package tripit

import (
	"http"
	"fmt"
	"encoding/base64"
)

type WebAuthCredential struct {
	username string
	password string
}

func (a *WebAuthCredential) Username() string {
	return a.username
}

func (a *WebAuthCredential) Password() string {
	return a.password
}

func (a *WebAuthCredential) Authorize(request *http.Request, args map[string]string) {
	pair := fmt.Sprintf("%s:%s", a.username, a.password)
	b := []byte(pair)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(dst, b)
	token := string(dst)
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
}
