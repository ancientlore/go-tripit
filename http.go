package tripit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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
	FilterPageNum        = "page_num"        // valid on trip, object. Integer, page number to retrieve
	FilterPageSize       = "page_size"       // valid on trip, object. Integer, number of items per page
)

// Authorizable is the interface for authorization objects.
type Authorizable interface {
	Authorize(request *http.Request, args map[string]string)
}

// TripIt class to used to communicate with the API.
type TripIt struct {
	baseUrl     string
	version     string
	httpClient  *http.Client
	credentials Authorizable
}

// New creates a new TripIt object using the given HTTP client and authorization object.
func New(apiUrl string, apiVersion string, client *http.Client, creds Authorizable) *TripIt {
	return &TripIt{apiUrl, apiVersion, client, creds}
}

// Makes an HTTP request to the TripIt API and returns the response.
func (t *TripIt) makeRequest(req *http.Request) (*Response, error) {
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	// Copy buffer and change @attributes to _attributes since json package doesn't support @
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)
	b := bytes.Replace(buf.Bytes(), []byte("\"@attributes\""), []byte("\"_attributes\""), -1)

	// debug logging
	// f, _ := os.Create("output.json")
	// defer f.Close()
	// f.Write(b)

	result := new(Response)
	err = json.Unmarshal(b, result)
	if err != nil {
		return nil, err
	}

	// debug logging
	// log.Print(result)

	return result, nil
}

// Get gets an Object of the given type and ID, and returns the Response object from TripIt.
// supports: air, activity, car, cruise, directions, lodging, map, note, rail, restaurant, transport, trip
func (t *TripIt) Get(objectType string, objectId uint) (*Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/get/%s/id/%d/format/json", t.baseUrl, t.version, objectType, objectId), nil)
	if err != nil {
		return nil, err
	}
	t.credentials.Authorize(req, nil)
	return t.makeRequest(req)
}

// List lists objects of the given type, filtered by the given filter parameters. Returns
// the response object from TripIt. To understand filter parameters and which filters
// can be combined, see the TripIt API documentation.
// supports: trip, object, points_program
func (t *TripIt) List(objectType string, filterParms map[string]string) (*Response, error) {
	var x string
	for p, v := range filterParms {
		x += fmt.Sprintf("/%s/%s", p, v)
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/list/%s%s/format/json", t.baseUrl, t.version, objectType, x), nil)
	if err != nil {
		return nil, err
	}
	t.credentials.Authorize(req, nil)
	return t.makeRequest(req)
}

// encodeForm encodes form arguments to send to TripIt
func encodeForm(r *Request) (*bytes.Buffer, map[string]string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, nil, err
	}
	s := string(b)
	// s = `{"Trip":{"start_date":"2011-12-09","end_date":"2011-12-27","primary_location":"Cancun, Mexico","display_name":"My Test Trip"}}`
	m := make(map[string][]string)
	m["json"] = []string{s}
	args := make(map[string]string)
	args["json"] = s
	return bytes.NewBuffer([]byte(url.Values(m).Encode())), args, nil
}

// Create creates an object in TripIt based on the given Request, returning the Response object from TripIt.
// supports: air, activity, car, cruise, directions, lodging, map, note, rail, restaurant, transport, trip
func (t *TripIt) Create(r *Request) (*Response, error) {
	buf, args, err := encodeForm(r)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/create/format/json", t.baseUrl, t.version), ioutil.NopCloser(buf))
	if err != nil {
		return nil, err
	}
	t.credentials.Authorize(req, args)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(buf.Len()))
	req.ContentLength = int64(buf.Len())
	return t.makeRequest(req)
}

// Replace replaces the object of the given type and ID with the new object in the Request. Returns
// the Response object from TripIt.
// supports: air, activity, car, cruise, directions, lodging, map, note, rail, restaurant, transport, trip
func (t *TripIt) Replace(objectType string, objectId uint, r *Request) (*Response, error) {
	b := new(bytes.Buffer)
	json := json.NewEncoder(b)
	err := json.Encode(r)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/replace/%s/id/%d/format/json", t.baseUrl, t.version, objectType, objectId), b)
	if err != nil {
		return nil, err
	}
	t.credentials.Authorize(req, nil)
	return t.makeRequest(req)
}

// Delete deletes the object of the given type and ID from TripIt, and returns the Response object
// from TripIt.
// supports: air, activity, car, cruise, directions, lodging, map, note, rail, restaurant, transport, trip
func (t *TripIt) Delete(objectType string, objectId uint) (*Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/delete/%s/id/%d/format/json", t.baseUrl, t.version, objectType, objectId), nil)
	if err != nil {
		return nil, err
	}
	t.credentials.Authorize(req, nil)
	return t.makeRequest(req)
}

// GetRequestToken is step 1 of the OAuth process. The function returns the token and secret
// from TripIt that is used in subsequent authentication requests. This token and secret
// is not the permanent one - if the user aborts the authentication process, these can
// be discarded.
func (t *TripIt) GetRequestToken() (map[string]string, error) {
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
		return nil, errors.New(resp.Status)
	}
	return parseQS(resp.Body)
}

// GetAccessToken gets the final OAuth token and token secret for an
// authenticated user. These should be saved with the user's ID for
// future used of the API on the user's behalf.
func (t *TripIt) GetAccessToken() (map[string]string, error) {
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
		return nil, errors.New(resp.Status)
	}
	return parseQS(resp.Body)
}

// parseQS parses the query string in the body and returns a simple map of the values.
func parseQS(body io.Reader) (map[string]string, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("reading http body: %v", err)
	}
	qm, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, fmt.Errorf("parsing body query: %v", err)
	}
	result := make(map[string]string)
	for k, v := range qm {
		result[k] = strings.TrimSpace(v[0])
	}
	return result, nil
}

// @TODO Add CRS elements
