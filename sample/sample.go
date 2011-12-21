package main

import (
	"flag"
	"log"
	"template"
	"http"
	"fmt"
	"strconv"
	"os"
	"json"
	"../_obj/tripit"
	"url"
)

var addr = flag.String("addr", "localhost:8080", "HTTP Service Address")
var oauthConsumerKey = flag.String("key", "", "OAuth Consumer Key from TripIt")
var oauthConsumerSecret = flag.String("secret", "", "OAuth Consumer Secret from TripIt")
var url_ = flag.String("url", tripit.ApiUrl, "TripIt API URL")

var indexT = template.Must(template.ParseFile("index.html"))
var tripsT = template.Must(template.ParseFile("trips.html"))
var detailsT = template.Must(template.ParseFile("details.html"))
var editT = template.Must(template.ParseFile("edit.html"))
var errorT = template.Must(template.ParseFile("error.html"))
var apierrorT = template.Must(template.ParseFile("apierror.html"))

type getsess struct {
	id    int
	reply chan map[string]string
}

var session chan getsess

func sessionManager() {
	var sessions []map[string]string
	for {
		select {
		case r := <-session:
			if r.id >= 0 && r.id < len(sessions) {
				r.reply <- sessions[r.id]
			} else {
				s := make(map[string]string)
				sessions = append(sessions, s)
				s["id"] = strconv.Itoa(len(sessions) - 1)
				r.reply <- s
			}
		}
	}
}

func getSession(w http.ResponseWriter, req *http.Request) map[string]string {
	var gs getsess
	gs.id = -1
	var err os.Error
	for _, c := range req.Cookies() {
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
	http.Handle("/details", http.HandlerFunc(Details))
	http.Handle("/list", http.HandlerFunc(List))
	http.Handle("/edit", http.HandlerFunc(Edit))
	http.Handle("/save", http.HandlerFunc(Save))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Index(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	m := make(map[string]interface{})
	if sess["oauth_token"] != "" {
		m["Authorized"] = true
	} else {
		m["NotAuthorized"] = true
	}
	indexT.Execute(w, m)
}

func Auth(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	cred := tripit.NewOAuthRequestCredential(*oauthConsumerKey, *oauthConsumerSecret)
	log.Print("Cred ", cred)
	var client http.Client
	t := tripit.New(*url_, tripit.ApiVersion, &client, cred)
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
	http.Redirect(w, req, fmt.Sprintf(tripit.UrlObtainUserAuthorization, url.QueryEscape(m["oauth_token"]), url.QueryEscape(aurl)), http.StatusFound)
}

func CheckAuth(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	cred := tripit.NewOAuth3LeggedCredential(*oauthConsumerKey, *oauthConsumerSecret, sess["oauth_token"], sess["oauth_token_secret"])
	var client http.Client
	t := tripit.New(*url_, tripit.ApiVersion, &client, cred)
	m, err := t.GetAccessToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	sess["oauth_token"] = m["oauth_token"]
	sess["oauth_token_secret"] = m["oauth_token_secret"]
	aurl := fmt.Sprintf("http://%s/", *addr)
	http.Redirect(w, req, aurl, http.StatusFound)
}

func Trips(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	cred := tripit.NewOAuth3LeggedCredential(*oauthConsumerKey, *oauthConsumerSecret, sess["oauth_token"], sess["oauth_token_secret"])
	var client http.Client
	t := tripit.New(*url_, tripit.ApiVersion, &client, cred)
	resp, err := t.List(tripit.ObjectTypeTrip, map[string]string{tripit.FilterTraveler: "true", tripit.FilterPast: "true"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	m := make(map[string]interface{})
	m["Result"] = resp
	if (resp.Warning != nil && len(resp.Warning) > 0) || (resp.Error != nil && len(resp.Error) > 0) {
		if resp.Warning != nil {
			m["Warning"] = resp.Warning
		}
		if resp.Error != nil {
			m["Error"] = resp.Error
		}
		apierrorT.Execute(w, m)
		return
	}
	m["Trip"] = resp.Trip
	b, _ := json.MarshalIndent(resp, "", "\t")
	m["JSON"] = string(b)
	tripsT.Execute(w, m)
}

func Details(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	cred := tripit.NewOAuth3LeggedCredential(*oauthConsumerKey, *oauthConsumerSecret, sess["oauth_token"], sess["oauth_token_secret"])
	var client http.Client
	t := tripit.New(*url_, tripit.ApiVersion, &client, cred)
	q, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	objType := q["t"][0]
	objId, err := strconv.Atoui(q["id"][0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	resp, err := t.Get(objType, objId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	m := make(map[string]interface{})
	m["Result"] = resp
	if (resp.Warning != nil && len(resp.Warning) > 0) || (resp.Error != nil && len(resp.Error) > 0) {
		if resp.Warning != nil {
			m["Warning"] = resp.Warning
		}
		if resp.Error != nil {
			m["Error"] = resp.Error
		}
		apierrorT.Execute(w, m)
		return
	}
	b, _ := json.MarshalIndent(resp, "", "\t")
	m["JSON"] = string(b)
	detailsT.Execute(w, m)
}

func List(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	cred := tripit.NewOAuth3LeggedCredential(*oauthConsumerKey, *oauthConsumerSecret, sess["oauth_token"], sess["oauth_token_secret"])
	var client http.Client
	t := tripit.New(*url_, tripit.ApiVersion, &client, cred)
	q, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	objType := q["t"][0]
	filters := make(map[string]string)
	for k, v := range q {
		if k != "objType" {
			filters[k] = v[0]
		}
	}

	resp, err := t.List(objType, filters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	m := make(map[string]interface{})
	m["Result"] = resp
	if (resp.Warning != nil && len(resp.Warning) > 0) || (resp.Error != nil && len(resp.Error) > 0) {
		if resp.Warning != nil {
			m["Warning"] = resp.Warning
		}
		if resp.Error != nil {
			m["Error"] = resp.Error
		}
		apierrorT.Execute(w, m)
		return
	}
	b, _ := json.MarshalIndent(resp, "", "\t")
	m["JSON"] = string(b)
	detailsT.Execute(w, m)
}

func Edit(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	cred := tripit.NewOAuth3LeggedCredential(*oauthConsumerKey, *oauthConsumerSecret, sess["oauth_token"], sess["oauth_token_secret"])
	var client http.Client
	t := tripit.New(*url_, tripit.ApiVersion, &client, cred)
	q, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return
	}
	objType := q["t"][0]
	var objId uint = 0
	tmp, ok := q["id"]
	if ok {
		objId, err = strconv.Atoui(tmp[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorT.Execute(w, err)
			return
		}
	}
	var trip *tripit.Trip
	m := make(map[string]interface{})
	if objId > 0 {
		resp, err := t.Get(objType, objId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorT.Execute(w, err)
			return
		}
		if (resp.Warning != nil && len(resp.Warning) > 0) || (resp.Error != nil && len(resp.Error) > 0) {
			if resp.Warning != nil {
				m["Warning"] = resp.Warning
			}
			if resp.Error != nil {
				m["Error"] = resp.Error
			}
			apierrorT.Execute(w, m)
			return
		}
		m["Result"] = resp
		trip = (resp.Trip)[0]
	} else {
		trip = new(tripit.Trip)
	}
	m["Trip"] = trip
	editT.Execute(w, m)
}

func Save(w http.ResponseWriter, req *http.Request) {
	sess := getSession(w, req)
	cred := tripit.NewOAuth3LeggedCredential(*oauthConsumerKey, *oauthConsumerSecret, sess["oauth_token"], sess["oauth_token_secret"])
	var client http.Client
	t := tripit.New(*url_, tripit.ApiVersion, &client, cred)
	err := req.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorT.Execute(w, err)
		return

	}
	objType := req.Form["t"][0]
	var objId uint = 0
	tmp, ok := req.Form["id"]
	if ok && tmp[0] != "" && tmp[0] != "<nil>" {
		objId, err = strconv.Atoui(tmp[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorT.Execute(w, err)
			return
		}
	}
	var trip tripit.Trip
	m := make(map[string]interface{})
	if objId > 0 {
		resp, err := t.Get(objType, objId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorT.Execute(w, err)
			return
		}
		if (resp.Warning != nil && len(resp.Warning) > 0) || (resp.Error != nil && len(resp.Error) > 0) {
			if resp.Warning != nil {
				m["Warning"] = resp.Warning
			}
			if resp.Error != nil {
				m["Error"] = resp.Error
			}
			apierrorT.Execute(w, m)
			return
		}
		trip = *(resp.Trip)[0]
		trip.DisplayName = req.Form["DisplayName"][0]
		trip.Description = req.Form["Description"][0]
		request := new(tripit.Request)
		request.Trip = &trip
		resp, err = t.Replace(objType, objId, request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorT.Execute(w, err)
			return
		}
		if (resp.Warning != nil && len(resp.Warning) > 0) || (resp.Error != nil && len(resp.Error) > 0) {
			if resp.Warning != nil {
				m["Warning"] = resp.Warning
			}
			if resp.Error != nil {
				m["Error"] = resp.Error
			}
			apierrorT.Execute(w, m)
			return
		}
	} else {
		trip.DisplayName = req.Form["DisplayName"][0]
		trip.Description = req.Form["Description"][0]
		request := new(tripit.Request)
		request.Trip = &trip
		b, err := json.Marshal(request)
		log.Print(string(b))
		resp, err := t.Create(request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorT.Execute(w, err)
			return
		}
		if (resp.Warning != nil && len(resp.Warning) > 0) || (resp.Error != nil && len(resp.Error) > 0) {
			if resp.Warning != nil {
				m["Warning"] = resp.Warning
			}
			if resp.Error != nil {
				m["Error"] = resp.Error
			}
			apierrorT.Execute(w, m)
			return
		}
	}

	aurl := fmt.Sprintf("http://%s/details?t=%s&id=%d", *addr, objType, objId)
	http.Redirect(w, req, aurl, http.StatusFound)
}
