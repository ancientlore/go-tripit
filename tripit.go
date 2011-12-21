// The go-tripit library is a Go API library for accessing the TripIt service. The API supports
// two forms of authorization - simple web authorization and OAuth. The library uses TripIt's
// JSON interface, and has structs representing all of the TripIt types. Within these structs,
// elements ending in an underscore have access functions that set or get the value in a more
// pleasant form for use in Go programs.
package tripit

import (
	"time"
	"fmt"
	"os"
	"strconv"
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

// Request contains the objects that can be sent to TripIt in a request
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

// Error is returned from TripIt on error conditions
type Error struct {
	Code              int     `json:"code,string,omitempty"`                // read-only
	DetailedErrorCode float64 `json:"detailed_error_code,string,omitempty"` // optional, read-only
	Description       string  `json:"description,omitempty"`                // read-only
	EntityType        string  `json:"entity_type,omitempty"`                // read-only
	Timestamp         string  `json:"timestamp,omitempty"`                  // read-only, xs:datetime
}

// returns a time.Time object for the Timestamp
func (e *Error) Time() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, e.Timestamp)
}

// Returns a string containing the error information
func (e *Error) String() string {
	return fmt.Sprintf("TripIt Error %s: %s", e.Code, e.Description)
}

// Warning is returned from TripIt to indicate warning conditions
type Warning struct {
	Description string `json:"description,omitempty"` // read-only
	EntityType  string `json:"entity_type,omitempty"` // read-only
	Timestamp   string `json:"timestamp,omitempty"`   // read-only, xs:datetime
}

// returns a time.Time object for the Timestamp
func (w *Warning) Time() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, w.Timestamp)
}

// Returns a string containing the warning information
func (w *Warning) String() string {
	return fmt.Sprintf("TripIt Warning: %s", w.Description)
}

// Note that the Vectors are pointers - otherwise the JSON marshaler doesn't notice the custom methods

// Represents a TripIt API Response
type Response struct {
	Timestamp        string                    `json:"timestamp,omitempty"`
	NumBytes         int                       `json:"num_bytes,string,omitempty"`
	Error            ErrorVector               `json:"Error,omitempty"`            // optional
	Warning          WarningVector             `json:"Warning,omitempty"`          // optional
	Trip             TripPtrVector             `json:"Trip,omitempty"`             // optional
	ActivityObject   ActivityObjectPtrVector   `json:"ActivityObject,omitempty"`   // optional
	AirObject        AirObjectPtrVector        `json:"AirObject,omitempty"`        // optional
	CarObject        CarObjectPtrVector        `json:"CarObject,omitempty"`        // optional
	CruiseObject     CruiseObjectPtrVector     `json:"CruiseObject,omitempty"`     // optional
	DirectionsObject DirectionsObjectPtrVector `json:"DirectionsObject,omitempty"` // optional
	LodgingObject    LodgingObjectPtrVector    `json:"LodgingObject,omitempty"`    // optional
	MapObject        MapObjectPtrVector        `json:"MapObject,omitempty"`        // optional
	NoteObject       NoteObjectPtrVector       `json:"NoteObject,omitempty"`       // optional
	RailObject       RailObjectPtrVector       `json:"RailObject,omitempty"`       // optional
	RestaurantObject RestaurantObjectPtrVector `json:"RestaurantObject,omitempty"` // optional
	TransportObject  TransportObjectPtrVector  `json:"TransportObject,omitempty"`  // optional
	WeatherObject    WeatherObjectVector       `json:"WeatherObject,omitempty"`    // optional
	PointsProgram    PointsProgramVector       `json:"PointsProgram,omitempty"`    // optional
	Profile          ProfileVector             `json:"Profile,omitempty"`          // optional
	// @TODO need to add invitee stuff
}

// returns a time.Time object for the Timestamp
func (r *Response) Time() (*time.Time, os.Error) {
	t, err := strconv.Atoi64(r.Timestamp)
	if err != nil {
		return nil, err
	}
	return time.SecondsToUTC(t), nil
}

/*
   	For create, use either:
	- address for single-line addresses.
	- addr1, addr2, city, state, zip, and country for multi-line addresses.
	Multi-line address will be ignored if single-line address is present.
	See documentation for more information.
*/
type Address struct {
	Address   string  `json:"address,omitempty"`          // optional
	Addr1     string  `json:"addr1,omitempty"`            // optional
	Addr2     string  `json:"addr2,omitempty"`            // optional
	City      string  `json:"city,omitempty"`             // optional
	State     string  `json:"state,omitempty"`            // optional
	Zip       string  `json:"zip,omitempty"`              // optional
	Country   string  `json:"country,omitempty"`          // optional
	Latitude  float64 `json:"latitude,string,omitempty"`  // optional, read-only
	Longitude float64 `json:"longitude,string,omitempty"` // optional, read-only
}

// Information about a traveler
type Traveler struct {
	FirstName                string `json:"first_name,omitempty"`                 // optional
	MiddleName               string `json:"middle_name,omitempty"`                // optional
	LastName                 string `json:"last_name,omitempty"`                  // optional
	FrequentTravelerNum      string `json:"frequent_traveler_num,omitempty"`      // optional
	FrequentTravelerSupplier string `json:"frequent_traveler_supplier,omitempty"` // optional
	MealPreference           string `json:"meal_preference,omitempty"`            // optional
	SeatPreference           string `json:"seat_preference,omitempty"`            // optional
	TicketNum                string `json:"ticket_num,omitempty"`                 //optional
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

// All FlightStatus fields are read-only and only available for monitored TripIt Pro AirSegments
type FlightStatus struct {
	ScheduledDepartureDateTime *DateTime `json:"ScheduledDepartureDateTime,omitempty"`   // optional, read-only
	EstimatedDepartureDateTime *DateTime `json:"EstimatedDepartureDateTime,omitempty"`   // optional, read-only
	ScheduledArrivalDateTime   *DateTime `json:"ScheduledArrivalDateTime,omitempty"`     // optional, read-only
	EstimatedArrivalDateTime   *DateTime `json:"EstimatedArrivalDateTime,omitempty"`     // optional, read-only
	FlightStatus               int       `json:"flight_status,string,omitempty"`         // optional, read-only
	IsConnectionAtRisk         bool      `json:"is_connection_at_risk,string,omitempty"` // optional, read-only
	DepartureTerminal          string    `json:"departure_terminal,omitempty"`           // optional, read-only
	DepartureGate              string    `json:"departure_gate,omitempty"`               // optional, read-only
	ArrivalTerminal            string    `json:"arrival_terminal,omitempty"`             // optional, read-only
	ArrivalGate                string    `json:"arrival_gate,omitempty"`                 // optional, read-only
	LayoverMinutes             string    `json:"layover_minutes,omitempty"`              // optional, read-only
	BaggageClaim               string    `json:"baggage_claim,omitempty"`                // optional, read-only
	DivertedAirportCode        string    `json:"diverted_airport_code,omitempty"`        // optional, read-only
	LastModified               string    `json:"last_modified,omitempty"`                // read-only
}

// returns a time.Time object for LastModified
func (fs *FlightStatus) LastModifiedTime() (*time.Time, os.Error) {
	l, err := strconv.Atoi64(fs.LastModified)
	if err != nil {
		return nil, err
	}
	return time.SecondsToUTC(l), nil
}

// Information about images
type Image struct {
	Caption string `json:"caption,omitempty"` // optional
	Url     string `json:"url"`
}

// Stores date and time zone information, for example:
// {
//	 "date":"2009-11-10",
//   "time":"14:00:00",
//    "timezone":"America\/Los_Angeles",
//    "utc_offset":"-08:00"
// }
type DateTime struct {
	Date      string `json:"date,omitempty"`       // optional, xs:date
	Time      string `json:"time,omitempty"`       // optional, xs:time
	Timezone  string `json:"timezone,omitempty"`   // optional, read-only
	UtcOffset string `json:"utc_offset,omitempty"` // optional, read-only
}

// Convert to a time.Time
func (dt DateTime) GetTime() (*time.Time, os.Error) {
	if dt.UtcOffset == "" {
		return time.Parse(time.RFC3339, fmt.Sprintf("%sT%sZ", dt.Date, dt.Time))
	}
	return time.Parse(time.RFC3339, fmt.Sprintf("%sT%s%s", dt.Date, dt.Time, dt.UtcOffset))
}

// Sets the values of the DateTime strucure from a time.Time
func (dt *DateTime) SetTime(t *time.Time) {
	dt.Date = t.Format("2006-01-02")
	dt.Time = t.Format("15:04:05")
	dt.UtcOffset = t.Format("-07:00")
	dt.Timezone = t.Format("MST")
}

// PointsProgram contains information about tracked travel programs for TripIt Pro users.
// All PointsProgram elements are read-only
type PointsProgram struct {
	Id                  uint                          `json:"id,string,omitempty"`                    // read-only
	Name                string                        `json:"name,omitempty"`                         // optional, read-only
	AccountNumber       string                        `json:"account_number,omitempty"`               // optional, read-only
	AccountLogin        string                        `json:"account_login,omitempty"`                // optional, read-only
	Balance             string                        `json:"balance,omitempty"`                      // optional, read-only
	EliteStatus         string                        `json:"elite_status,omitempty"`                 // optional, read-only
	EliteNextStatus     string                        `json:"elite_next_status,omitempty"`            // optional, read-only
	EliteYtdQualify     string                        `json:"elite_ytd_qualify,omitempty"`            // optional, read-only
	EliteNeedToEarn     string                        `json:"elite_need_to_earn,omitempty"`           // optional, read-only
	LastModified        string                        `json:"last_modified,omitempty"`                // read-only
	TotalNumActivities  int                           `json:"total_num_activities,string,omitempty"`  // read-only
	TotalNumExpirations int                           `json:"total_num_expirations,string,omitempty"` // read-only
	ErrorMessage        string                        `json:"error_message,omitempty"`                // optional, read-only
	Activity            PointsProgramActivityVector   `json:"Activity,omitempty"`                     // optional, read-only
	Expiration          PointsProgramExpirationVector `json:"Expiration,omitempty"`                   // optional, read-only
}

// returns a time.Time object for LastModified
func (pp *PointsProgram) LastModifiedTime() (*time.Time, os.Error) {
	v, err := strconv.Atoi64(pp.LastModified)
	if err != nil {
		return nil, err
	}
	return time.SecondsToUTC(v), nil
}

// PointsProgramActivity contains program transactions
// All PointsProgramActivity elements are read-only
type PointsProgramActivity struct {
	Date        string `json:"date,omitempty"`        // read-only, xs:date
	Description string `json:"description,omitempty"` // optional, read-only
	Base        string `json:"base,omitempty"`        // optional, read-only
	Bonus       string `json:"bonus,omitempty"`       // optional, read-only
	Total       string `json:"total,omitempty"`       // optional, read-only
}

// returns a time.Time object for Date
// Note: This won't have proper time zone information
func (pa *PointsProgramActivity) Time() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", pa.Date)
}

// All PointsProgramExpiration elements are read-only
type PointsProgramExpiration struct {
	Date   string `json:"date,omitempty"`   // read-only, xs:date
	Amount string `json:"amount,omitempty"` // optional, read-only
}

// returns a time.Time object for Date
// Note: This won't have proper time zone information
func (pe *PointsProgramExpiration) Time() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", pe.Date)
}

// TripShare contains information about which users a trip is shared with
type TripShare struct {
	TripId            uint `json:"trip_id,string,omitempty"`
	IsTraveler        bool `json:"is_traveler,string,omitempty"`
	IsReadOnly        bool `json:"is_read_only,string,omitempty"`
	IsSentWithDetails bool `json:"is_sent_with_details,string,omitempty"`
}

// Connection request
type ConnectionRequest struct {

}

// Invitation contains a list of users invited to see the trip
type Invitation struct {
	EmailAddresses    []string           `json:"EmailAddresses,omitempty"`
	TripShare         *TripShare         `json:"TripShare,omitempty"`         // optional
	ConnectionRequest *ConnectionRequest `json:"ConnectionRequest,omitempty"` // optional
	Message           string             `json:"message,omitempty"`           // optional
}

// Profile contains user information
// All Profile elements are read-only
type Profile struct {
	Attributes            ProfileAttributes      `json:"_attributes"`                     // read-only
	ProfileEmailAddresses *ProfileEmailAddresses `json:"ProfileEmailAddresses,omitempty"` // optional, read-only
	GroupMemberships      *GroupMemberships      `json:"GroupMemberships,omitempty"`      // optional, read-only
	IsClient              bool                   `json:"is_client,string,omitempty"`      // read-only
	IsPro                 bool                   `json:"is_pro,string,omitempty"`         // read-only
	ScreenName            string                 `json:"screen_name,omitempty"`           // read-only
	PublicDisplayName     string                 `json:"public_display_name,omitempty"`   // read-only
	ProfileUrl            string                 `json:"profile_url,omitempty"`           // read-only
	HomeCity              string                 `json:"home_city,omitempty"`             // optional, read-only
	Company               string                 `json:"company,omitempty"`               // optional, read-only
	AboutMeInfo           string                 `json:"about_me_info,omitempty"`         // optional, read-only
	PhotoUrl              string                 `json:"photo_url,omitempty"`             // optional, read-only
	ActivityFeedUrl       string                 `json:"activity_feed_url,omitempty"`     // optional, read-only
	AlertsFeedUrl         string                 `json:"alerts_feed_url,omitempty"`       // optional, read-only
	IcalUrl               string                 `json:"ical_url,omitempty"`              // optional, read-only
}

// ProfileEmailAddresses contains the list of email addresses for a user
type ProfileEmailAddresses struct {
	ProfileEmailAddress ProfileEmailAddressVector `json:"ProfileEmailAddress,omitempty"`
}

// GroupMemberships contains a list of groups that the user is a member of
type GroupMemberships struct {
	Group GroupVector `json:"Group,omitempty"` // optional, read-only
}

// ProfileAttributes represent links to profiles
type ProfileAttributes struct {
	Ref string `json:"ref,omitempty"` // read-only
}

// ProfileEmailAddress contains an email address and its properties
// All ProfileEmailAddress elements are read-only
type ProfileEmailAddress struct {
	Address      string `json:"address"`                         // read-only
	IsAutoImport bool   `json:"is_auto_import,string,omitempty"` // read-only
	IsConfirmed  bool   `json:"is_confirmed,string,omitempty"`   // read-only
	IsPrimary    bool   `json:"is_primary,string,omitempty"`     // read-only
}

// Group contains data about a group in TripIt
// All Group elements are read-only
type Group struct {
	DisplayName string `json:"display_name,omitempty"` // read-only
	Url         string `json:"url"`                    // read-only
}

// Trip Invitee
// All Invitee elements are read-only
type Invitee struct {
	IsReadOnly bool              `json:"is_read_only,string,omitempty"` // read-only
	IsTraveler bool              `json:"is_traveler,string,omitempty"`  // read-only
	Attributes InviteeAttributes `json:"_attributes"`                   // read-only, Use the profile_ref attribute to reference a Profile
}

// Used to link to user profiles
type InviteeAttributes struct {
	ProfileRef string `json:"profile_ref"` // read-only, used to reference a profile
}

// A CRS remark
// All TripCrsRemark elements are read-only
type TripCrsRemark struct {
	RecordLocator string `json:"record_locator,omitempty"` // read-only
	Notes         string `json:"notes,omitempty"`          // read-only
}

// List of nearby users
// All ClosenessMatch elements are read-only
type ClosenessMatch struct {
	Attributes ClosenessMatchAttributes `json:"_attributes"` // read-only, Use the profile_ref attribute to reference a Profile
}

// Links to profiles of nearby users
type ClosenessMatchAttributes struct {
	ProfileRef string `json:"profile_ref"` // read-only, Use the profile_ref attribute to reference a Profile
}

// The Trip
type Trip struct {
	ClosenessMatches       *ClosenessMatches `json:"ClosenessMatches,omitempty"`         // optional, ClosenessMatches are read-only
	TripInvitees           *TripInvitees     `json:"TripInvitees,omitempty"`             // optional, TripInvitees are read-only
	TripCrsRemarks         *TripCrsRemarks   `json:"TripCrsRemarks,omitempty"`           // optional, TripCrsRemarks are read-only
	Id                     uint              `json:"id,string,omitempty"`                // optional, id is a read-only field
	RelativeUrl            string            `json:"relative_url,omitempty"`             // optional, relative_url is a read-only field
	StartDate              string            `json:"start_date,omitempty"`               // optional, xs:date
	EndDate                string            `json:"end_date,omitempty"`                 // optional, xs:date
	Description            string            `json:"description,omitempty"`              // optional
	DisplayName            string            `json:"display_name,omitempty"`             // optional
	ImageUrl               string            `json:"image_url,omitempty"`                // optional
	IsPrivate              bool              `json:"is_private,string,omitempty"`        // optional
	PrimaryLocation        string            `json:"primary_location,omitempty"`         // optional
	PrimaryLocationAddress *Address          `json:"primary_location_address,omitempty"` // optional, PrimaryLocationAddress is a read-only field
}

// People invited to view a trip
type TripInvitees struct {
	Invitee InviteeVector `json:"Invitee,omitempty"` // optional, TripInvitees are read-only
}

// TripIt users who are near this trip
type ClosenessMatches struct {
	ClosenessMatch ClosenessMatchVector `json:"Match,omitempty"` // optional, ClosenessMatches are read-only
}

// Remarks from a reservation system
type TripCrsRemarks struct {
	TripCrsRemark TripCrsRemarkVector `json:"TripCrsRemark,omitempty"` // optional, TripCrsRemarks are read-only
}

// returns a time.Time object for StartDate
// Note: This won't have proper time zone information
func (t *Trip) StartTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", t.StartDate)
}

// returns a time.Time object for EndDate
// Note: This won't have proper time zone information
func (t *Trip) EndTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", t.EndDate)
}

// AirObject contains data about a flight
type AirObject struct {
	Id                   uint                `json:"id,string,omitempty"`                 // optional, read-only
	TripId               uint                `json:"trip_id,string,omitempty"`            // optional
	IsClientTraveler     bool                `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl          string              `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string              `json:"display_name,omitempty"`              // optional
	Image                ImagePtrVector      `json:"Image,omitempty"`                     // optional
	CancellationDateTime *DateTime           `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string              `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string              `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string              `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string              `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string              `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteUrl       string              `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string              `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string              `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string              `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string              `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string              `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string              `json:"supplier_phone,omitempty"`            // optional
	SupplierUrl          string              `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool                `json:"is_purchased,string,omitempty"`       // optional
	Notes                string              `json:"notes,omitempty"`                     // optional
	Restrictions         string              `json:"restrictions,omitempty"`              // optional
	TotalCost            string              `json:"total_cost,omitempty"`                // optional
	Segment              AirSegmentPtrVector `json:"Segment,omitempty"`
	Traveler             TravelerPtrVector   `json:"Traveler,omitempty"` // optional
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *AirObject) BookingTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// AirSegment contains details about individual flights
type AirSegment struct {
	Status                *FlightStatus `json:"Status,omitempty"`                         // optional
	StartDateTime         *DateTime     `json:"StartDateTime,omitempty"`                  // optional
	EndDateTime           *DateTime     `json:"EndDateTime,omitempty"`                    // optional
	StartAirportCode      string        `json:"start_airport_code,omitempty"`             // optional
	StartAirportLatitude  float64       `json:"start_airport_latitude,string,omitempty"`  // optional, read-only
	StartAirportLongitude float64       `json:"start_airport_longitude,string,omitempty"` // optional, read-only
	StartCityName         string        `json:"start_city_name,omitempty"`                // optional
	StartGate             string        `json:"start_gate,omitempty"`                     // optional
	StartTerminal         string        `json:"start_terminal,omitempty"`                 // optional
	EndAirportCode        string        `json:"end_airport_code,omitempty"`               // optional
	EndAirportLatitude    float64       `json:"end_airport_latitude,string,omitempty"`    // optional, read-only
	EndAirportLongitude   float64       `json:"end_airport_longitude,string,omitempty"`   // optional, read-only
	EndCityName           string        `json:"end_city_name,omitempty"`                  // optional
	EndGate               string        `json:"end_gate,omitempty"`                       // optional
	EndTerminal           string        `json:"end_terminal,omitempty"`                   // optional
	MarketingAirline      string        `json:"marketing_airline,omitempty"`              // optional
	MarketingAirlineCode  string        `json:"marketing_airline_code,omitempty"`         // optional, read-only
	MarketingFlightNumber string        `json:"marketing_flight_number,omitempty"`        // optional
	OperatingAirline      string        `json:"operating_airline,omitempty"`              // optional
	OperatingAirlineCode  string        `json:"operating_airline_code,omitempty"`         // optional, read-only
	OperatingFlightNumber string        `json:"operating_flight_number,omitempty"`        // optional
	AlternativeFlightsUrl string        `json:"alternate_flights_url,omitempty"`          // optional, read-only
	Aircraft              string        `json:"aircraft,omitempty"`                       // optional
	AircraftDisplayName   string        `json:"aircraft_display_name,omitempty"`          // optional, read-only
	Distance              string        `json:"distance,omitempty"`                       // optional
	Duration              string        `json:"duration,omitempty"`                       // optional
	Entertainment         string        `json:"entertainment,omitempty"`                  // optional
	Meal                  string        `json:"meal,omitempty"`                           // optional
	Notes                 string        `json:"notes,omitempty"`                          // optional
	OntimePerc            string        `json:"ontime_perc,omitempty"`                    // optional
	Seats                 string        `json:"seats,omitempty"`                          // optional
	ServiceClass          string        `json:"service_class,omitempty"`                  // optional
	Stops                 string        `json:"stops,omitempty"`                          // optional
	BaggageClaim          string        `json:"baggage_claim,omitempty"`                  // optional
	CheckInUrl            string        `json:"check_in_url,omitempty"`                   // optional
	ConflictResolutionUrl string        `json:"conflict_resolution_url,omitempty"`        // optional, read-only
	IsHidden              bool          `json:"is_hidden,string,omitempty"`               // optional, read-only
	Id                    uint          `json:"id,string,omitempty"`                      // optional, read-only
}

// LodgingObject contains information about hotels or other lodging
// hotel cancellation remarks should be in restrictions
// hotel room description should be in notes
// hotel average daily rate should be in booking_rate
type LodgingObject struct {
	Id                   uint              `json:"id,string,omitempty"`                 // optional, read-only
	TripId               uint              `json:"trip_id,string,omitempty"`            // optional
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
	EndDateTime          *DateTime         `json:"EndDateTime,omitempty"`               // optional
	Address              *Address          `json:"Address,omitempty"`                   // optional
	Guest                TravelerPtrVector `json:"Guest,omitempty"`                     // optional
	NumberGuests         string            `json:"number_guests,omitempty"`             // optional
	NumberRooms          string            `json:"numer_rooms,omitempty"`               // optional
	RoomType             string            `json:"room_type,omitempty"`                 // optional
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *LodgingObject) BookingTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// CarObject contains information about rental cars
// car cancellation remarks should be in restrictions
// car pickup instructions should be in notes
// car daily rate should be in booking_rate
type CarObject struct {
	Id                   uint              `json:"id,string,omitempty"`                 // optional, read-only
	TripId               uint              `json:"trip_id,string,omitempty"`            // optional
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
	EndDateTime          *DateTime         `json:"EndDateTime,omitempty"`               // optional
	StartLocationAddress *Address          `json:"StartLocationAddress,omitempty"`      // optional
	EndLocationAddress   *Address          `json:"EndLocationAddress,omitempty"`        // optional
	Driver               TravelerPtrVector `json:"Driver,omitempty"`                    // optional
	StartLocationHours   string            `json:"start_location_hours,omitempty"`      // optional
	StartLocationName    string            `json:"start_location_name,omitempty"`       // optional
	StartLocationPhone   string            `json:"start_location_phone,omitempty"`      // optional
	EndLocationHours     string            `json:"end_location_hours,omitempty"`        // optional
	EndLocationName      string            `json:"end_location_name,omitempty"`         // optional
	EndLocationPhone     string            `json:"end_location_phone,omitempty"`        // optional
	CarDescription       string            `json:"car_description,omitempty"`           // optional
	CarType              string            `json:"car_type,omitempty"`                  // optional
	MileageCharges       string            `json:"mileage_charges,omitempty"`           // optional
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *CarObject) BookingTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// RailObject contains information about trains
type RailObject struct {
	Id                   uint                 `json:"id,string,omitempty"`                 // optional, read-only
	TripId               uint                 `json:"trip_id,string,omitempty"`            // optional
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

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *RailObject) BookingTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// RailSegment contains details about an indivual train ride
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
	Id                  uint      `json:"id,string,omitempty"`           // optional, read-only
}

// Transport Detail Types
const (
	TransportDetailTypeFerry                = "F"
	TransportDetailTypeGroundTransportation = "G"
)

// TransportObject contains details about other forms of transport like bus rides
type TransportObject struct {
	Id                   uint                      `json:"id,string,omitempty"`                 // optional, read-only
	TripId               uint                      `json:"trip_id,string,omitempty"`            // optional
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

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *TransportObject) BookingTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// TransportSegment contains details about indivual transport rides
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
	Id                   uint      `json:"id,string,omitempty"`            // optional, read-only
}

// Cruise Detail Types
const (
	CruiseDetailTypePortOfCall = "P"
)

// CruiseObject contains information about cruises
type CruiseObject struct {
	Id                   uint                   `json:"id,string,omitempty"`                 // optional, read-only
	TripId               uint                   `json:"trip_id,string,omitempty"`            // optional
	IsClientTraveler     bool                   `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl          string                 `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string                 `json:"display_name,omitempty"`              // optional
	Image                ImagePtrVector         `json:"Image,omitempty"`                     // optional
	CancellationDateTime *DateTime              `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string                 `json:"booking_date,omitempty`               // optional, xs:date
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

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *CruiseObject) BookingTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// CruiseSegment contains details about indivual cruise segments
type CruiseSegment struct {
	StartDateTime   *DateTime `json:"StartDateTime,omitempty"`    // optional
	EndDateTime     *DateTime `json:"EndDateTime,omitempty"`      // optional
	LocationAddress *Address  `json:"LocationAddress,omitempty"`  // optional
	LocationName    string    `json:"location_name,omitempty"`    // optional
	DetailTypeCode  string    `json:"detail_type_code,omitempty"` // optional
	Id              uint      `json:"id,string,omitempty"`        // optional, read-only
}

// RestaurantObject contains details about dining reservations
// restaurant name should be in supplier_name
// restaurant notes should be in notes
type RestaurantObject struct {
	Id                   uint           `json:"id,string,omitempty"`                 // optional, read-only
	TripId               uint           `json:"trip_id,string,omitempty"`            // optional
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

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *RestaurantObject) BookingTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// Activity Detail Types
const (
	ActivityDetailTypeConcert = "C"
	ActivityDetailTypeTheatre = "H"
	ActivityDetailTypeMeeting = "M"
	ActivityDetailTypeTour    = "T"
)

// ActivityObject contains details about activities like museum, theatre, and other events
type ActivityObject struct {
	Id                   uint              `json:"id,string,omitempty"`                 // optional, read-only
	TripId               uint              `json:"trip_id,string,omitempty"`            // optional
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

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *ActivityObject) BookingTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

// Note Detail Types
const (
	NoteDetailTypeArticle = "A"
)

// NoteObject contains information about notes added by the traveler
type NoteObject struct {
	Id               uint           `json:"id,string,omitempty"`                 // optional, read-only
	TripId           uint           `json:"trip_id,string,omitempty"`            // optional
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

// MapObject contains addresses to show on a map
type MapObject struct {
	Id               uint           `json:"id,string,omitempty"`                 // optional, read-only
	TripId           uint           `json:"trip_id,string,omitempty"`            // optional
	IsClientTraveler bool           `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl      string         `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName      string         `json:"display_name,omitempty"`              // optional
	Image            ImagePtrVector `json:"Image,omitempty"`                     // optional
	DateTime         *DateTime      `json:"DateTime,omitempty"`                  // optional
	Address          *Address       `json:"Address,omitempty"`                   // optional
}

// DirectionsObject contains addresses to show directions for on the trip
type DirectionsObject struct {
	Id               uint           `json:"id,string,omitempty"`                 // optional, read-only
	TripId           uint           `json:"trip_id,string,omitempty"`            // optional
	IsClientTraveler bool           `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeUrl      string         `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName      string         `json:"display_name,omitempty"`              // optional
	Image            ImagePtrVector `json:"Image,omitempty"`                     // optional
	DateTime         *DateTime      `json:"DateTime,omitempty"`                  // optional
	StartAddress     *Address       `json:"StartAddress,omitempty"`              // optional
	EndAddress       *Address       `json:"EndAddress,omitempty"`                // optional
}

// WeatherObject contains information about the weather at a particular destination
// Weather is read-only
type WeatherObject struct {
	Id                 uint           `json:"id,string,omitempty"`                   // optional, read-only
	TripId             uint           `json:"trip_id,string,omitempty"`              // optional
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

// returns a time.Time object for StartDate
// Note: This won't have proper time zone information
func (w *WeatherObject) Time() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", w.Date)
}
