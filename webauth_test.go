package tripit

import (
	"testing"
	"http"
)

func TestWebAuth(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://www.google.com", nil)
	c := &WebAuthCredential{"user@site.com", "password"}
	c.Authorize(r, nil)
	if r.Header["Authorization"][0] != "Basic dXNlckBzaXRlLmNvbTpwYXNzd29yZA==" {
		t.Errorf("Basic authorization header incorrect: %s", r.Header["Authorization"][0])
	}
}
