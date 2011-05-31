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
	dst := make([]byte, len(pair)+4)
	base64.StdEncoding.Encode(dst, []byte(pair))
	token := string(dst)
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
}
