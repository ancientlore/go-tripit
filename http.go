package tripit

import (
	"http"
)

// TripIt API URLs for methods. All return tripit.Response. All POSTs use tripit.Request
const (
	// GET object_type, object_id
	UrlGet = "https://api.tripit.com/v1/get/%s/id/%s/format/json"

	// GET object_type, filter_parameters*
	// supports: trip, object, points_program
	UrlList = "https://api.tripit.com/v1/list/%s%s/format/json"

	// filter_parameter, filter_value (you can have multiple)
	UrlListFilter = "/%s/%s"

	// POST
	UrlCreate = "https://api.tripit.com/v1/create/format/json"

	// POST object_type, object_id
	// supports: air, activity, car, cruise, directions, lodging, map, note, rail, restaurant, transport, trip
	UrlReplace = "https://api.tripit.com/v1/replace/%s/id/%s/format/json"

	// GET object_type, object_id
	// supports: air, activity, car, cruise, directions, lodging, map, note, rail, restaurant, transport, trip
	UrlDelete = "https://api.tripit.com/v1/delete/%s/id/%s/format/json"
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





type TripIt struct {
	httpClient *http.Client

}

func New(client *http.Client) *TripIt {
	return &TripIt{client}
}

func (t *TripIt) Get(objectType string, objectId uint) *Response {
	return nil
}

func (t *TripIt) List(objectType string, filterParms map[string]string) *Response {
	return nil
}

func (t *TripIt) Create(req *Request) *Response {
	return nil
}

func (t *TripIt) Replace(objectType string, objectId uint, req *Request) *Response {
	return nil
}

func (t *TripIt) Delete(objectType string, objectId uint) *Response {
	return nil
}

