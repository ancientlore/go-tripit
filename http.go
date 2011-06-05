package tripit

import (
	"http"
	"os"
	"fmt"
	"json"
	"io"
	"bytes"
	"strings"
)

// TripIt API information
const (
	ApiUrl     = "https://api.tripit.com"
	ApiVersion = "v1"
)

// List objects
const (
	ListTrip          = "trip"
	ListObject        = "object"
	ListPointsProgram = "points_program"
)

// Filter Parameters
const (
	FilterNone           = ""                // valid on trip, object, points_program
	FilterTraveler       = "traveler"        // valid on trip, object. Values: true, false, all
	FilterPast           = "past"            // valid on trip, object. Values: true, false
	FilterModifiedSince  = "modified_since"  // valid on trip, object. Values: integer
	FilterIncludeObjects = "include_objects" // valid on trip. Values: true, false
	FilterTripId         = "trip_id"         // valid on object. Values: integer trip id
	FilterType           = "type"            // valid on object. Values: all object types
)

// Interface for authorization objects
type Authorizable interface {
	Authorize(request *http.Request, args map[string]string)
}

// TripIt class to used to communicate with the API
type TripIt struct {
	baseUrl     string
	version     string
	httpClient  *http.Client
	credentials Authorizable
}

// Creates a new TripIt object using the given HTTP client and authorization object
func New(apiUrl string, apiVersion string, client *http.Client, creds Authorizable) *TripIt {
	return &TripIt{apiUrl, apiVersion, client, creds}
}

// makes request
func (t *TripIt) makeRequest(req *http.Request) (*Response, os.Error) {
	t.credentials.Authorize(req, nil)
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, os.NewError(resp.Status)
	}
	//f, _ := os.Create("output.json")
	//defer f.Close()
	//io.Copy(f, resp.Body)
	//f.Seek(0, 0)
	// json := json.NewDecoder(f)

	// Copy buffer and change @attributes to _attributes since json package doesn't support @
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)
	b := bytes.Replace(buf.Bytes(), []byte("\"@attributes\""), []byte("\"_attributes\""), -1)
	result := new(Response)
	err = json.Unmarshal(b, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// supports: air, activity, car, cruise, directions, lodging, map, note, rail, restaurant, transport, trip
func (t *TripIt) Get(objectType string, objectId uint) (*Response, os.Error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/get/%s/id/%u/format/json", t.baseUrl, t.version, objectType, objectId), nil)
	if err != nil {
		return nil, err
	}
	return t.makeRequest(req)
}

// supports: trip, object, points_program
func (t *TripIt) List(objectType string, filterParms map[string]string) (*Response, os.Error) {
	var x string
	for p, v := range filterParms {
		x += fmt.Sprintf("/%s/%s", p, v)
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/list/%s%s/format/json", t.baseUrl, t.version, objectType, x), nil)
	if err != nil {
		return nil, err
	}
	return t.makeRequest(req)
}

// supports: air, activity, car, cruise, directions, lodging, map, note, rail, restaurant, transport, trip
func (t *TripIt) Create(r *Request) (*Response, os.Error) {
	b := new(bytes.Buffer)
	json := json.NewEncoder(b)
	err := json.Encode(r)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/create/format/json", t.baseUrl, t.version), b)
	if err != nil {
		return nil, err
	}
	return t.makeRequest(req)
}

// supports: air, activity, car, cruise, directions, lodging, map, note, rail, restaurant, transport, trip
func (t *TripIt) Replace(objectType string, objectId uint, r *Request) (*Response, os.Error) {
	b := new(bytes.Buffer)
	json := json.NewEncoder(b)
	err := json.Encode(r)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/replace/%s/id/%u/format/json", t.baseUrl, t.version, objectType, objectId), b)
	if err != nil {
		return nil, err
	}
	return t.makeRequest(req)
}

// supports: air, activity, car, cruise, directions, lodging, map, note, rail, restaurant, transport, trip
func (t *TripIt) Delete(objectType string, objectId uint) (*Response, os.Error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/delete/%s/id/%u/format/json", t.baseUrl, t.version, objectType, objectId), nil)
	if err != nil {
		return nil, err
	}
	return t.makeRequest(req)
}

func (t *TripIt) GetRequestToken() (map[string]string, os.Error) {
	req, err := http.NewRequest("GET", t.baseUrl+UrlObtainRequestToken, nil)
	if err != nil {
		return nil, err
	}
	t.credentials.Authorize(req, nil)
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, os.NewError(resp.Status)
	}
	return parseQS(resp.Body)
}

func (t *TripIt) GetAccessToken() (map[string]string, os.Error) {
	req, err := http.NewRequest("GET", t.baseUrl+UrlObtainAccessToken, nil)
	if err != nil {
		return nil, err
	}
	t.credentials.Authorize(req, nil)
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, os.NewError(resp.Status)
	}
	return parseQS(resp.Body)
}

func parseQS(body io.Reader) (map[string]string, os.Error) {
	buf := make([]byte, 1024) // assume oauth token response won't be larger
	l, err := body.Read(buf)
	if err != nil {
		return nil, err
	}
	qm, err := http.ParseQuery(string(buf[0:l]))
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for k, v := range qm {
		result[k] = strings.TrimSpace(v[0])
	}
	return result, nil
}

// @TODO Add CRS elements
