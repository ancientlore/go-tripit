package tripit

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestWarning(t *testing.T) {
	x := func() error { return &Warning{"Something went wrong", "trip", "2011-05-27T13:38:33"} }()
	t.Log(x)
}

func TestError(t *testing.T) {
	x := func() error { return &Error{"500", nil, "Something else went wrong", "trip", "2011-05-27T13:38:34"} }()
	t.Log(x)
}

func TestJsonWrite(t *testing.T) {
	var r Response
	log.Print("Marshal JSON")
	r.Error = new(ErrorVector)
	*r.Error = append(*r.Error, Error{Code_: "304", Description: "WTF"})
	b, _ := json.Marshal(r)
	os.Stdout.Write(b)
	fmt.Fprintf(os.Stdout, "\n")
}

func TestJsonRead(t *testing.T) {
	s := `
{
"Error": [
	{
		"code":"502",
		"description":"wtf",
		"timestamp":"2011-05-26T23:44:33"
	},
	{
		"code":"503",
		"description":"WTF Happened?",
		"timestamp":"2011-05-26T23:44:34"
	}
]
}`
	log.Print("Unmarshal JSON")
	var r Response
	b := make([]uint8, 300)
	l, err := strings.NewReader(s).Read(b)
	if err != nil {
		log.Print(err)
	}
	err = json.Unmarshal(b[0:l], &r)
	if err != nil {
		log.Print(err)
	}
	log.Print(r)
	log.Print("Marshal it back")
	b2, _ := json.Marshal(r)
	os.Stdout.Write(b2)
	fmt.Fprintf(os.Stdout, "\n")
}

func TestDateTime(t *testing.T) {

	d := &DateTime{"2009-11-10", "14:00:00", "America/Los_Angeles", "-08:00"}
	s, err := d.DateTime()
	log.Print("Parsed time: ", s, " err: ", err)

	log.Print(time.Now().Format(time.RFC3339))
	d.SetDateTime(time.Now())

	log.Print("Assigned time: ", d)
}
