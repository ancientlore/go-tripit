package tripit

import (
	"http"
	"json"
	"time"
	"strings"
	"rand"
	"fmt"
	"sort"
	"strconv"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"crypto/hmac"
)

// TripIt API URLs for OAuth
const (
	UrlObtainRequestToken            = "/oauth/request_token"                                                    // POST
	UrlObtainUserAuthorization       = "https://www.tripit.com/oauth/authorize?oauth_token=%s&oauth_callback=%s" // Redirect
	UrlObtainUserAuthorizationMobile = "https://m.tripit.com/oauth/authorize?oauth_token=%s&oauth_callback=%s"   // Redirect
	UrlObtainAccessToken             = "/oauth/access_token"                                                     // POST
)

// Signature method and version
const (
	OAUTH_SIGNATURE_METHOD = "HMAC-SHA1"
	OAUTH_VERSION          = "1.0"
)

// OAuth consumer credential for use with TripIt API
type OAuthConsumerCredential struct {
	oauthConsumerKey    string // Consumer key provided by TripIt
	oauthConsumerSecret string // Consumer secret provided by TripIt
	oauthOauthToken     string // OAuth token
	oauthTokenSecret    string // OAuth token secret
	oauthRequestorId    string // Requestor ID
}

// Get a credential with no token (to get a request token)
func NewOAuthRequestCredential(consumerKey string, consumerSecret string) *OAuthConsumerCredential {
	a := new(OAuthConsumerCredential)
	a.oauthConsumerKey = consumerKey
	a.oauthConsumerSecret = consumerSecret
	return a
}

// Get a 3 legged OAuth credential (request or authorized token)
func NewOAuth3LeggedCredential(consumerKey string, consumerSecret string, token string, tokenSecret string) *OAuthConsumerCredential {
	a := new(OAuthConsumerCredential)
	a.oauthConsumerKey = consumerKey
	a.oauthConsumerSecret = consumerSecret
	a.oauthOauthToken = token
	a.oauthTokenSecret = tokenSecret
	return a
}

// Get a 2 legged OAuth credential
func NewOAuth2LeggedCredential(consumerKey string, consumerSecret string, requestorId string) *OAuthConsumerCredential {
	a := new(OAuthConsumerCredential)
	a.oauthConsumerKey = consumerKey
	a.oauthConsumerSecret = consumerSecret
	a.oauthRequestorId = requestorId
	return a
}

// Returns the consumer key
func (a *OAuthConsumerCredential) OAuthConsumerKey() string {
	return a.oauthConsumerKey
}

// Returns the consumer secret
func (a *OAuthConsumerCredential) OAuthConsumerSecret() string {
	return a.oauthConsumerSecret
}

// Returns the OAuth token
func (a *OAuthConsumerCredential) OAuthOAuthToken() string {
	return a.oauthOauthToken
}

// Returns the consumer credential
func (a *OAuthConsumerCredential) OAuthTokenSecret() string {
	return a.oauthTokenSecret
}

// Returns the requestor ID
func (a *OAuthConsumerCredential) OAuthRequestorId() string {
	return a.oauthRequestorId
}

// Adds the authorization header for OAuth to the request, including any additional arguments.
// Additional arguments are used in signature generation.
func (a *OAuthConsumerCredential) Authorize(request *http.Request, args map[string]string) {
	request.Header.Set("Authorization", a.generateAuthorizationHeader(request, args))
}

// Validates the URL's OAuth signature in the given url
func (a *OAuthConsumerCredential) ValidateSignature(url string) bool {
	u, err := http.ParseURL(url)
	if err != nil {
		return false
	}
	q, err := http.ParseQuery(u.RawQuery)
	if err != nil {
		return false
	}
	u.RawQuery = ""
	u.Fragment = ""
	args := make(map[string]string)
	for k, v := range q {
		args[k] = v[0]
	}
	return q["oauth_signature"][0] == a.generateSignature("GET", u.String(), args)
}

// Returns the OAuth parameters for a given session
func (a *OAuthConsumerCredential) GetSessionParameters(redirectUrl string, action string) string {
	params := a.generateOAuthParameters("GET", action, map[string]string{"redirect_url": redirectUrl})
	params["redirect_url"] = redirectUrl
	params["action"] = action
	b, _ := json.Marshal(params)
	return string(b)
}

// Generates the authorization header string
func (a *OAuthConsumerCredential) generateAuthorizationHeader(request *http.Request, args map[string]string) string {
	httpMethod := strings.ToUpper(request.Method)
	realm := request.URL.Scheme + "://" + request.URL.Host
	httpUrl := request.URL.Scheme + "://" + request.URL.Host + request.URL.Path
	s := fmt.Sprintf("OAuth realm=\"%s\",", realm)
	p := a.generateOAuthParameters(httpMethod, httpUrl, args)
	arr := make([]string, len(p))
	i := 0
	for k, v := range p {
		arr[i] = fmt.Sprintf("%s=\"%s\"", http.URLEscape(k), http.URLEscape(v))
		i++
	}
	s += strings.Join(arr, ",")
	return s
}

// Generates the OAuth parameters and stores them in a map
func (a *OAuthConsumerCredential) generateOAuthParameters(httpMethod string, httpUrl string, args map[string]string) map[string]string {
	p := map[string]string{
		"oauth_consumer_key":     a.oauthConsumerKey,
		"oauth_nonce":            generateNonce(),
		"oauth_timestamp":        strconv.Itoa64(time.LocalTime().Seconds()),
		"oauth_signature_method": OAUTH_SIGNATURE_METHOD,
		"oauth_version":          OAUTH_VERSION,
	}
	if a.oauthOauthToken != "" {
		p["oauth_token"] = a.oauthOauthToken
	}
	if a.oauthRequestorId != "" {
		p["xoauth_requestor_id"] = a.oauthRequestorId
	}
	oauthParmsForBaseString := make(map[string]string)
	for k, v := range p {
		oauthParmsForBaseString[k] = v
	}
	if args != nil {
		for k, v := range args {
			oauthParmsForBaseString[k] = v
		}
	}
	p["oauth_signature"] = a.generateSignature(httpMethod, httpUrl, oauthParmsForBaseString)
	return p
}

// Generates the OAuth signature for a given URL
func (a *OAuthConsumerCredential) generateSignature(httpMethod string, baseUrl string, params map[string]string) string {
	params["oauth_signature"] = "", false
	arr := sort.StringArray(make([]string, len(params)))
	i := 0
	for k, v := range params {
		arr[i] = fmt.Sprintf("%s=%s", http.URLEscape(k), http.URLEscape(v))
		i++
	}
	arr.Sort()
	sigBaseString := strings.Join([]string{httpMethod, http.URLEscape(baseUrl), http.URLEscape(strings.Join(arr, "&"))}, "&")
	key := a.oauthConsumerSecret + "&" + a.oauthTokenSecret
	h := hmac.NewSHA1([]byte(key))
	h.Write([]byte(sigBaseString))
	b := h.Sum()
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(dst, b)
	return string(dst)
}

// Generates a unique one-time-use value
func generateNonce() string {
	arr := make([]string, 40)
	for i := 0; i < 40; i++ {
		arr[i] = string(rand.Int31n(10))
	}
	s := string(time.Nanoseconds()) + strings.Join(arr, "")
	h := md5.New()
	fmt.Fprintf(h, "%s", s)
	return hex.EncodeToString(h.Sum())
}
