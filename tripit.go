// Package tripit is a Go API library for accessing the TripIt service. The API supports
// two forms of authorization - simple web authorization and OAuth. The library uses TripIt's
// JSON interface, and has structs representing all of the TripIt types. Within these structs,
// elements ending in an underscore have access functions that set or get the value in a more
// pleasant form for use in Go programs.
package tripit

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// TripIt Object Types
const (
	ObjectTypeAir        = "air"
	ObjectTypeActivity   = "activity"
	ObjectTypeCar        = "car"
	ObjectTypeCruise     = "cruise"
	ObjectTypeDirections = "directions"
	ObjectTypeLodging    = "lodging"
	ObjectTypeMap        = "map"
	ObjectTypeNote       = "note"
	ObjectTypeRail       = "rail"
	ObjectTypeRestaurant = "restaurant"
	ObjectTypeTransport  = "transport"
	ObjectTypeTrip       = "trip"
)

// Request contains the objects that can be sent to TripIt in a request.
type Request struct {
	Invitation       []Invitation      `json:"Invitation,omitempty"`       // optional
	Trip             *Trip             `json:"Trip,omitempty"`             // optional
	ActivityObject   *ActivityObject   `json:"ActivityObject,omitempty"`   // optional
	AirObject        *AirObject        `json:"AirObject,omitempty"`        // optional
	CarObject        *CarObject        `json:"CarObject,omitempty"`        // optional
	CruiseObject     *CruiseObject     `json:"CruiseObject,omitempty"`     // optional
	DirectionsObject *DirectionsObject `json:"DirectionsObject,omitempty"` // optional
	LodgingObject    *LodgingObject    `json:"LodgingObject,omitempty"`    // optional
	MapObject        *MapObject        `json:"MapObject,omitempty"`        // optional
	NoteObject       *NoteObject       `json:"NoteObject,omitempty"`       // optional
	RailObject       *RailObject       `json:"RailObject,omitempty"`       // optional
	RestaurantObject *RestaurantObject `json:"RestaurantObject,omitempty"` // optional
	TransportObject  *TransportObject  `json:"TransportObject,omitempty"`  // optional
}

// Error is returned from TripIt on error conditions.
type Error struct {
	Code              int     `json:"code,string,omitempty" xml:"code"`                               // read-only
	DetailedErrorCode float64 `json:"detailed_error_code,string,omitempty" xml:"detailed_error_code"` // optional, read-only
	Description       string  `json:"description,omitempty" xml:"description"`                        // read-only
	EntityType        string  `json:"entity_type,omitempty" xml:"entity_type"`                        // read-only
	Timestamp         string  `json:"timestamp,omitempty" xml:"timestamp"`                            // read-only, xs:datetime
}

// Time returns a time.Time object for the Timestamp.
func (e *Error) Time() (time.Time, error) {
	return time.Parse(time.RFC3339, e.Timestamp)
}

// String returns a string containing the error information.
func (e *Error) String() string {
	return e.Error()
}

// Error returns a string containing the error information.
func (e *Error) Error() string {
	return fmt.Sprintf("TripIt Error %d: %s", e.Code, e.Description)
}

// Warning is returned from TripIt to indicate warning conditions
type Warning struct {
	Description string `json:"description,omitempty"` // read-only
	EntityType  string `json:"entity_type,omitempty"` // read-only
	Timestamp   string `json:"timestamp,omitempty"`   // read-only, xs:datetime
}

// Time returns a time.Time object for the Timestamp.
func (w *Warning) Time() (time.Time, error) {
	return time.Parse(time.RFC3339, w.Timestamp)
}

// String returns a string containing the warning information.
func (w *Warning) String() string {
	return w.Error()
}

// Error returns a string containing the warning information.
func (w *Warning) Error() string {
	return fmt.Sprintf("TripIt Warning: %s", w.Description)
}

// Note that the Vectors are pointers - otherwise the JSON marshaler doesn't notice the custom methods

// Response represents a TripIt API Response
type Response struct {
	Timestamp        string                    `json:"timestamp,omitempty" xml:"timestamp"`
	NumBytes         int                       `json:"num_bytes,string,omitempty" xml:"num_bytes"`
	Error            ErrorVector               `json:"Error,omitempty" xml:"Error"`                       // optional
	Warning          WarningVector             `json:"Warning,omitempty" xml:"Warning"`                   // optional
	Trip             TripPtrVector             `json:"Trip,omitempty" xml:"Trip"`                         // optional
	ActivityObject   ActivityObjectPtrVector   `json:"ActivityObject,omitempty" xml:"ActivityObject"`     // optional
	AirObject        AirObjectPtrVector        `json:"AirObject,omitempty" xml:"AirObject"`               // optional
	CarObject        CarObjectPtrVector        `json:"CarObject,omitempty" xml:"CarObject"`               // optional
	CruiseObject     CruiseObjectPtrVector     `json:"CruiseObject,omitempty" xml:"CruiseObject"`         // optional
	DirectionsObject DirectionsObjectPtrVector `json:"DirectionsObject,omitempty" xml:"DirectionsObject"` // optional
	LodgingObject    LodgingObjectPtrVector    `json:"LodgingObject,omitempty" xml:"LodgingObject"`       // optional
	MapObject        MapObjectPtrVector        `json:"MapObject,omitempty" xml:"MapObject"`               // optional
	NoteObject       NoteObjectPtrVector       `json:"NoteObject,omitempty" xml:"NoteObject"`             // optional
	RailObject       RailObjectPtrVector       `json:"RailObject,omitempty" xml:"RailObject"`             // optional
	RestaurantObject RestaurantObjectPtrVector `json:"RestaurantObject,omitempty" xml:"RestaurantObject"` // optional
	TransportObject  TransportObjectPtrVector  `json:"TransportObject,omitempty" xml:"TransportObject"`   // optional
	WeatherObject    WeatherObjectVector       `json:"WeatherObject,omitempty" xml:"WeatherObject"`       // optional
	PointsProgram    PointsProgramVector       `json:"PointsProgram,omitempty" xml:"PointsProgram"`       // optional
	Profile          ProfileVector             `json:"Profile,omitempty" xml:"Profile"`                   // optional

	PageNumber json.Number `json:"page_num,omitempty" xml:"page_num"`   // when pagination is activated
	PageSize   json.Number `json:"page_size,omitempty" xml:"page_size"` // when pagination is activated
	MaxPage    json.Number `json:"max_page,omitempty" xml:"max_page"`   // when pagination is activated

	// @TODO need to add invitee stuff
}

// Time returns a time.Time object for the Timestamp
func (r *Response) Time() (time.Time, error) {
	t, err := strconv.ParseInt(r.Timestamp, 10, 64)
	if err != nil {
		return time.Unix(0, 0), err
	}
	return time.Unix(t, 0).UTC(), nil
}

// Address represents the address of a location.
// For create, use either:
// - address for single-line addresses.
// - addr1, addr2, city, state, zip, and country for multi-line addresses.
// Multi-line address will be ignored if single-line address is present.
// See documentation for more information.
type Address struct {
	Address   string  `json:"address,omitempty" xml:"address"`            // optional
	Addr1     string  `json:"addr1,omitempty" xml:"addr1"`                // optional
	Addr2     string  `json:"addr2,omitempty" xml:"addr2"`                // optional
	City      string  `json:"city,omitempty" xml:"city"`                  // optional
	State     string  `json:"state,omitempty" xml:"state"`                // optional
	Zip       string  `json:"zip,omitempty" xml:"zip"`                    // optional
	Country   string  `json:"country,omitempty" xml:"country"`            // optional
	Latitude  float64 `json:"latitude,string,omitempty" xml:"latitude"`   // optional, read-only
	Longitude float64 `json:"longitude,string,omitempty" xml:"longitude"` // optional, read-only
}

// Traveler contains information about a traveler.
type Traveler struct {
	FirstName                string `json:"first_name,omitempty" xml:"first_name"`                                 // optional
	MiddleName               string `json:"middle_name,omitempty" xml:"middle_name"`                               // optional
	LastName                 string `json:"last_name,omitempty" xml:"last_name"`                                   // optional
	FrequentTravelerNum      string `json:"frequent_traveler_num,omitempty" xml:"frequent_traveler_num"`           // optional
	FrequentTravelerSupplier string `json:"frequent_traveler_supplier,omitempty" xml:"frequent_traveler_supplier"` // optional
	MealPreference           string `json:"meal_preference,omitempty" xml:"meal_preference"`                       // optional
	SeatPreference           string `json:"seat_preference,omitempty" xml:"seat_preference"`                       // optional
	TicketNum                string `json:"ticket_num,omitempty" xml:"ticket_num"`                                 //optional
}

// Flight status values
const (
	FlightStatusNotMonitorable = 100
	FlightStatusNotMonitored   = 200
	FlightStatusScheduled      = 300
	FlightStatusOnTime         = 301
	FlightStatusInFlightOnTime = 302
	FlightStatusArrivedOnTime  = 303
	FlightStatusCancelled      = 400
	FlightStatusDelayed        = 401
	FlightStatusInFlightLate   = 402
	FlightStatusArrivedLate    = 403
	FlightStatusDiverted       = 404
)

// FlightStatus fields are read-only and only available for monitored TripIt Pro AirSegments.
type FlightStatus struct {
	ScheduledDepartureDateTime *DateTime `json:"ScheduledDepartureDateTime,omitempty" xml:"ScheduledDepartureDateTime"` // optional, read-only
	EstimatedDepartureDateTime *DateTime `json:"EstimatedDepartureDateTime,omitempty" xml:"EstimatedDepartureDateTime"` // optional, read-only
	ScheduledArrivalDateTime   *DateTime `json:"ScheduledArrivalDateTime,omitempty" xml:"ScheduledArrivalDateTime"`     // optional, read-only
	EstimatedArrivalDateTime   *DateTime `json:"EstimatedArrivalDateTime,omitempty" xml:"EstimatedArrivalDateTime"`     // optional, read-only
	FlightStatus               int       `json:"flight_status,string,omitempty" xml:"flight_status"`                    // optional, read-only
	IsConnectionAtRisk         bool      `json:"is_connection_at_risk,string,omitempty" xml:"is_connection_at_risk"`    // optional, read-only
	DepartureTerminal          string    `json:"departure_terminal,omitempty" xml:"departure_terminal"`                 // optional, read-only
	DepartureGate              string    `json:"departure_gate,omitempty" xml:"departure_gate"`                         // optional, read-only
	ArrivalTerminal            string    `json:"arrival_terminal,omitempty" xml:"arrival_terminal"`                     // optional, read-only
	ArrivalGate                string    `json:"arrival_gate,omitempty" xml:"arrival_gate"`                             // optional, read-only
	LayoverMinutes             string    `json:"layover_minutes,omitempty" xml:"layover_minutes"`                       // optional, read-only
	BaggageClaim               string    `json:"baggage_claim,omitempty" xml:"baggage_claim"`                           // optional, read-only
	DivertedAirportCode        string    `json:"diverted_airport_code,omitempty" xml:"diverted_airport_code"`           // optional, read-only
	LastModified               string    `json:"last_modified,omitempty" xml:"last_modified"`                           // read-only
}

// LastModifiedTime returns a time.Time object for LastModified.
func (fs *FlightStatus) LastModifiedTime() (time.Time, error) {
	l, err := strconv.ParseInt(fs.LastModified, 10, 64)
	if err != nil {
		return time.Unix(0, 0), err
	}
	return time.Unix(l, 0).UTC(), nil
}

// Image stores information about images.
type Image struct {
	Caption string `json:"caption,omitempty" xml:"caption"` // optional
	Url     string `json:"url" xml:"url"`
}

// DateTime Stores date and time zone information, for example:
// {
//	 "date":"2009-11-10",
//   "time":"14:00:00",
//    "timezone":"America\/Los_Angeles",
//    "utc_offset":"-08:00"
// }
type DateTime struct {
	Date      string `json:"date,omitempty" xml:"date"`             // optional, xs:date
	Time      string `json:"time,omitempty" xml:"time"`             // optional, xs:time
	Timezone  string `json:"timezone,omitempty" xml:"timezone"`     // optional, read-only
	UtcOffset string `json:"utc_offset,omitempty" xml:"utc_offset"` // optional, read-only
}

// GetTime converts the time to a time.Time.
func (dt DateTime) GetTime() (time.Time, error) {
	if dt.UtcOffset == "" {
		return time.Parse(time.RFC3339, fmt.Sprintf("%sT%sZ", dt.Date, dt.Time))
	}
	return time.Parse(time.RFC3339, fmt.Sprintf("%sT%s%s", dt.Date, dt.Time, dt.UtcOffset))
}

// SetTime sets the values of the DateTime strucure from a time.Time.
func (dt *DateTime) SetTime(t time.Time) {
	dt.Date = t.Format("2006-01-02")
	dt.Time = t.Format("15:04:05")
	dt.UtcOffset = t.Format("-07:00")
	dt.Timezone = t.Format("MST")
}

// PointsProgram contains information about tracked travel programs for TripIt Pro users.
// All PointsProgram elements are read-only.
type PointsProgram struct {
	Id                  uint                          `json:"id,string,omitempty" xml:"id"`                                       // read-only
	Name                string                        `json:"name,omitempty" xml:"name"`                                          // optional, read-only
	AccountNumber       string                        `json:"account_number,omitempty" xml:"account_number"`                      // optional, read-only
	AccountLogin        string                        `json:"account_login,omitempty" xml:"account_login"`                        // optional, read-only
	Balance             string                        `json:"balance,omitempty" xml:"balance"`                                    // optional, read-only
	EliteStatus         string                        `json:"elite_status,omitempty" xml:"elite_status"`                          // optional, read-only
	EliteNextStatus     string                        `json:"elite_next_status,omitempty" xml:"elite_next_status"`                // optional, read-only
	EliteYtdQualify     string                        `json:"elite_ytd_qualify,omitempty" xml:"elite_ytd_qualify"`                // optional, read-only
	EliteNeedToEarn     string                        `json:"elite_need_to_earn,omitempty" xml:"elite_need_to_earn"`              // optional, read-only
	LastModified        string                        `json:"last_modified,omitempty" xml:"last_modified"`                        // read-only
	TotalNumActivities  int                           `json:"total_num_activities,string,omitempty" xml:"total_num_activities"`   // read-only
	TotalNumExpirations int                           `json:"total_num_expirations,string,omitempty" xml:"total_num_expirations"` // read-only
	ErrorMessage        string                        `json:"error_message,omitempty" xml:"error_message"`                        // optional, read-only
	Activity            PointsProgramActivityVector   `json:"Activity,omitempty" xml:"Activity"`                                  // optional, read-only
	Expiration          PointsProgramExpirationVector `json:"Expiration,omitempty" xml:"Expiration"`                              // optional, read-only
}

// LastModifiedTime returns a time.Time object for LastModified.
func (pp *PointsProgram) LastModifiedTime() (time.Time, error) {
	v, err := strconv.ParseInt(pp.LastModified, 10, 64)
	if err != nil {
		return time.Unix(0, 0), err
	}
	return time.Unix(v, 0).UTC(), nil
}

// PointsProgramActivity contains program transactions
// All PointsProgramActivity elements are read-only
type PointsProgramActivity struct {
	Date        string `json:"date,omitempty" xml:"date"`               // read-only, xs:date
	Description string `json:"description,omitempty" xml:"description"` // optional, read-only
	Base        string `json:"base,omitempty" xml:"base"`               // optional, read-only
	Bonus       string `json:"bonus,omitempty" xml:"bonus"`             // optional, read-only
	Total       string `json:"total,omitempty" xml:"total"`             // optional, read-only
}

// Time returns a time.Time object for Date.
// Note: This won't have proper time zone information
func (pa *PointsProgramActivity) Time() (time.Time, error) {
	return time.Parse("2006-01-02", pa.Date)
}

// PointsProgramExpiration elements are read-only.
type PointsProgramExpiration struct {
	Date   string `json:"date,omitempty" xml:"date"`     // read-only, xs:date
	Amount string `json:"amount,omitempty" xml:"amount"` // optional, read-only
}

// Time returns a time.Time object for Date.
// Note: This won't have proper time zone information.
func (pe *PointsProgramExpiration) Time() (time.Time, error) {
	return time.Parse("2006-01-02", pe.Date)
}

// TripShare contains information about which users a trip is shared with.
type TripShare struct {
	TripId            uint `json:"trip_id,string,omitempty" xml:"trip_id"`
	IsTraveler        bool `json:"is_traveler,string,omitempty" xml:"is_traveler"`
	IsReadOnly        bool `json:"is_read_only,string,omitempty" xml:"is_read_only"`
	IsSentWithDetails bool `json:"is_sent_with_details,string,omitempty" xml:"is_sent_with_details"`
}

// ConnectionRequest stores connection request data.
type ConnectionRequest struct {
}

// Invitation contains a list of users invited to see the trip.
type Invitation struct {
	EmailAddresses    []string           `json:"EmailAddresses,omitempty" xml:"EmailAddresses"`
	TripShare         *TripShare         `json:"TripShare,omitempty" xml:"TripShare"`                 // optional
	ConnectionRequest *ConnectionRequest `json:"ConnectionRequest,omitempty" xml:"ConnectionRequest"` // optional
	Message           string             `json:"message,omitempty" xml:"message"`                     // optional
}

// Profile contains user information.
// All Profile elements are read-only.
type Profile struct {
	Attributes            ProfileAttributes      `json:"_attributes" xml:"attributes"`                                // read-only
	ProfileEmailAddresses *ProfileEmailAddresses `json:"ProfileEmailAddresses,omitempty" xml:"ProfileEmailAddresses"` // optional, read-only
	GroupMemberships      *GroupMemberships      `json:"GroupMemberships,omitempty" xml:"GroupMemberships"`           // optional, read-only
	IsClient              bool                   `json:"is_client,string,omitempty" xml:"is_client"`                  // read-only
	IsPro                 bool                   `json:"is_pro,string,omitempty" xml:"is_pro"`                        // read-only
	ScreenName            string                 `json:"screen_name,omitempty" xml:"screen_name"`                     // read-only
	PublicDisplayName     string                 `json:"public_display_name,omitempty" xml:"public_display_name"`     // read-only
	ProfileUrl            string                 `json:"profile_url,omitempty" xml:"profile_url"`                     // read-only
	HomeCity              string                 `json:"home_city,omitempty" xml:"home_city"`                         // optional, read-only
	Company               string                 `json:"company,omitempty" xml:"company"`                             // optional, read-only
	AboutMeInfo           string                 `json:"about_me_info,omitempty" xml:"about_me_info"`                 // optional, read-only
	PhotoUrl              string                 `json:"photo_url,omitempty" xml:"photo_url"`                         // optional, read-only
	ActivityFeedUrl       string                 `json:"activity_feed_url,omitempty" xml:"activity_feed_url"`         // optional, read-only
	AlertsFeedUrl         string                 `json:"alerts_feed_url,omitempty" xml:"alerts_feed_url"`             // optional, read-only
	IcalUrl               string                 `json:"ical_url,omitempty" xml:"ical_url"`                           // optional, read-only
}

// ProfileEmailAddresses contains the list of email addresses for a user.
type ProfileEmailAddresses struct {
	ProfileEmailAddress ProfileEmailAddressVector `json:"ProfileEmailAddress,omitempty" xml:"ProfileEmailAddress"`
}

// GroupMemberships contains a list of groups that the user is a member of.
type GroupMemberships struct {
	Group GroupVector `json:"Group,omitempty" xml:"Group"` // optional, read-only
}

// ProfileAttributes represent links to profiles.
type ProfileAttributes struct {
	Ref string `json:"ref,omitempty" xml:"ref"` // read-only
}

// ProfileEmailAddress contains an email address and its properties.
// All ProfileEmailAddress elements are read-only.
type ProfileEmailAddress struct {
	Address      string `json:"address" xml:"address"`                                // read-only
	IsAutoImport bool   `json:"is_auto_import,string,omitempty" xml:"is_auto_import"` // read-only
	IsConfirmed  bool   `json:"is_confirmed,string,omitempty" xml:"is_confirmed"`     // read-only
	IsPrimary    bool   `json:"is_primary,string,omitempty" xml:"is_primary"`         // read-only
}

// Group contains data about a group in TripIt.
// All Group elements are read-only.
type Group struct {
	DisplayName string `json:"display_name,omitempty" xml:"display_name"` // read-only
	Url         string `json:"url" xml:"url"`                             // read-only
}

// Invitee stores attributes about invitees to a trip.
// All Invitee elements are read-only.
type Invitee struct {
	IsReadOnly bool              `json:"is_read_only,string,omitempty" xml:"is_read_only"` // read-only
	IsTraveler bool              `json:"is_traveler,string,omitempty" xml:"is_traveler"`   // read-only
	Attributes InviteeAttributes `json:"_attributes" xml:"attributes"`                     // read-only, Use the profile_ref attribute to reference a Profile
}

// InviteeAttributes are used to link to user profiles.
type InviteeAttributes struct {
	ProfileRef string `json:"profile_ref" xml:"profile_ref"` // read-only, used to reference a profile
}

// TripCrsRemark is a reservation system remark.
// All TripCrsRemark elements are read-only.
type TripCrsRemark struct {
	RecordLocator string `json:"record_locator,omitempty" xml:"record_locator"` // read-only
	Notes         string `json:"notes,omitempty" xml:"notes"`                   // read-only
}

// ClosenessMatch refers to nearby users.
// All ClosenessMatch elements are read-only.
type ClosenessMatch struct {
	Attributes ClosenessMatchAttributes `json:"_attributes" xml:"attributes"` // read-only, Use the profile_ref attribute to reference a Profile
}

// ClosenessMatchAttributes links to profiles of nearby users.
type ClosenessMatchAttributes struct {
	ProfileRef string `json:"profile_ref" xml:"profile_ref"` // read-only, Use the profile_ref attribute to reference a Profile
}

// Trip represents a trip in the TripIt model.
type Trip struct {
	ClosenessMatches       *ClosenessMatches `json:"ClosenessMatches,omitempty" xml:"ClosenessMatches"`                 // optional, ClosenessMatches are read-only
	TripInvitees           *TripInvitees     `json:"TripInvitees,omitempty" xml:"TripInvitees"`                         // optional, TripInvitees are read-only
	TripCrsRemarks         *TripCrsRemarks   `json:"TripCrsRemarks,omitempty" xml:"TripCrsRemarks"`                     // optional, TripCrsRemarks are read-only
	Id                     string            `json:"id,omitempty" xml:"id"`                                             // optional, id is a read-only field
	RelativeUrl            string            `json:"relative_url,omitempty" xml:"relative_url"`                         // optional, relative_url is a read-only field
	StartDate              string            `json:"start_date,omitempty" xml:"start_date"`                             // optional, xs:date
	EndDate                string            `json:"end_date,omitempty" xml:"end_date"`                                 // optional, xs:date
	Description            string            `json:"description,omitempty" xml:"description"`                           // optional
	DisplayName            string            `json:"display_name,omitempty" xml:"display_name"`                         // optional
	ImageUrl               string            `json:"image_url,omitempty" xml:"image_url"`                               // optional
	IsPrivate              bool              `json:"is_private,string,omitempty" xml:"is_private"`                      // optional
	PrimaryLocation        string            `json:"primary_location,omitempty" xml:"primary_location"`                 // optional
	PrimaryLocationAddress *Address          `json:"primary_location_address,omitempty" xml:"primary_location_address"` // optional, PrimaryLocationAddress is a read-only field
}

// TripInvitees are people invited to view a trip.
type TripInvitees struct {
	Invitee InviteeVector `json:"Invitee,omitempty" xml:"Invitee"` // optional, TripInvitees are read-only
}

// ClosenessMatches are TripIt users who are near this trip.
type ClosenessMatches struct {
	ClosenessMatch ClosenessMatchVector `json:"Match,omitempty" xml:"Match"` // optional, ClosenessMatches are read-only
}

// TripCrsRemarks are remarks from a reservation system.
type TripCrsRemarks struct {
	TripCrsRemark TripCrsRemarkVector `json:"TripCrsRemark,omitempty" xml:"TripCrsRemark"` // optional, TripCrsRemarks are read-only
}

// StartTime returns a time.Time object for StartDate.
// Note: This won't have proper time zone information.
func (t *Trip) StartTime() (time.Time, error) {
	return time.Parse("2006-01-02", t.StartDate)
}

// EndTime returns a time.Time object for EndDate.
// Note: This won't have proper time zone information.
func (t *Trip) EndTime() (time.Time, error) {
	return time.Parse("2006-01-02", t.EndDate)
}

// AirObject contains data about a flight.
type AirObject struct {
	Id                   string              `json:"id,omitempty" xml:"id"`                                         // optional, read-only
	TripId               string              `json:"trip_id,omitempty" xml:"trip_id"`                               // optional
	IsClientTraveler     bool                `json:"is_client_traveler,string,omitempty" xml:"is_client_traveler"`  // optional, read-only
	RelativeUrl          string              `json:"relative_url,omitempty" xml:"relative_url"`                     // optional, read-only
	DisplayName          string              `json:"display_name,omitempty" xml:"display_name"`                     // optional
	Image                ImagePtrVector      `json:"Image,omitempty" xml:"Image"`                                   // optional
	CancellationDateTime *DateTime           `json:"CancellationDateTime,omitempty" xml:"CancellationDateTime"`     // optional
	BookingDate          string              `json:"booking_date,omitempty" xml:"booking_date"`                     // optional, xs:date
	BookingRate          string              `json:"booking_rate,omitempty" xml:"booking_rate"`                     // optional
	BookingSiteConfNum   string              `json:"booking_site_conf_num,omitempty" xml:"booking_site_conf_num"`   // optional
	BookingSiteName      string              `json:"booking_site_name,omitempty" xml:"booking_site_name"`           // optional
	BookingSitePhone     string              `json:"booking_site_phone,omitempty" xml:"booking_site_phone"`         // optional
	BookingSiteUrl       string              `json:"booking_site_url,omitempty" xml:"booking_site_url"`             // optional
	RecordLocator        string              `json:"record_locator,omitempty" xml:"record_locator"`                 // optional
	SupplierConfNum      string              `json:"supplier_conf_num,omitempty" xml:"supplier_conf_num"`           // optional
	SupplierContact      string              `json:"supplier_contact,omitempty" xml:"supplier_contact"`             // optional
	SupplierEmailAddress string              `json:"supplier_email_address,omitempty" xml:"supplier_email_address"` // optional
	SupplierName         string              `json:"supplier_name,omitempty" xml:"supplier_name"`                   // optional
	SupplierPhone        string              `json:"supplier_phone,omitempty" xml:"supplier_phone"`                 // optional
	SupplierUrl          string              `json:"supplier_url,omitempty" xml:"supplier_url"`                     // optional
	IsPurchased          bool                `json:"is_purchased,string,omitempty" xml:"is_purchased"`              // optional
	Notes                string              `json:"notes,omitempty" xml:"notes"`                                   // optional
	Restrictions         string              `json:"restrictions,omitempty" xml:"restrictions"`                     // optional
	TotalCost            string              `json:"total_cost,omitempty" xml:"total_cost"`                         // optional
	Segment              AirSegmentPtrVector `json:"Segment,omitempty" xml:"Segment"`
	Traveler             TravelerPtrVector   `json:"Traveler,omitempty" xml:"Traveler"` // optional
}

// BookingTime returns a time.Time object for BookingDate.
// Note: This won't have proper time zone information.
func (r *AirObject) BookingTime() (time.Time, error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// AirSegment contains details about individual flights.
type AirSegment struct {
	Status                *FlightStatus `json:"Status,omitempty" xml:"Status"`                                          // optional
	StartDateTime         *DateTime     `json:"StartDateTime,omitempty" xml:"StartDateTime"`                            // optional
	EndDateTime           *DateTime     `json:"EndDateTime,omitempty" xml:"EndDateTime"`                                // optional
	StartAirportCode      string        `json:"start_airport_code,omitempty" xml:"start_airport_code"`                  // optional
	StartAirportLatitude  float64       `json:"start_airport_latitude,string,omitempty" xml:"start_airport_latitude"`   // optional, read-only
	StartAirportLongitude float64       `json:"start_airport_longitude,string,omitempty" xml:"start_airport_longitude"` // optional, read-only
	StartCityName         string        `json:"start_city_name,omitempty" xml:"start_city_name"`                        // optional
	StartGate             string        `json:"start_gate,omitempty" xml:"start_gate"`                                  // optional
	StartTerminal         string        `json:"start_terminal,omitempty" xml:"start_terminal"`                          // optional
	EndAirportCode        string        `json:"end_airport_code,omitempty" xml:"end_airport_code"`                      // optional
	EndAirportLatitude    float64       `json:"end_airport_latitude,string,omitempty" xml:"end_airport_latitude"`       // optional, read-only
	EndAirportLongitude   float64       `json:"end_airport_longitude,string,omitempty" xml:"end_airport_longitude"`     // optional, read-only
	EndCityName           string        `json:"end_city_name,omitempty" xml:"end_city_name"`                            // optional
	EndGate               string        `json:"end_gate,omitempty" xml:"end_gate"`                                      // optional
	EndTerminal           string        `json:"end_terminal,omitempty" xml:"end_terminal"`                              // optional
	MarketingAirline      string        `json:"marketing_airline,omitempty" xml:"marketing_airline"`                    // optional
	MarketingAirlineCode  string        `json:"marketing_airline_code,omitempty" xml:"marketing_airline_code"`          // optional, read-only
	MarketingFlightNumber string        `json:"marketing_flight_number,omitempty" xml:"marketing_flight_number"`        // optional
	OperatingAirline      string        `json:"operating_airline,omitempty" xml:"operating_airline"`                    // optional
	OperatingAirlineCode  string        `json:"operating_airline_code,omitempty" xml:"operating_airline_code"`          // optional, read-only
	OperatingFlightNumber string        `json:"operating_flight_number,omitempty" xml:"operating_flight_number"`        // optional
	AlternativeFlightsUrl string        `json:"alternate_flights_url,omitempty" xml:"alternate_flights_url"`            // optional, read-only
	Aircraft              string        `json:"aircraft,omitempty" xml:"aircraft"`                                      // optional
	AircraftDisplayName   string        `json:"aircraft_display_name,omitempty" xml:"aircraft_display_name"`            // optional, read-only
	Distance              string        `json:"distance,omitempty" xml:"distance"`                                      // optional
	Duration              string        `json:"duration,omitempty" xml:"duration"`                                      // optional
	Entertainment         string        `json:"entertainment,omitempty" xml:"entertainment"`                            // optional
	Meal                  string        `json:"meal,omitempty" xml:"meal"`                                              // optional
	Notes                 string        `json:"notes,omitempty" xml:"notes"`                                            // optional
	OntimePerc            string        `json:"ontime_perc,omitempty" xml:"ontime_perc"`                                // optional
	Seats                 string        `json:"seats,omitempty" xml:"seats"`                                            // optional
	ServiceClass          string        `json:"service_class,omitempty" xml:"service_class"`                            // optional
	Stops                 string        `json:"stops,omitempty" xml:"stops"`                                            // optional
	BaggageClaim          string        `json:"baggage_claim,omitempty" xml:"baggage_claim"`                            // optional
	CheckInUrl            string        `json:"check_in_url,omitempty" xml:"check_in_url"`                              // optional
	ConflictResolutionUrl string        `json:"conflict_resolution_url,omitempty" xml:"conflict_resolution_url"`        // optional, read-only
	IsHidden              bool          `json:"is_hidden,string,omitempty" xml:"is_hidden"`                             // optional, read-only
	Id                    string        `json:"id,omitempty" xml:"id"`                                                  // optional, read-only
}

// LodgingObject contains information about hotels or other lodging.
// hotel cancellation remarks should be in restrictions.
// hotel room description should be in notes.
// hotel average daily rate should be in booking_rate.
type LodgingObject struct {
	Id                   string            `json:"id,omitempty" xml:"id"`                                         // optional, read-only
	TripId               string            `json:"trip_id,omitempty" xml:"trip_id"`                               // optional
	IsClientTraveler     bool              `json:"is_client_traveler,string,omitempty" xml:"is_client_traveler"`  // optional, read-only
	RelativeUrl          string            `json:"relative_url,omitempty" xml:"relative_url"`                     // optional, read-only
	DisplayName          string            `json:"display_name,omitempty" xml:"display_name"`                     // optional
	Image                ImagePtrVector    `json:"Image,omitempty" xml:"Image"`                                   // optional
	CancellationDateTime *DateTime         `json:"CancellationDateTime,omitempty" xml:"CancellationDateTime"`     // optional
	BookingDate          string            `json:"booking_date,omitempty" xml:"booking_date"`                     // optional, xs:date
	BookingRate          string            `json:"booking_rate,omitempty" xml:"booking_rate"`                     // optional
	BookingSiteConfNum   string            `json:"booking_site_conf_num,omitempty" xml:"booking_site_conf_num"`   // optional
	BookingSiteName      string            `json:"booking_site_name,omitempty" xml:"booking_site_name"`           // optional
	BookingSitePhone     string            `json:"booking_site_phone,omitempty" xml:"booking_site_phone"`         // optional
	BookingSiteUrl       string            `json:"booking_site_url,omitempty" xml:"booking_site_url"`             // optional
	RecordLocator        string            `json:"record_locator,omitempty" xml:"record_locator"`                 // optional
	SupplierConfNum      string            `json:"supplier_conf_num,omitempty" xml:"supplier_conf_num"`           // optional
	SupplierContact      string            `json:"supplier_contact,omitempty" xml:"supplier_contact"`             // optional
	SupplierEmailAddress string            `json:"supplier_email_address,omitempty" xml:"supplier_email_address"` // optional
	SupplierName         string            `json:"supplier_name,omitempty" xml:"supplier_name"`                   // optional
	SupplierPhone        string            `json:"supplier_phone,omitempty" xml:"supplier_phone"`                 // optional
	SupplierUrl          string            `json:"supplier_url,omitempty" xml:"supplier_url"`                     // optional
	IsPurchased          bool              `json:"is_purchased,string,omitempty" xml:"is_purchased"`              // optional
	Notes                string            `json:"notes,omitempty" xml:"notes"`                                   // optional
	Restrictions         string            `json:"restrictions,omitempty" xml:"restrictions"`                     // optional
	TotalCost            string            `json:"total_cost,omitempty" xml:"total_cost"`                         // optional
	StartDateTime        *DateTime         `json:"StartDateTime,omitempty" xml:"StartDateTime"`                   // optional
	EndDateTime          *DateTime         `json:"EndDateTime,omitempty" xml:"EndDateTime"`                       // optional
	Address              *Address          `json:"Address,omitempty" xml:"Address"`                               // optional
	Guest                TravelerPtrVector `json:"Guest,omitempty" xml:"Guest"`                                   // optional
	NumberGuests         string            `json:"number_guests,omitempty" xml:"number_guests"`                   // optional
	NumberRooms          string            `json:"number_rooms,omitempty" xml:"number_rooms"`                     // optional
	RoomType             string            `json:"room_type,omitempty" xml:"room_type"`                           // optional
}

// BookingTime returns a time.Time object for BookingDate.
// Note: This won't have proper time zone information.
func (r *LodgingObject) BookingTime() (time.Time, error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// CarObject contains information about rental cars.
// car cancellation remarks should be in restrictions.
// car pickup instructions should be in notes.
// car daily rate should be in booking_rate.
type CarObject struct {
	Id                   string            `json:"id,omitempty" xml:"id"`                                         // optional, read-only
	TripId               string            `json:"trip_id,omitempty" xml:"trip_id"`                               // optional
	IsClientTraveler     bool              `json:"is_client_traveler,string,omitempty" xml:"is_client_traveler"`  // optional, read-only
	RelativeUrl          string            `json:"relative_url,omitempty" xml:"relative_url"`                     // optional, read-only
	DisplayName          string            `json:"display_name,omitempty" xml:"display_name"`                     // optional
	Image                ImagePtrVector    `json:"Image,omitempty" xml:"Image"`                                   // optional
	CancellationDateTime *DateTime         `json:"CancellationDateTime,omitempty" xml:"CancellationDateTime"`     // optional
	BookingDate          string            `json:"booking_date,omitempty" xml:"booking_date"`                     // optional, xs:date
	BookingRate          string            `json:"booking_rate,omitempty" xml:"booking_rate"`                     // optional
	BookingSiteConfNum   string            `json:"booking_site_conf_num,omitempty" xml:"booking_site_conf_num"`   // optional
	BookingSiteName      string            `json:"booking_site_name,omitempty" xml:"booking_site_name"`           // optional
	BookingSitePhone     string            `json:"booking_site_phone,omitempty" xml:"booking_site_phone"`         // optional
	BookingSiteUrl       string            `json:"booking_site_url,omitempty" xml:"booking_site_url"`             // optional
	RecordLocator        string            `json:"record_locator,omitempty" xml:"record_locator"`                 // optional
	SupplierConfNum      string            `json:"supplier_conf_num,omitempty" xml:"supplier_conf_num"`           // optional
	SupplierContact      string            `json:"supplier_contact,omitempty" xml:"supplier_contact"`             // optional
	SupplierEmailAddress string            `json:"supplier_email_address,omitempty" xml:"supplier_email_address"` // optional
	SupplierName         string            `json:"supplier_name,omitempty" xml:"supplier_name"`                   // optional
	SupplierPhone        string            `json:"supplier_phone,omitempty" xml:"supplier_phone"`                 // optional
	SupplierUrl          string            `json:"supplier_url,omitempty" xml:"supplier_url"`                     // optional
	IsPurchased          bool              `json:"is_purchased,string,omitempty" xml:"is_purchased"`              // optional
	Notes                string            `json:"notes,omitempty" xml:"notes"`                                   // optional
	Restrictions         string            `json:"restrictions,omitempty" xml:"restrictions"`                     // optional
	TotalCost            string            `json:"total_cost,omitempty" xml:"total_cost"`                         // optional
	StartDateTime        *DateTime         `json:"StartDateTime,omitempty" xml:"StartDateTime"`                   // optional
	EndDateTime          *DateTime         `json:"EndDateTime,omitempty" xml:"EndDateTime"`                       // optional
	StartLocationAddress *Address          `json:"StartLocationAddress,omitempty" xml:"StartLocationAddress"`     // optional
	EndLocationAddress   *Address          `json:"EndLocationAddress,omitempty" xml:"EndLocationAddress"`         // optional
	Driver               TravelerPtrVector `json:"Driver,omitempty" xml:"Driver"`                                 // optional
	StartLocationHours   string            `json:"start_location_hours,omitempty" xml:"start_location_hours"`     // optional
	StartLocationName    string            `json:"start_location_name,omitempty" xml:"start_location_name"`       // optional
	StartLocationPhone   string            `json:"start_location_phone,omitempty" xml:"start_location_phone"`     // optional
	EndLocationHours     string            `json:"end_location_hours,omitempty" xml:"end_location_hours"`         // optional
	EndLocationName      string            `json:"end_location_name,omitempty" xml:"end_location_name"`           // optional
	EndLocationPhone     string            `json:"end_location_phone,omitempty" xml:"end_location_phone"`         // optional
	CarDescription       string            `json:"car_description,omitempty" xml:"car_description"`               // optional
	CarType              string            `json:"car_type,omitempty" xml:"car_type"`                             // optional
	MileageCharges       string            `json:"mileage_charges,omitempty" xml:"mileage_charges"`               // optional
}

// BookingTime returns a time.Time object for BookingDate.
// Note: This won't have proper time zone information.
func (r *CarObject) BookingTime() (time.Time, error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// RailObject contains information about trains.
type RailObject struct {
	Id                   string               `json:"id,omitempty"`                        // optional, read-only
	TripId               string               `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler     bool                 `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl          string               `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string               `json:"display_name,omitempty"`              // optional
	Image                ImagePtrVector       `json:"Image,omitempty"`                     // optional
	CancellationDateTime *DateTime            `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string               `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string               `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string               `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string               `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string               `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteUrl       string               `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string               `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string               `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string               `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string               `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string               `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string               `json:"supplier_phone,omitempty"`            // optional
	SupplierUrl          string               `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool                 `json:"is_purchased,string,omitempty"`       // optional
	Notes                string               `json:"notes,omitempty"`                     // optional
	Restrictions         string               `json:"restrictions,omitempty"`              // optional
	TotalCost            string               `json:"total_cost,omitempty"`                // optional
	Segment              RailSegmentPtrVector `json:"Segment,omitempty"`
	Traveler             TravelerPtrVector    `json:"Traveler,omitempty"` // optional
}

// BookingTime returns a time.Time object for BookingDate.
// Note: This won't have proper time zone information.
func (r *RailObject) BookingTime() (time.Time, error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// RailSegment contains details about an indivual train ride.
type RailSegment struct {
	StartDateTime       *DateTime `json:"StartDateTime,omitempty"`       // optional
	EndDateTime         *DateTime `json:"EndDateTime,omitempty"`         // optional
	StartStationAddress *Address  `json:"StartStationAddress,omitempty"` // optional
	EndStationAddress   *Address  `json:"EndStationAddress,omitempty"`   // optional
	StartStationName    string    `json:"start_station_name,omitempty"`  // optional
	EndStationName      string    `json:"end_station_name,omitempty"`    // optional
	CarrierName         string    `json:"carrier_name,omitempty"`        // optional
	CoachNumber         string    `json:"coach_number,omitempty"`        // optional
	ConfirmationNum     string    `json:"confirmation_num,omitempty"`    // optional
	Seats               string    `json:"seats,omitempty"`               // optional
	ServiceClass        string    `json:"service_class,omitempty"`       // optional
	TrainNumber         string    `json:"train_number,omitempty"`        // optional
	TrainType           string    `json:"train_type,omitempty"`          // optional
	Id                  string    `json:"id,omitempty"`                  // optional, read-only
}

// Transport Detail Types
const (
	TransportDetailTypeFerry                = "F"
	TransportDetailTypeGroundTransportation = "G"
)

// TransportObject contains details about other forms of transport like bus rides.
type TransportObject struct {
	Id                   string                    `json:"id,omitempty"`                        // optional, read-only
	TripId               string                    `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler     bool                      `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl          string                    `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string                    `json:"display_name,omitempty"`              // optional
	Image                ImagePtrVector            `json:"Image,omitempty"`                     // optional
	CancellationDateTime *DateTime                 `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string                    `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string                    `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string                    `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string                    `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string                    `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteUrl       string                    `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string                    `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string                    `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string                    `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string                    `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string                    `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string                    `json:"supplier_phone,omitempty"`            // optional
	SupplierUrl          string                    `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool                      `json:"is_purchased,string,omitempty"`       // optional
	Notes                string                    `json:"notes,omitempty"`                     // optional
	Restrictions         string                    `json:"restrictions,omitempty"`              // optional
	TotalCost            string                    `json:"total_cost,omitempty"`                // optional
	Segment              TransportSegmentPtrVector `json:"Segment,omitempty"`
	Traveler             TravelerPtrVector         `json:"Traveler,omitempty"` // optional
}

// BookingTime returns a time.Time object for BookingDate.
// Note: This won't have proper time zone information
func (r *TransportObject) BookingTime() (time.Time, error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// TransportSegment contains details about indivual transport rides.
type TransportSegment struct {
	StartDateTime        *DateTime `json:"StartDateTime,omitempty"`        // optional
	EndDateTime          *DateTime `json:"EndDateTime,omitempty"`          // optional
	StartLocationAddress *Address  `json:"StartLocationAddress,omitempty"` // optional
	EndLocationAddress   *Address  `json:"EndLocationAddress,omitempty"`   // optional
	StartLocationName    string    `json:"start_location_name,omitempty"`  // optional
	EndLocationName      string    `json:"end_location_name,omitempty"`    // optional
	DetailTypeCode       string    `json:"detail_type_code,omitempty"`     // optional
	CarrierName          string    `json:"carrier_name,omitempty"`         // optional
	ConfirmationNum      string    `json:"confirmation_num,omitempty"`     // optional
	NumberPassengers     string    `json:"number_passengers,omitempty"`    // optional
	VehicleDescription   string    `json:"vehicle_description,omitempty"`  // optional
	Id                   string    `json:"id,omitempty"`                   // optional, read-only
}

// Cruise Detail Types
const (
	CruiseDetailTypePortOfCall = "P"
)

// CruiseObject contains information about cruises.
type CruiseObject struct {
	Id                   string                 `json:"id,omitempty"`                        // optional, read-only
	TripId               string                 `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler     bool                   `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl          string                 `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string                 `json:"display_name,omitempty"`              // optional
	Image                ImagePtrVector         `json:"Image,omitempty"`                     // optional
	CancellationDateTime *DateTime              `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string                 `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string                 `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string                 `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string                 `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string                 `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteUrl       string                 `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string                 `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string                 `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string                 `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string                 `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string                 `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string                 `json:"supplier_phone,omitempty"`            // optional
	SupplierUrl          string                 `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool                   `json:"is_purchased,string,omitempty"`       // optional
	Notes                string                 `json:"notes,omitempty"`                     // optional
	Restrictions         string                 `json:"restrictions,omitempty"`              // optional
	TotalCost            string                 `json:"total_cost,omitempty"`                // optional
	Segment              CruiseSegmentPtrVector `json:"Segment,omitempty"`
	Traveler             TravelerPtrVector      `json:"Traveler,omitempty"`     // optional
	CabinNumber          string                 `json:"cabin_number,omitempty"` // optional
	CabinType            string                 `json:"cabin_type,omitempty"`   // optional
	Dining               string                 `json:"dining,omitempty"`       // optional
	ShipName             string                 `json:"ship_name,omitempty"`    // optional
}

// BookingTime returns a time.Time object for BookingDate.
// Note: This won't have proper time zone information.
func (r *CruiseObject) BookingTime() (time.Time, error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// CruiseSegment contains details about indivual cruise segments.
type CruiseSegment struct {
	StartDateTime   *DateTime `json:"StartDateTime,omitempty"`    // optional
	EndDateTime     *DateTime `json:"EndDateTime,omitempty"`      // optional
	LocationAddress *Address  `json:"LocationAddress,omitempty"`  // optional
	LocationName    string    `json:"location_name,omitempty"`    // optional
	DetailTypeCode  string    `json:"detail_type_code,omitempty"` // optional
	Id              string    `json:"id,omitempty"`               // optional, read-only
}

// RestaurantObject contains details about dining reservations.
// restaurant name should be in supplier_name.
// restaurant notes should be in notes.
type RestaurantObject struct {
	Id                   string         `json:"id,omitempty"`                        // optional, read-only
	TripId               string         `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler     bool           `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl          string         `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string         `json:"display_name,omitempty"`              // optional
	Image                ImagePtrVector `json:"Image,omitempty"`                     // optional
	CancellationDateTime *DateTime      `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string         `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string         `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string         `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string         `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string         `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteUrl       string         `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string         `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string         `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string         `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string         `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string         `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string         `json:"supplier_phone,omitempty"`            // optional
	SupplierUrl          string         `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool           `json:"is_purchased,string,omitempty"`       // optional
	Notes                string         `json:"notes,omitempty"`                     // optional
	Restrictions         string         `json:"restrictions,omitempty"`              // optional
	TotalCost            string         `json:"total_cost,omitempty"`                // optional
	DateTime             *DateTime      `json:"DateTime,omitempty"`                  // optional
	Address              *Address       `json:"Address,omitempty"`                   // optional
	ReservationHolder    *Traveler      `json:"ReservationHolder,omitempty"`         // optional
	Cuisine              string         `json:"cuisine,omitempty"`                   // optional
	DressCode            string         `json:"dress_code,omitempty"`                // optional
	Hours                string         `json:"hours,omitempty"`                     // optional
	NumberPatrons        string         `json:"number_patrons,omitempty"`            // optional
	PriceRange           string         `json:"price_range,omitempty"`               // optional
}

// BookingTime returns a time.Time object for BookingDate.
// Note: This won't have proper time zone information.
func (r *RestaurantObject) BookingTime() (time.Time, error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// Activity Detail Types
const (
	ActivityDetailTypeConcert = "C"
	ActivityDetailTypeTheatre = "H"
	ActivityDetailTypeMeeting = "M"
	ActivityDetailTypeTour    = "T"
)

// ActivityObject contains details about activities like museum, theatre, and other events.
type ActivityObject struct {
	Id                   string            `json:"id,omitempty"`                        // optional, read-only
	TripId               string            `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler     bool              `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl          string            `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string            `json:"display_name,omitempty"`              // optional
	Image                ImagePtrVector    `json:"Image,omitempty"`                     // optional
	CancellationDateTime *DateTime         `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string            `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string            `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string            `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string            `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string            `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteUrl       string            `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string            `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string            `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string            `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string            `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string            `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string            `json:"supplier_phone,omitempty"`            // optional
	SupplierUrl          string            `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool              `json:"is_purchased,string,omitempty"`       // optional
	Notes                string            `json:"notes,omitempty"`                     // optional
	Restrictions         string            `json:"restrictions,omitempty"`              // optional
	TotalCost            string            `json:"total_cost,omitempty"`                // optional
	StartDateTime        *DateTime         `json:"StartDateTime,omitempty"`             // optional
	EndTime              string            `json:"end_time,omitempty"`                  // optional, xs:time
	Address              *Address          `json:"Address,omitempty"`                   // optional
	Participant          TravelerPtrVector `json:"Participant,omitempty"`               // optional
	DetailTypeCode       string            `json:"detail_type_code,omitempty"`          // optional
	LocationName         string            `json:"location_name,omitempty"`             // optional
}

// BookingTime returns a time.Time object for BookingDate.
// Note: This won't have proper time zone information.
func (r *ActivityObject) BookingTime() (time.Time, error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// Note Detail Types
const (
	NoteDetailTypeArticle = "A"
)

// NoteObject contains information about notes added by the traveler.
type NoteObject struct {
	Id               string         `json:"id,omitempty"`                        // optional, read-only
	TripId           string         `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler bool           `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl      string         `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName      string         `json:"display_name,omitempty"`              // optional
	Image            ImagePtrVector `json:"Image,omitempty"`                     // optional
	DateTime         *DateTime      `json:"DateTime,omitempty"`                  // optional
	Address          *Address       `json:"Address,omitempty"`                   // optional
	DetailTypeCode   string         `json:"detail_type_code,omitempty"`          // optional
	Source           string         `json:"source,omitempty"`                    // optional
	Text             string         `json:"text,omitempty"`                      // optional
	Url              string         `json:"url,omitempty"`                       // optional
	Notes            string         `json:"notes,omitempty"`                     // optional
}

// MapObject contains addresses to show on a map.
type MapObject struct {
	Id               string         `json:"id,omitempty"`                        // optional, read-only
	TripId           string         `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler bool           `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl      string         `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName      string         `json:"display_name,omitempty"`              // optional
	Image            ImagePtrVector `json:"Image,omitempty"`                     // optional
	DateTime         *DateTime      `json:"DateTime,omitempty"`                  // optional
	Address          *Address       `json:"Address,omitempty"`                   // optional
}

// DirectionsObject contains addresses to show directions for on the trip.
type DirectionsObject struct {
	Id               string         `json:"id,omitempty"`                        // optional, read-only
	TripId           string         `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler bool           `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl      string         `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName      string         `json:"display_name,omitempty"`              // optional
	Image            ImagePtrVector `json:"Image,omitempty"`                     // optional
	DateTime         *DateTime      `json:"DateTime,omitempty"`                  // optional
	StartAddress     *Address       `json:"StartAddress,omitempty"`              // optional
	EndAddress       *Address       `json:"EndAddress,omitempty"`                // optional
}

// WeatherObject contains information about the weather at a particular destination.
// Weather is read-only.
type WeatherObject struct {
	Id                 string         `json:"id,omitempty"`                          // optional, read-only
	TripId             string         `json:"trip_id,omitempty"`                     // optional
	IsClientTraveler   bool           `json:"is_client_traveler,string,omitempty"`   // optional, read-only
	RelativeUrl        string         `json:"relative_url,omitempty"`                // optional, read-only
	DisplayName        string         `json:"display_name,omitempty"`                // optional
	Image              ImagePtrVector `json:"Image,omitempty"`                       // optional
	Date               string         `json:"date,omitempty"`                        // optional, read-only, xs:date
	Location           string         `json:"location,omitempty"`                    // optional, read-only
	AvgHighTempC       float64        `json:"avg_high_temp_c,string,omitempty"`      // optional, read-only
	AvgLowTempC        float64        `json:"avg_low_temp_c,string,omitempty"`       // optional, read-only
	AvgWindSpeedKn     float64        `json:"avg_wind_speed_kn,string,omitempty"`    // optional, read-only
	AvgPrecipitationCm float64        `json:"avg_precipitation_cm,string,omitempty"` // optional, read-only
	AvgSnowDepthCm     float64        `json:"avg_snow_depth_cm,string,omitempty"`    // optional, read-only
}

// Time returns a time.Time object for StartDate.
// Note: This won't have proper time zone information.
func (w *WeatherObject) Time() (time.Time, error) {
	return time.Parse("2006-01-02", w.Date)
}
