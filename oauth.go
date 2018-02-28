package tripit

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
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

// OAuthConsumerCredential is the OAuth consumer credential for use with TripIt API.
type OAuthConsumerCredential struct {
	oauthConsumerKey    string // Consumer key provided by TripIt
	oauthConsumerSecret string // Consumer secret provided by TripIt
	oauthOauthToken     string // OAuth token
	oauthTokenSecret    string // OAuth token secret
	oauthRequestorId    string // Requestor ID
}

// NewOAuthRequestCredential gets a credential with no token (to get a request token).
func NewOAuthRequestCredential(consumerKey string, consumerSecret string) *OAuthConsumerCredential {
	a := new(OAuthConsumerCredential)
	a.oauthConsumerKey = consumerKey
	a.oauthConsumerSecret = consumerSecret
	return a
}

// NewOAuth3LeggedCredential gets a 3 legged OAuth credential (request or authorized token).
func NewOAuth3LeggedCredential(consumerKey string, consumerSecret string, token string, tokenSecret string) *OAuthConsumerCredential {
	a := new(OAuthConsumerCredential)
	a.oauthConsumerKey = consumerKey
	a.oauthConsumerSecret = consumerSecret
	a.oauthOauthToken = token
	a.oauthTokenSecret = tokenSecret
	return a
}

// NewOAuth2LeggedCredential gets a 2 legged OAuth credential.
func NewOAuth2LeggedCredential(consumerKey string, consumerSecret string, requestorId string) *OAuthConsumerCredential {
	a := new(OAuthConsumerCredential)
	a.oauthConsumerKey = consumerKey
	a.oauthConsumerSecret = consumerSecret
	a.oauthRequestorId = requestorId
	return a
}

// OAuthConsumerKey returns the consumer key.
func (a *OAuthConsumerCredential) OAuthConsumerKey() string {
	return a.oauthConsumerKey
}

// OAuthConsumerSecret returnss the consumer secret.
func (a *OAuthConsumerCredential) OAuthConsumerSecret() string {
	return a.oauthConsumerSecret
}

// OAuthOAuthToken returns the OAuth token.
func (a *OAuthConsumerCredential) OAuthOAuthToken() string {
	return a.oauthOauthToken
}

// OAuthTokenSecret returns the OAuth token secret.
func (a *OAuthConsumerCredential) OAuthTokenSecret() string {
	return a.oauthTokenSecret
}

// OAuthRequestorId returns the requestor ID.
func (a *OAuthConsumerCredential) OAuthRequestorId() string {
	return a.oauthRequestorId
}

// Authorize adds the authorization header for OAuth to the request, including any additional arguments.
// Additional arguments are used in signature generation.
func (a *OAuthConsumerCredential) Authorize(request *http.Request, args map[string]string) {
	request.Header.Set("Authorization", a.generateAuthorizationHeader(request, args))
}

// ValidateSignature validates the URL's OAuth signature in the given url.
func (a *OAuthConsumerCredential) ValidateSignature(url_ string) bool {
	u, err := url.Parse(url_)
	if err != nil {
		return false
	}
	q, err := url.ParseQuery(u.RawQuery)
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

// GetSessionParameters returns the OAuth parameters for a given session.
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
		arr[i] = fmt.Sprintf("%s=\"%s\"", url.QueryEscape(k), url.QueryEscape(v))
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
		"oauth_timestamp":        strconv.FormatInt(time.Now().Unix(), 10),
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
	delete(params, "oauth_signature")
	arr := make([]string, len(params))
	i := 0
	for k, v := range params {
		arr[i] = fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(v))
		i++
	}
	sort.Sort(sort.StringSlice(arr))
	sigBaseString := strings.Join([]string{httpMethod, url.QueryEscape(baseUrl), url.QueryEscape(strings.Join(arr, "&"))}, "&")
	key := a.oauthConsumerSecret + "&" + a.oauthTokenSecret
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(sigBaseString))
	b := h.Sum(nil)
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
	s := string(time.Now().Unix()) + strings.Join(arr, "")
	h := md5.New()
	fmt.Fprintf(h, "%s", s)
	return hex.EncodeToString(h.Sum(nil))
}
