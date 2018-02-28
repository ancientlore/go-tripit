package tripit

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// WebAuthCredential us used for web authorization. Web authorization is for testing
// and needs to be enabled on your TripIt account.
type WebAuthCredential struct {
	username string // Account user name
	password string // Account password
}

// Username returns the user name.
func (a *WebAuthCredential) Username() string {
	return a.username
}

// Password returns the password.
func (a *WebAuthCredential) Password() string {
	return a.password
}

// Authorize adds the authorization header to the HTTP request, including any additional arguments.
// Note that web authorization ignores extra arguments.
func (a *WebAuthCredential) Authorize(request *http.Request, args map[string]string) {
	pair := fmt.Sprintf("%s:%s", a.username, a.password)
	b := []byte(pair)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(dst, b)
	token := string(dst)
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", token))
}
