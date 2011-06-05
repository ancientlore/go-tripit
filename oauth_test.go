package tripit

import (
	"testing"
	"http"
	"strings"
)

func TestOAuth(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://www.google.com", nil)
	c := NewOAuth2LeggedCredential("foo", "bar", "app")
	c.Authorize(r, nil)
	t.Log(r.Header["Authorization"][0])
	arr := strings.Split(r.Header["Authorization"][0], ",", -1)
	p := make([]string, len(arr))
	j := 0
	for i := range arr {
		kva := strings.Split(arr[i], "=", -1)
		if kva[0] != "OAuth realm" {
			t.Log(kva)
			p[j] = kva[0] + "=" + strings.Trim(kva[1], "\"")
			j++
		}
	}
	s := "http://www.google.com?" + strings.Join(p, "&")
	t.Log(s)
	if c.ValidateSignature(s) != true {
		t.Error("Signature is invalid")
	}
}
