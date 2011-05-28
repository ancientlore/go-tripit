package tripit

import (
	"http"
	//"fmt"
)

// TripIt API URLs for OAuth
const (
	UrlObtainRequestToken            = "https://api.tripit.com/oauth/request_token"                              // POST
	UrlObtainUserAuthorization       = "https://www.tripit.com/oauth/authorize?oauth_token=%s&oauth_callback=%s" // Redirect
	UrlObtainUserAuthorizationMobile = "https://m.tripit.com/oauth/authorize?oauth_token=%s&oauth_callback=%s"   // Redirect
	UrlObtainAccessToken             = "https://api.tripit.com/oauth/access_token"                               // POST
)

const (
	OAUTH_SIGNATURE_METHOD = "HMAC-SHA1"
	OAUTH_VERSION = "1.0"
)

type OAuthConsumerCredential struct {
	oauthConsumerKey string
	oauthConsumerSecret string
	oauthOauthToken string
	oauthTokenSecret string
	oauthRequestorId string
}

// Get a credential with no token (to get a request token)
func newOAuthRequestCredential(consumerKey string, consumerSecret string) *OAuthConsumerCredential {
	a := new(OAuthConsumerCredential)
	a.oauthConsumerKey = consumerKey
	a.oauthConsumerSecret = consumerSecret
	return a
}

// Get a 3 legged OAuth credential (request or authorized token)
func newOAuth3LeggedCredential(consumerKey string, consumerSecret string, token string, tokenSecret string) *OAuthConsumerCredential {
	a := new(OAuthConsumerCredential)
	a.oauthConsumerKey = consumerKey
	a.oauthConsumerSecret = consumerSecret
	a.oauthOauthToken = token
	a.oauthTokenSecret = tokenSecret
	return a
}

// Get a 2 legged OAuth credential
func newOAuth2LeggedCredential(consumerKey string, consumerSecret string, requestorId string) *OAuthConsumerCredential {
	a := new(OAuthConsumerCredential)
	a.oauthConsumerKey = consumerKey
	a.oauthConsumerSecret = consumerSecret
	a.oauthRequestorId = requestorId
	return a
}

func (a *OAuthConsumerCredential) OAuthConsumerKey() string {
	return a.oauthConsumerKey
}

func (a *OAuthConsumerCredential) OAuthConsumerSecret() string {
	return a.oauthConsumerSecret
}

func (a *OAuthConsumerCredential) OAuthOAuthToken() string {
	return a.oauthOauthToken
}

func (a *OAuthConsumerCredential) OAuthTokenSecret() string {
	return a.oauthTokenSecret
}

func (a *OAuthConsumerCredential) OAuthRequestorId() string {
	return a.oauthRequestorId
}

func (a *OAuthConsumerCredential) Authorize(request *http.Request, args map[string]string) {

}

func (a *OAuthConsumerCredential) ValidateSignature(url string) bool {
	return true
}

func (a* OAuthConsumerCredential) GetSessionParameters(redirectUrl string, action string) string {
	return ""
}

func (a *OAuthConsumerCredential) generateAuthorizationHeader(request *http.Request, args map[string]string) string {
	return ""
}

func (a *OAuthConsumerCredential) generateOAuthParameters(httpMethod string, httpUrl string, args map[string]string) map[string]string {
	return make(map[string]string)
}

func (a *OAuthConsumerCredential) generateSignature(httpMethod string, baseUrl string, params map[string]string) string {
	return ""
}


