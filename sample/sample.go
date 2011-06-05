package main

import (
	"flag"
	"log"
	"template"
	"http"
	"fmt"
	"strconv"
	"os"
	"container/vector"
	"../_obj/tripit"
)

var addr = flag.String("addr", "localhost:8080", "HTTP Service Address")
var oauthConsumerKey = flag.String("key", "", "OAuth Consumer Key from TripIt")
var oauthConsumerSecret = flag.String("secret", "", "OAuth Consumer Secret from TripIt")
var url = flag.String("url", tripit.ApiUrl, "TripIt API URL")

var indexT = template.MustParseFile("index.html", nil)
var tripsT = template.MustParseFile("trips.html", nil)
var errorT = template.MustParseFile("error.html", nil)
var apierrorT = template.MustParseFile("apierror.html", nil)

type getsess struct {
	id    int
	reply chan map[string]string
}

var session chan getsess

func sessionManager() {
	var sessions vector.Vector
	for {
		select {
		case r := <-session:
			if r.id >= 0 && r.id < sessions.Len() {
				r.reply <- sessions.At(r.id).(map[string]string)
			} else {
				s := make(map[string]string)
				sessions.Push(s)
				s["id"] = strconv.Itoa(sessions.Len() - 1)
				r.reply <- s
			}
		}
	}
}

func getSession(w http.ResponseWriter, req *http.Request) map[string]string {
	var gs getsess
	gs.id = -1
	var err os.Error
	for _, c := range req.Cookie {
		if c.Name == "samplesession" {
			gs.id, err = strconv.Atoi(c.Value)
			if err != nil {
				gs.id = -1
			}
			break
		}
	}
	gs.reply = make(chan map[string]string)
	session <- gs
	m := <-gs.reply
	sid, _ := strconv.Atoi(m["id"])
	if gs.id < 0 || gs.id != sid {
		w.Header().Add("Set-Cookie", "samplesession="+m["id"])
	}
	return m
}

func startSessionManager() {
	session = make(chan getsess, 8)
	go sessionManager()
}

func main() {
	flag.Parse()
	startSessionManager()
	http.Handle("/", http.HandlerFunc(Index))
	http.Handle("/auth", http.HandlerFunc(Auth))
	http.Handle("/auth2", http.HandlerFunc(CheckAuth))
	http.Handle("/trips", http.HandlerFunc(Trips))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Index(w http.ResponseWriter, req *http.Request) {
	getSession(w, req)
	indexT.Execute(w, nil)
}

func Auth(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	cred := tripit.NewOAuthRequestCredential(*oauthConsumerKey, *oauthConsumerSecret)
	log.Print("Cred ", cred)
	var client http.Client
	t := tripit.New(*url, tripit.ApiVersion, &client, cred)
	m, err := t.GetRequestToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	log.Print("requested token ", m)
	sess["oauth_token"] = m["oauth_token"]
	sess["oauth_token_secret"] = m["oauth_token_secret"]
	aurl := fmt.Sprintf("http://%s/auth2", *addr)
	http.Redirect(w, req, fmt.Sprintf(tripit.UrlObtainUserAuthorization, http.URLEscape(m["oauth_token"]), http.URLEscape(aurl)), http.StatusFound)
}

func CheckAuth(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	cred := tripit.NewOAuth3LeggedCredential(*oauthConsumerKey, *oauthConsumerSecret, sess["oauth_token"], sess["oauth_token_secret"])
	var client http.Client
	t := tripit.New(*url, tripit.ApiVersion, &client, cred)
	m, err := t.GetAccessToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	sess["oauth_token"] = m["oauth_token"]
	sess["oauth_token_secret"] = m["oauth_token_secret"]
	aurl := fmt.Sprintf("http://%s/trips", *addr)
	http.Redirect(w, req, aurl, http.StatusFound)
}

func Trips(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	cred := tripit.NewOAuth3LeggedCredential(*oauthConsumerKey, *oauthConsumerSecret, sess["oauth_token"], sess["oauth_token_secret"])
	var client http.Client
	t := tripit.New(*url, tripit.ApiVersion, &client, cred)
	resp, err := t.List(tripit.ObjectTypeTrip, map[string]string{tripit.FilterTraveler: "true", tripit.FilterPast: "true"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	if (resp.Warning != nil && len(resp.Warning) > 0) || (resp.Error != nil && len(resp.Error) > 0) {
		apierrorT.Execute(w, resp)
		return
	}
	tripsT.Execute(w, resp)
}
