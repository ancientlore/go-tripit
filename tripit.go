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
	Code_              string  `json:"code"`                // read-only
	DetailedErrorCode_ *string `json:"detailed_error_code"` // optional, read-only
	Description        string  `json:"description"`         // read-only
	EntityType         string  `json:"entity_type"`         // read-only
	Timestamp          string  `json:"timestamp"`           // read-only, xs:datetime
}

// Returns error code
func (e *Error) Code() (int, os.Error) {
	return strconv.Atoi(e.Code_)
}

// Returns detailed error code
func (e *Error) DetailedErrorCode() (float64, os.Error) {
	if e.DetailedErrorCode_ == nil {
		return 0.0, nil
	}
	return strconv.Atof64(*e.DetailedErrorCode_)
}

// returns a time.Time object for the Timestamp
func (e *Error) Time() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, e.Timestamp)
}

// Returns a string containing the error information
func (e *Error) String() string {
	return fmt.Sprintf("TripIt Error %s: %s", e.Code_, e.Description)
}

// Warning is returned from TripIt to indicate warning conditions
type Warning struct {
	Description string `json:"description"` // read-only
	EntityType  string `json:"entity_type"` // read-only
	Timestamp   string `json:"timestamp"`   // read-only, xs:datetime
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
	Timestamp_       string                     `json:"timestamp"`
	NumBytes_        string                     `json:"num_bytes"`
	Error            *ErrorVector               `json:"Error,omitempty"`            // optional
	Warning          *WarningVector             `json:"Warning,omitempty"`          // optional
	Trip             *TripPtrVector             `json:"Trip,omitempty"`             // optional
	ActivityObject   *ActivityObjectPtrVector   `json:"ActivityObject,omitempty"`   // optional
	AirObject        *AirObjectPtrVector        `json:"AirObject,omitempty"`        // optional
	CarObject        *CarObjectPtrVector        `json:"CarObject,omitempty"`        // optional
	CruiseObject     *CruiseObjectPtrVector     `json:"CruiseObject,omitempty"`     // optional
	DirectionsObject *DirectionsObjectPtrVector `json:"DirectionsObject,omitempty"` // optional
	LodgingObject    *LodgingObjectPtrVector    `json:"LodgingObject,omitempty"`    // optional
	MapObject        *MapObjectPtrVector        `json:"MapObject,omitempty"`        // optional
	NoteObject       *NoteObjectPtrVector       `json:"NoteObject,omitempty"`       // optional
	RailObject       *RailObjectPtrVector       `json:"RailObject,omitempty"`       // optional
	RestaurantObject *RestaurantObjectPtrVector `json:"RestaurantObject,omitempty"` // optional
	TransportObject  *TransportObjectPtrVector  `json:"TransportObject,omitempty"`  // optional
	WeatherObject    *WeatherObjectVector       `json:"WeatherObject,omitempty"`    // optional
	PointsProgram    *PointsProgramVector       `json:"PointsProgram,omitempty"`    // optional
	Profile          *ProfileVector             `json:"Profile,omitempty"`          // optional
	// @TODO need to add invitee stuff
}

// returns a time.Time object for the Timestamp
func (r *Response) Timestamp() (*time.Time, os.Error) {
	t, err := strconv.Atoi64(r.Timestamp_)
	if err != nil {
		return nil, err
	}
	return time.SecondsToUTC(t), nil
}

// Returns number of bytes in response
func (r *Response) NumBytes() (int, os.Error) {
	return strconv.Atoi(r.NumBytes_)
}

/*
   	For create, use either:
	- address for single-line addresses.
	- addr1, addr2, city, state, zip, and country for multi-line addresses.
	Multi-line address will be ignored if single-line address is present.
	See documentation for more information.
*/
type Address struct {
	Address    string  `json:"address"`   // optional
	Addr1      string  `json:"addr1"`     // optional
	Addr2      string  `json:"addr2"`     // optional
	City       string  `json:"city"`      // optional
	State      string  `json:"state"`     // optional
	Zip        string  `json:"zip"`       // optional
	Country    string  `json:"country"`   // optional
	Latitude_  *string `json:"latitude"`  // optional, read-only
	Longitude_ *string `json:"longitude"` // optional, read-only
}

// Return the latitude of the address
func (a *Address) Latitude() (float64, os.Error) {
	if a.Latitude_ == nil {
		return 0.0, os.NewError("Latitude not provided")
	}
	return strconv.Atof64(*a.Latitude_)
}

// Return the longitude of the address
func (a *Address) Longitude() (float64, os.Error) {
	if a.Longitude_ == nil {
		return 0.0, os.NewError("Longitude not provided")
	}
	return strconv.Atof64(*a.Longitude_)
}

// Information about a traveler
type Traveler struct {
	FirstName                string `json:"first_name"`                 // optional
	MiddleName               string `json:"middle_name"`                // optional
	LastName                 string `json:"last_name"`                  // optional
	FrequentTravelerNum      string `json:"frequent_traveler_num"`      // optional
	FrequentTravelerSupplier string `json:"frequent_traveler_supplier"` // optional
	MealPreference           string `json:"meal_preference"`            // optional
	SeatPreference           string `json:"seat_preference"`            // optional
	TicketNum                string `json:"ticket_num"`                 //optional
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
	ScheduledDepartureDateTime *DateTime `json:"ScheduledDepartureDateTime"` // optional, read-only
	EstimatedDepartureDateTime *DateTime `json:"EstimatedDepartureDateTime"` // optional, read-only
	ScheduledArrivalDateTime   *DateTime `json:"ScheduledArrivalDateTime"`   // optional, read-only
	EstimatedArrivalDateTime   *DateTime `json:"EstimatedArrivalDateTime"`   // optional, read-only
	FlightStatus_              *string   `json:"flight_status"`              // optional, read-only
	IsConnectionAtRisk_        *string   `json:"is_connection_at_risk"`      // optional, read-only
	DepartureTerminal          string    `json:"departure_terminal"`         // optional, read-only
	DepartureGate              string    `json:"departure_gate"`             // optional, read-only
	ArrivalTerminal            string    `json:"arrival_terminal"`           // optional, read-only
	ArrivalGate                string    `json:"arrival_gate"`               // optional, read-only
	LayoverMinutes             string    `json:"layover_minutes"`            // optional, read-only
	BaggageClaim               string    `json:"baggage_claim"`              // optional, read-only
	DivertedAirportCode        string    `json:"diverted_airport_code"`      // optional, read-only
	LastModified_              string    `json:"last_modified"`              // read-only
}

// Returns the flight status as an int
func (fs *FlightStatus) FlightStatus() (int, os.Error) {
	if fs.FlightStatus_ == nil {
		return 0, os.NewError("Flight status not specified")
	}
	return strconv.Atoi(*fs.FlightStatus_)
}

// Returns whether the connection is at risk as a boolean
func (fs *FlightStatus) IsConnectionAtRisk() (bool, os.Error) {
	if fs.IsConnectionAtRisk_ == nil {
		return false, os.NewError("Is connection at risk not specified")
	}
	return strconv.Atob(*fs.IsConnectionAtRisk_)
}

// returns a time.Time object for LastModified
func (fs *FlightStatus) LastModified() (*time.Time, os.Error) {
	l, err := strconv.Atoi64(fs.LastModified_)
	if err != nil {
		return nil, err
	}
	return time.SecondsToUTC(l), nil
}

// Information about images
type Image struct {
	Caption string `json:"caption"` // optional
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
	Date_     string `json:"date"`       // optional, xs:date
	Time_     string `json:"time"`       // optional, xs:time
	Timezone  string `json:"timezone"`   // optional, read-only
	UtcOffset string `json:"utc_offset"` // optional, read-only
}

// Convert to a time.Time
func (dt DateTime) DateTime() (*time.Time, os.Error) {
	if dt.UtcOffset == "" {
		return time.Parse(time.RFC3339, fmt.Sprintf("%sT%sZ", dt.Date_, dt.Time_))
	}
	return time.Parse(time.RFC3339, fmt.Sprintf("%sT%s%s", dt.Date_, dt.Time_, dt.UtcOffset))
}

// Sets the values of the DateTime strucure from a time.Time
func (dt *DateTime) SetDateTime(t *time.Time) {
	dt.Date_ = t.Format("2006-01-02")
	dt.Time_ = t.Format("15:04:05")
	dt.UtcOffset = t.Format("-07:00")
	dt.Timezone = t.Format("MST")
}

// PointsProgram contains information about tracked travel programs for TripIt Pro users.
// All PointsProgram elements are read-only
type PointsProgram struct {
	Id_                  string                         `json:"id"`                    // read-only
	Name                 string                         `json:"name"`                  // optional, read-only
	AccountNumber        string                         `json:"account_number"`        // optional, read-only
	AccountLogin         string                         `json:"account_login"`         // optional, read-only
	Balance              string                         `json:"balance"`               // optional, read-only
	EliteStatus          string                         `json:"elite_status"`          // optional, read-only
	EliteNextStatus      string                         `json:"elite_next_status"`     // optional, read-only
	EliteYtdQualify      string                         `json:"elite_ytd_qualify"`     // optional, read-only
	EliteNeedToEarn      string                         `json:"elite_need_to_earn"`    // optional, read-only
	LastModified_        string                         `json:"last_modified"`         // read-only
	TotalNumActivities_  string                         `json:"total_num_activities"`  // read-only
	TotalNumExpirations_ string                         `json:"total_num_expirations"` // read-only
	ErrorMessage         string                         `json:"error_message"`         // optional, read-only
	Activity             *PointsProgramActivityVector   `json:"Activity"`              // optional, read-only
	Expiration           *PointsProgramExpirationVector `json:"Expiration"`            // optional, read-only
}

// Returns the ID
func (pp *PointsProgram) Id() (uint, os.Error) {
	return strconv.Atoui(pp.Id_)
}

// Returns the total number of activities
func (pp *PointsProgram) TotalNumActivities() (int, os.Error) {
	return strconv.Atoi(pp.TotalNumActivities_)
}

// Returns the total number of expirations
func (pp *PointsProgram) TotalNumExpirations() (int, os.Error) {
	return strconv.Atoi(pp.TotalNumExpirations_)
}

// returns a time.Time object for LastModified
func (pp *PointsProgram) LastModified() (*time.Time, os.Error) {
	v, err := strconv.Atoi64(pp.LastModified_)
	if err != nil {
		return nil, err
	}
	return time.SecondsToUTC(v), nil
}

// PointsProgramActivity contains program transactions
// All PointsProgramActivity elements are read-only
type PointsProgramActivity struct {
	Date_       string `json:"date"`        // read-only, xs:date
	Description string `json:"description"` // optional, read-only
	Base        string `json:"base"`        // optional, read-only
	Bonus       string `json:"bonus"`       // optional, read-only
	Total       string `json:"total"`       // optional, read-only
}

// returns a time.Time object for Date
// Note: This won't have proper time zone information
func (pa *PointsProgramActivity) Date() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", pa.Date_)
}

// All PointsProgramExpiration elements are read-only
type PointsProgramExpiration struct {
	Date_  string `json:"date"`   // read-only, xs:date
	Amount string `json:"amount"` // optional, read-only
}

// returns a time.Time object for Date
// Note: This won't have proper time zone information
func (pe *PointsProgramExpiration) Date() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", pe.Date_)
}

// TripShare contains information about which users a trip is shared with
type TripShare struct {
	TripId_            string `json:"trip_id"`
	IsTraveler_        string `json:"is_traveler"`
	IsReadOnly_        string `json:"is_read_only"`
	IsSentWithDetails_ string `json:"is_sent_with_details"`
}

// Returns the TripId
func (ts *TripShare) TripId() (uint, os.Error) {
	return strconv.Atoui(ts.TripId_)
}

// Returns a boolean indicating whether this record is for a traveler
func (ts *TripShare) IsTraveler() (bool, os.Error) {
	return strconv.Atob(ts.IsTraveler_)
}

// Returns a boolean indicating whether the item is read-only
func (ts *TripShare) IsReadOnly() (bool, os.Error) {
	return strconv.Atob(ts.IsReadOnly_)
}

// Returns a boolean for is sent with details
func (ts *TripShare) IsSentWithDetails() (bool, os.Error) {
	return strconv.Atob(ts.IsSentWithDetails_)
}

// Connection request
type ConnectionRequest struct {

}

// Invitation contains a list of users invited to see the trip
type Invitation struct {
	EmailAddresses    []string           `json:"EmailAddresses"`
	TripShare         *TripShare         `json:"TripShare"`         // optional
	ConnectionRequest *ConnectionRequest `json:"ConnectionRequest"` // optional
	Message           string             `json:"message"`           // optional
}

// Profile contains user information
// All Profile elements are read-only
type Profile struct {
	Attributes            ProfileAttributes      `json:"_attributes"`           // read-only
	ProfileEmailAddresses *ProfileEmailAddresses `json:"ProfileEmailAddresses"` // optional, read-only
	GroupMemberships      *GroupMemberships      `json:"GroupMemberships"`      // optional, read-only
	IsClient_             string                 `json:"is_client"`             // read-only
	IsPro_                string                 `json:"is_pro"`                // read-only
	ScreenName            string                 `json:"screen_name"`           // read-only
	PublicDisplayName     string                 `json:"public_display_name"`   // read-only
	ProfileUrl            string                 `json:"profile_url"`           // read-only
	HomeCity              string                 `json:"home_city"`             // optional, read-only
	Company               string                 `json:"company"`               // optional, read-only
	AboutMeInfo           string                 `json:"about_me_info"`         // optional, read-only
	PhotoUrl              string                 `json:"photo_url"`             // optional, read-only
	ActivityFeedUrl       string                 `json:"activity_feed_url"`     // optional, read-only
	AlertsFeedUrl         string                 `json:"alerts_feed_url"`       // optional, read-only
	IcalUrl               string                 `json:"ical_url"`              // optional, read-only
}

// ProfileEmailAddresses contains the list of email addresses for a user
type ProfileEmailAddresses struct {
	ProfileEmailAddress *ProfileEmailAddressVector `json:"ProfileEmailAddress"`
}

// GroupMemberships contains a list of groups that the user is a member of
type GroupMemberships struct {
	Group *GroupVector `json:"Group"` // optional, read-only
}

// ProfileAttributes represent links to profiles
type ProfileAttributes struct {
	Ref string `json:"ref"` // read-only
}

// Returns whether the profile is a client
func (p *Profile) IsClient() (bool, os.Error) {
	return strconv.Atob(p.IsClient_)
}

// Returns whether the profile has TripIt pro
func (p *Profile) IsPro() (bool, os.Error) {
	return strconv.Atob(p.IsPro_)
}

// ProfileEmailAddress contains an email address and its properties
// All ProfileEmailAddress elements are read-only
type ProfileEmailAddress struct {
	Address       string `json:"address"`        // read-only
	IsAutoImport_ string `json:"is_auto_import"` // read-only
	IsConfirmed_  string `json:"is_confirmed"`   // read-only
	IsPrimary_    string `json:"is_primary"`     // read-only
}

// Returns whether the email address 
func (e *ProfileEmailAddress) IsAutoImport() (bool, os.Error) {
	return strconv.Atob(e.IsAutoImport_)
}

// Returns whether the email address 
func (e *ProfileEmailAddress) IsConfirmed() (bool, os.Error) {
	return strconv.Atob(e.IsConfirmed_)
}

// Returns whether the email address 
func (e *ProfileEmailAddress) IsPrimary() (bool, os.Error) {
	return strconv.Atob(e.IsPrimary_)
}

// Group contains data about a group in TripIt
// All Group elements are read-only
type Group struct {
	DisplayName string `json:"display_name"` // read-only
	Url         string `json:"url"`          // read-only
}

// Trip Invitee
// All Invitee elements are read-only
type Invitee struct {
	IsReadOnly_ string            `json:"is_read_only"` // read-only
	IsTraveler_ string            `json:"is_traveler"`  // read-only
	Attributes  InviteeAttributes `json:"_attributes"`  // read-only, Use the profile_ref attribute to reference a Profile
}

// Returns whether the Invitee is read-only
func (i *Invitee) IsReadOnly() (bool, os.Error) {
	return strconv.Atob(i.IsReadOnly_)
}

// Returns whether the Invitee is a traveler on the trip
func (i *Invitee) IsTraveler() (bool, os.Error) {
	return strconv.Atob(i.IsTraveler_)
}

// Used to link to user profiles
type InviteeAttributes struct {
	ProfileRef string `json:"profile_ref"` // read-only, used to reference a profile
}

// A CRS remark
// All TripCrsRemark elements are read-only
type TripCrsRemark struct {
	RecordLocator string `json:"record_locator"` // read-only
	Notes         string `json:"notes"`          // read-only
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
	ClosenessMatches       *ClosenessMatches `json:"ClosenessMatches"`         // optional, ClosenessMatches are read-only
	TripInvitees           *TripInvitees     `json:"TripInvitees"`             // optional, TripInvitees are read-only
	TripCrsRemarks         *TripCrsRemarks   `json:"TripCrsRemarks"`           // optional, TripCrsRemarks are read-only
	Id_                    *string           `json:"id"`                       // optional, id is a read-only field
	RelativeUrl            string            `json:"relative_url"`             // optional, relative_url is a read-only field
	StartDate_             string            `json:"start_date"`               // optional, xs:date
	EndDate_               string            `json:"end_date"`                 // optional, xs:date
	Description            string            `json:"description"`              // optional
	DisplayName            string            `json:"display_name"`             // optional
	ImageUrl               string            `json:"image_url"`                // optional
	IsPrivate_             *string           `json:"is_private"`               // optional
	PrimaryLocation        string            `json:"primary_location"`         // optional
	PrimaryLocationAddress *Address          `json:"primary_location_address"` // optional, PrimaryLocationAddress is a read-only field
}

// People invited to view a trip
type TripInvitees struct {
	Invitee *InviteeVector `json:"Invitee"` // optional, TripInvitees are read-only
}

// TripIt users who are near this trip
type ClosenessMatches struct {
	ClosenessMatch *ClosenessMatchVector `json:"Match"` // optional, ClosenessMatches are read-only
}

// Remarks from a reservation system
type TripCrsRemarks struct {
	TripCrsRemark *TripCrsRemarkVector `json:"TripCrsRemark"` // optional, TripCrsRemarks are read-only
}

// returns the ID of the trip
func (t *Trip) Id() (uint, os.Error) {
	if t.Id_ == nil {
		return 0, os.NewError("Id field is not specified")
	}
	return strconv.Atoui(*t.Id_)
}

// Returns whether the trip is private
func (t *Trip) IsPrivate() (bool, os.Error) {
	if t.IsPrivate_ == nil {
		return false, os.NewError("IsPrivate field is not specified")
	}
	return strconv.Atob(*t.IsPrivate_)
}

// returns a time.Time object for StartDate
// Note: This won't have proper time zone information
func (t *Trip) StartDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", t.StartDate_)
}

// returns a time.Time object for EndDate
// Note: This won't have proper time zone information
func (t *Trip) EndDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", t.EndDate_)
}

// AirObject contains data about a flight
type AirObject struct {
	Id_                  *string              `json:"id"`                     // optional, read-only
	TripId_              *string              `json:"trip_id"`                // optional
	IsClientTraveler_    *string              `json:"is_client_traveler"`     // optional, read-only
	RelativeUrl          string               `json:"relative_url"`           // optional, read-only
	DisplayName          string               `json:"display_name"`           // optional
	Image                *ImagePtrVector      `json:"Image"`                  // optional
	CancellationDateTime *DateTime            `json:"CancellationDateTime"`   // optional
	BookingDate_         string               `json:"booking_date"`           // optional, xs:date
	BookingRate          string               `json:"booking_rate"`           // optional
	BookingSiteConfNum   string               `json:"booking_site_conf_num"`  // optional
	BookingSiteName      string               `json:"booking_site_name"`      // optional
	BookingSitePhone     string               `json:"booking_site_phone"`     // optional
	BookingSiteUrl       string               `json:"booking_site_url"`       // optional
	RecordLocator        string               `json:"record_locator"`         // optional
	SupplierConfNum      string               `json:"supplier_conf_num"`      // optional
	SupplierContact      string               `json:"supplier_contact"`       // optional
	SupplierEmailAddress string               `json:"supplier_email_address"` // optional
	SupplierName         string               `json:"supplier_name"`          // optional
	SupplierPhone        string               `json:"supplier_phone"`         // optional
	SupplierUrl          string               `json:"supplier_url"`           // optional
	IsPurchased_         *string              `json:"is_purchased"`           // optional
	Notes                string               `json:"notes"`                  // optional
	Restrictions         string               `json:"restrictions"`           // optional
	TotalCost            string               `json:"total_cost"`             // optional
	Segment              *AirSegmentPtrVector `json:"Segment"`
	Traveler             *TravelerPtrVector   `json:"Traveler"` // optional
}

// Returns the ID
func (o *AirObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

// Returns the associated trip ID
func (o *AirObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

// Returns whether the client is a traveler
func (o *AirObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *AirObject) BookingDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate_)
}

// returns whether the flights have been purchased
func (o *AirObject) IsPurchased() (bool, os.Error) {
	if o.IsPurchased_ == nil {
		return false, os.NewError("IsPurchased not specified")
	}
	return strconv.Atob(*o.IsPurchased_)
}

// AirSegment contains details about individual flights
type AirSegment struct {
	Status                 *FlightStatus `json:"Status"`                  // optional
	StartDateTime          *DateTime     `json:"StartDateTime"`           // optional
	EndDateTime            *DateTime     `json:"EndDateTime"`             // optional
	StartAirportCode       string        `json:"start_airport_code"`      // optional
	StartAirportLatitude_  *string       `json:"start_airport_latitude"`  // optional, read-only
	StartAirportLongitude_ *string       `json:"start_airport_longitude"` // optional, read-only
	StartCityName          string        `json:"start_city_name"`         // optional
	StartGate              string        `json:"start_gate"`              // optional
	StartTerminal          string        `json:"start_terminal"`          // optional
	EndAirportCode         string        `json:"end_airport_code"`        // optional
	EndAirportLatitude_    *string       `json:"end_airport_latitude"`    // optional, read-only
	EndAirportLongitude_   *string       `json:"end_airport_longitude"`   // optional, read-only
	EndCityName            string        `json:"end_city_name"`           // optional
	EndGate                string        `json:"end_gate"`                // optional
	EndTerminal            string        `json:"end_terminal"`            // optional
	MarketingAirline       string        `json:"marketing_airline"`       // optional
	MarketingAirlineCode   string        `json:"marketing_airline_code"`  // optional, read-only
	MarketingFlightNumber  string        `json:"marketing_flight_number"` // optional
	OperatingAirline       string        `json:"operating_airline"`       // optional
	OperatingAirlineCode   string        `json:"operating_airline_code"`  // optional, read-only
	OperatingFlightNumber  string        `json:"operating_flight_number"` // optional
	AlternativeFlightsUrl  string        `json:"alternate_flights_url"`   // optional, read-only
	Aircraft               string        `json:"aircraft"`                // optional
	AircraftDisplayName    string        `json:"aircraft_display_name"`   // optional, read-only
	Distance               string        `json:"distance"`                // optional
	Duration               string        `json:"duration"`                // optional
	Entertainment          string        `json:"entertainment"`           // optional
	Meal                   string        `json:"meal"`                    // optional
	Notes                  string        `json:"notes"`                   // optional
	OntimePerc             string        `json:"ontime_perc"`             // optional
	Seats                  string        `json:"seats"`                   // optional
	ServiceClass           string        `json:"service_class"`           // optional
	Stops                  string        `json:"stops"`                   // optional
	BaggageClaim           string        `json:"baggage_claim"`           // optional
	CheckInUrl             string        `json:"check_in_url"`            // optional
	ConflictResolutionUrl  string        `json:"conflict_resolution_url"` // optional, read-only
	IsHidden_              *string       `json:"is_hidden"`               // optional, read-only
	Id_                    *string       `json:"id"`                      // optional, read-only
}

func (s *AirSegment) StartAirportLatitude() (float64, os.Error) {
	if s.StartAirportLatitude_ == nil {
		return 0.0, os.NewError("StartAirportLatitude not specified")
	}
	return strconv.Atof64(*s.StartAirportLatitude_)
}

func (s *AirSegment) StartAirportLongitude() (float64, os.Error) {
	if s.StartAirportLongitude_ == nil {
		return 0.0, os.NewError("StartAirportLongitude not specified")
	}
	return strconv.Atof64(*s.StartAirportLongitude_)
}

func (s *AirSegment) EndAirportLatitude() (float64, os.Error) {
	if s.EndAirportLatitude_ == nil {
		return 0.0, os.NewError("EndAirportLatitude not specified")
	}
	return strconv.Atof64(*s.EndAirportLatitude_)
}

func (s *AirSegment) EndAirportLongitude() (float64, os.Error) {
	if s.EndAirportLongitude_ == nil {
		return 0.0, os.NewError("EndAirportLongitude not specified")
	}
	return strconv.Atof64(*s.EndAirportLongitude_)
}

func (s *AirSegment) Id() (uint, os.Error) {
	if s.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*s.Id_)
}

func (s *AirSegment) IsHidden() (bool, os.Error) {
	if s.IsHidden_ == nil {
		return false, os.NewError("IsHidden not specified")
	}
	return strconv.Atob(*s.IsHidden_)
}

// LodgingObject contains information about hotels or other lodging
// hotel cancellation remarks should be in restrictions
// hotel room description should be in notes
// hotel average daily rate should be in booking_rate
type LodgingObject struct {
	Id_                  *string            `json:"id"`                     // optional, read-only
	TripId_              *string            `json:"trip_id"`                // optional
	IsClientTraveler_    *string            `json:"is_client_traveler"`     // optional, read-only
	RelativeUrl          string             `json:"relative_url"`           // optional, read-only
	DisplayName          string             `json:"display_name"`           // optional
	Image                *ImagePtrVector    `json:"Image"`                  // optional
	CancellationDateTime *DateTime          `json:"CancellationDateTime"`   // optional
	BookingDate_         string             `json:"booking_date"`           // optional, xs:date
	BookingRate          string             `json:"booking_rate"`           // optional
	BookingSiteConfNum   string             `json:"booking_site_conf_num"`  // optional
	BookingSiteName      string             `json:"booking_site_name"`      // optional
	BookingSitePhone     string             `json:"booking_site_phone"`     // optional
	BookingSiteUrl       string             `json:"booking_site_url"`       // optional
	RecordLocator        string             `json:"record_locator"`         // optional
	SupplierConfNum      string             `json:"supplier_conf_num"`      // optional
	SupplierContact      string             `json:"supplier_contact"`       // optional
	SupplierEmailAddress string             `json:"supplier_email_address"` // optional
	SupplierName         string             `json:"supplier_name"`          // optional
	SupplierPhone        string             `json:"supplier_phone"`         // optional
	SupplierUrl          string             `json:"supplier_url"`           // optional
	IsPurchased_         *string            `json:"is_purchased"`           // optional
	Notes                string             `json:"notes"`                  // optional
	Restrictions         string             `json:"restrictions"`           // optional
	TotalCost            string             `json:"total_cost"`             // optional
	StartDateTime        *DateTime          `json:"StartDateTime"`          // optional
	EndDateTime          *DateTime          `json:"EndDateTime"`            // optional
	Address              *Address           `json:"Address"`                // optional
	Guest                *TravelerPtrVector `json:"Guest"`                  // optional
	NumberGuests         string             `json:"number_guests"`          // optional
	NumberRooms          string             `json:"numer_rooms"`            // optional
	RoomType             string             `json:"room_type"`              // optional
}

func (o *LodgingObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *LodgingObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *LodgingObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *LodgingObject) BookingDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate_)
}

func (o *LodgingObject) IsPurchased() (bool, os.Error) {
	if o.IsPurchased_ == nil {
		return false, os.NewError("IsPurchased not specified")
	}
	return strconv.Atob(*o.IsPurchased_)
}

// CarObject contains information about rental cars
// car cancellation remarks should be in restrictions
// car pickup instructions should be in notes
// car daily rate should be in booking_rate
type CarObject struct {
	Id_                  *string            `json:"id"`                     // optional, read-only
	TripId_              *string            `json:"trip_id"`                // optional
	IsClientTraveler_    *string            `json:"is_client_traveler"`     // optional, read-only
	RelativeUrl          string             `json:"relative_url"`           // optional, read-only
	DisplayName          string             `json:"display_name"`           // optional
	Image                *ImagePtrVector    `json:"Image"`                  // optional
	CancellationDateTime *DateTime          `json:"CancellationDateTime"`   // optional
	BookingDate_         string             `json:"booking_date"`           // optional, xs:date
	BookingRate          string             `json:"booking_rate"`           // optional
	BookingSiteConfNum   string             `json:"booking_site_conf_num"`  // optional
	BookingSiteName      string             `json:"booking_site_name"`      // optional
	BookingSitePhone     string             `json:"booking_site_phone"`     // optional
	BookingSiteUrl       string             `json:"booking_site_url"`       // optional
	RecordLocator        string             `json:"record_locator"`         // optional
	SupplierConfNum      string             `json:"supplier_conf_num"`      // optional
	SupplierContact      string             `json:"supplier_contact"`       // optional
	SupplierEmailAddress string             `json:"supplier_email_address"` // optional
	SupplierName         string             `json:"supplier_name"`          // optional
	SupplierPhone        string             `json:"supplier_phone"`         // optional
	SupplierUrl          string             `json:"supplier_url"`           // optional
	IsPurchased_         *string            `json:"is_purchased"`           // optional
	Notes                string             `json:"notes"`                  // optional
	Restrictions         string             `json:"restrictions"`           // optional
	TotalCost            string             `json:"total_cost"`             // optional
	StartDateTime        *DateTime          `json:"StartDateTime"`          // optional
	EndDateTime          *DateTime          `json:"EndDateTime"`            // optional
	StartLocationAddress *Address           `json:"StartLocationAddress"`   // optional
	EndLocationAddress   *Address           `json:"EndLocationAddress"`     // optional
	Driver               *TravelerPtrVector `json:"Driver"`                 // optional
	StartLocationHours   string             `json:"start_location_hours"`   // optional
	StartLocationName    string             `json:"start_location_name"`    // optional
	StartLocationPhone   string             `json:"start_location_phone"`   // optional
	EndLocationHours     string             `json:"end_location_hours"`     // optional
	EndLocationName      string             `json:"end_location_name"`      // optional
	EndLocationPhone     string             `json:"end_location_phone"`     // optional
	CarDescription       string             `json:"car_description"`        // optional
	CarType              string             `json:"car_type"`               // optional
	MileageCharges       string             `json:"mileage_charges"`        // optional
}

func (o *CarObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *CarObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *CarObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *CarObject) BookingDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate_)
}

func (o *CarObject) IsPurchased() (bool, os.Error) {
	if o.IsPurchased_ == nil {
		return false, os.NewError("IsPurchased not specified")
	}
	return strconv.Atob(*o.IsPurchased_)
}

// RailObject contains information about trains
type RailObject struct {
	Id_                  *string               `json:"id"`                     // optional, read-only
	TripId_              *string               `json:"trip_id"`                // optional
	IsClientTraveler_    *string               `json:"is_client_traveler"`     // optional, read-only
	RelativeUrl          string                `json:"relative_url"`           // optional, read-only
	DisplayName          string                `json:"display_name"`           // optional
	Image                *ImagePtrVector       `json:"Image"`                  // optional
	CancellationDateTime *DateTime             `json:"CancellationDateTime"`   // optional
	BookingDate_         string                `json:"booking_date"`           // optional, xs:date
	BookingRate          string                `json:"booking_rate"`           // optional
	BookingSiteConfNum   string                `json:"booking_site_conf_num"`  // optional
	BookingSiteName      string                `json:"booking_site_name"`      // optional
	BookingSitePhone     string                `json:"booking_site_phone"`     // optional
	BookingSiteUrl       string                `json:"booking_site_url"`       // optional
	RecordLocator        string                `json:"record_locator"`         // optional
	SupplierConfNum      string                `json:"supplier_conf_num"`      // optional
	SupplierContact      string                `json:"supplier_contact"`       // optional
	SupplierEmailAddress string                `json:"supplier_email_address"` // optional
	SupplierName         string                `json:"supplier_name"`          // optional
	SupplierPhone        string                `json:"supplier_phone"`         // optional
	SupplierUrl          string                `json:"supplier_url"`           // optional
	IsPurchased_         *string               `json:"is_purchased"`           // optional
	Notes                string                `json:"notes"`                  // optional
	Restrictions         string                `json:"restrictions"`           // optional
	TotalCost            string                `json:"total_cost"`             // optional
	Segment              *RailSegmentPtrVector `json:"Segment"`
	Traveler             *TravelerPtrVector    `json:"Traveler"` // optional
}

func (o *RailObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *RailObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *RailObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *RailObject) BookingDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate_)
}

func (o *RailObject) IsPurchased() (bool, os.Error) {
	if o.IsPurchased_ == nil {
		return false, os.NewError("IsPurchased not specified")
	}
	return strconv.Atob(*o.IsPurchased_)
}

// RailSegment contains details about an indivual train ride
type RailSegment struct {
	StartDateTime       *DateTime `json:"StartDateTime"`       // optional
	EndDateTime         *DateTime `json:"EndDateTime"`         // optional
	StartStationAddress *Address  `json:"StartStationAddress"` // optional
	EndStationAddress   *Address  `json:"EndStationAddress"`   // optional
	StartStationName    string    `json:"start_station_name"`  // optional
	EndStationName      string    `json:"end_station_name"`    // optional
	CarrierName         string    `json:"carrier_name"`        // optional
	CoachNumber         string    `json:"coach_number"`        // optional
	ConfirmationNum     string    `json:"confirmation_num"`    // optional
	Seats               string    `json:"seats"`               // optional
	ServiceClass        string    `json:"service_class"`       // optional
	TrainNumber         string    `json:"train_number"`        // optional
	TrainType           string    `json:"train_type"`          // optional
	Id_                 *string   `json:"id"`                  // optional, read-only
}

func (r *RailSegment) Id() (uint, os.Error) {
	if r.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*r.Id_)
}

// Transport Detail Types
const (
	TransportDetailTypeFerry                = "F"
	TransportDetailTypeGroundTransportation = "G"
)

// TransportObject contains details about other forms of transport like bus rides
type TransportObject struct {
	Id_                  *string                    `json:"id"`                     // optional, read-only
	TripId_              *string                    `json:"trip_id"`                // optional
	IsClientTraveler_    *string                    `json:"is_client_traveler"`     // optional, read-only
	RelativeUrl          string                     `json:"relative_url"`           // optional, read-only
	DisplayName          string                     `json:"display_name"`           // optional
	Image                *ImagePtrVector            `json:"Image"`                  // optional
	CancellationDateTime *DateTime                  `json:"CancellationDateTime"`   // optional
	BookingDate_         string                     `json:"booking_date"`           // optional, xs:date
	BookingRate          string                     `json:"booking_rate"`           // optional
	BookingSiteConfNum   string                     `json:"booking_site_conf_num"`  // optional
	BookingSiteName      string                     `json:"booking_site_name"`      // optional
	BookingSitePhone     string                     `json:"booking_site_phone"`     // optional
	BookingSiteUrl       string                     `json:"booking_site_url"`       // optional
	RecordLocator        string                     `json:"record_locator"`         // optional
	SupplierConfNum      string                     `json:"supplier_conf_num"`      // optional
	SupplierContact      string                     `json:"supplier_contact"`       // optional
	SupplierEmailAddress string                     `json:"supplier_email_address"` // optional
	SupplierName         string                     `json:"supplier_name"`          // optional
	SupplierPhone        string                     `json:"supplier_phone"`         // optional
	SupplierUrl          string                     `json:"supplier_url"`           // optional
	IsPurchased_         *string                    `json:"is_purchased"`           // optional
	Notes                string                     `json:"notes"`                  // optional
	Restrictions         string                     `json:"restrictions"`           // optional
	TotalCost            string                     `json:"total_cost"`             // optional
	Segment              *TransportSegmentPtrVector `json:"Segment"`
	Traveler             *TravelerPtrVector         `json:"Traveler"` // optional
}

func (o *TransportObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *TransportObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *TransportObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *TransportObject) BookingDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate_)
}

func (o *TransportObject) IsPurchased() (bool, os.Error) {
	if o.IsPurchased_ == nil {
		return false, os.NewError("IsPurchased not specified")
	}
	return strconv.Atob(*o.IsPurchased_)
}

// TransportSegment contains details about indivual transport rides
type TransportSegment struct {
	StartDateTime        *DateTime `json:"StartDateTime"`        // optional
	EndDateTime          *DateTime `json:"EndDateTime"`          // optional
	StartLocationAddress *Address  `json:"StartLocationAddress"` // optional
	EndLocationAddress   *Address  `json:"EndLocationAddress"`   // optional
	StartLocationName    string    `json:"start_location_name"`  // optional
	EndLocationName      string    `json:"end_location_name"`    // optional
	DetailTypeCode       string    `json:"detail_type_code"`     // optional
	CarrierName          string    `json:"carrier_name"`         // optional
	ConfirmationNum      string    `json:"confirmation_num"`     // optional
	NumberPassengers     string    `json:"number_passengers"`    // optional
	VehicleDescription   string    `json:"vehicle_description"`  // optional
	Id_                  *string   `json:"id"`                   // optional, read-only
}

func (r *TransportSegment) Id() (uint, os.Error) {
	if r.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*r.Id_)
}

// Cruise Detail Types
const (
	CruiseDetailTypePortOfCall = "P"
)

// CruiseObject contains information about cruises
type CruiseObject struct {
	Id_                  *string                 `json:"id"`                     // optional, read-only
	TripId_              *string                 `json:"trip_id"`                // optional
	IsClientTraveler_    *string                 `json:"is_client_traveler"`     // optional, read-only
	RelativeUrl          string                  `json:"relative_url"`           // optional, read-only
	DisplayName          string                  `json:"display_name"`           // optional
	Image                *ImagePtrVector         `json:"Image"`                  // optional
	CancellationDateTime *DateTime               `json:"CancellationDateTime"`   // optional
	BookingDate_         string                  `json:"booking_date"`           // optional, xs:date
	BookingRate          string                  `json:"booking_rate"`           // optional
	BookingSiteConfNum   string                  `json:"booking_site_conf_num"`  // optional
	BookingSiteName      string                  `json:"booking_site_name"`      // optional
	BookingSitePhone     string                  `json:"booking_site_phone"`     // optional
	BookingSiteUrl       string                  `json:"booking_site_url"`       // optional
	RecordLocator        string                  `json:"record_locator"`         // optional
	SupplierConfNum      string                  `json:"supplier_conf_num"`      // optional
	SupplierContact      string                  `json:"supplier_contact"`       // optional
	SupplierEmailAddress string                  `json:"supplier_email_address"` // optional
	SupplierName         string                  `json:"supplier_name"`          // optional
	SupplierPhone        string                  `json:"supplier_phone"`         // optional
	SupplierUrl          string                  `json:"supplier_url"`           // optional
	IsPurchased_         *string                 `json:"is_purchased"`           // optional
	Notes                string                  `json:"notes"`                  // optional
	Restrictions         string                  `json:"restrictions"`           // optional
	TotalCost            string                  `json:"total_cost"`             // optional
	Segment              *CruiseSegmentPtrVector `json:"Segment"`
	Traveler             *TravelerPtrVector      `json:"Traveler"`     // optional
	CabinNumber          string                  `json:"cabin_number"` // optional
	CabinType            string                  `json:"cabin_type"`   // optional
	Dining               string                  `json:"dining"`       // optional
	ShipName             string                  `json:"ship_name"`    // optional
}

func (o *CruiseObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *CruiseObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *CruiseObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *CruiseObject) BookingDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate_)
}

func (o *CruiseObject) IsPurchased() (bool, os.Error) {
	if o.IsPurchased_ == nil {
		return false, os.NewError("IsPurchased not specified")
	}
	return strconv.Atob(*o.IsPurchased_)
}

// CruiseSegment contains details about indivual cruise segments
type CruiseSegment struct {
	StartDateTime   *DateTime `json:"StartDateTime"`    // optional
	EndDateTime     *DateTime `json:"EndDateTime"`      // optional
	LocationAddress *Address  `json:"LocationAddress"`  // optional
	LocationName    string    `json:"location_name"`    // optional
	DetailTypeCode  string    `json:"detail_type_code"` // optional
	Id_             *string   `json:"id"`               // optional, read-only
}

func (r *CruiseSegment) Id() (uint, os.Error) {
	if r.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*r.Id_)
}

// RestaurantObject contains details about dining reservations
// restaurant name should be in supplier_name
// restaurant notes should be in notes
type RestaurantObject struct {
	Id_                  *string         `json:"id"`                     // optional, read-only
	TripId_              *string         `json:"trip_id"`                // optional
	IsClientTraveler_    *string         `json:"is_client_traveler"`     // optional, read-only
	RelativeUrl          string          `json:"relative_url"`           // optional, read-only
	DisplayName          string          `json:"display_name"`           // optional
	Image                *ImagePtrVector `json:"Image"`                  // optional
	CancellationDateTime *DateTime       `json:"CancellationDateTime"`   // optional
	BookingDate_         string          `json:"booking_date"`           // optional, xs:date
	BookingRate          string          `json:"booking_rate"`           // optional
	BookingSiteConfNum   string          `json:"booking_site_conf_num"`  // optional
	BookingSiteName      string          `json:"booking_site_name"`      // optional
	BookingSitePhone     string          `json:"booking_site_phone"`     // optional
	BookingSiteUrl       string          `json:"booking_site_url"`       // optional
	RecordLocator        string          `json:"record_locator"`         // optional
	SupplierConfNum      string          `json:"supplier_conf_num"`      // optional
	SupplierContact      string          `json:"supplier_contact"`       // optional
	SupplierEmailAddress string          `json:"supplier_email_address"` // optional
	SupplierName         string          `json:"supplier_name"`          // optional
	SupplierPhone        string          `json:"supplier_phone"`         // optional
	SupplierUrl          string          `json:"supplier_url"`           // optional
	IsPurchased_         *string         `json:"is_purchased"`           // optional
	Notes                string          `json:"notes"`                  // optional
	Restrictions         string          `json:"restrictions"`           // optional
	TotalCost            string          `json:"total_cost"`             // optional
	DateTime             *DateTime       `json:"DateTime"`               // optional
	Address              *Address        `json:"Address"`                // optional
	ReservationHolder    *Traveler       `json:"ReservationHolder"`      // optional
	Cuisine              string          `json:"cuisine"`                // optional
	DressCode            string          `json:"dress_code"`             // optional
	Hours                string          `json:"hours"`                  // optional
	NumberPatrons        string          `json:"number_patrons"`         // optional
	PriceRange           string          `json:"price_range"`            // optional
}

func (o *RestaurantObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *RestaurantObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *RestaurantObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *RestaurantObject) BookingDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate_)
}

func (o *RestaurantObject) IsPurchased() (bool, os.Error) {
	if o.IsPurchased_ == nil {
		return false, os.NewError("IsPurchased not specified")
	}
	return strconv.Atob(*o.IsPurchased_)
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
	Id_                  *string            `json:"id"`                     // optional, read-only
	TripId_              *string            `json:"trip_id"`                // optional
	IsClientTraveler_    *string            `json:"is_client_traveler"`     // optional, read-only
	RelativeUrl          string             `json:"relative_url"`           // optional, read-only
	DisplayName          string             `json:"display_name"`           // optional
	Image                *ImagePtrVector    `json:"Image"`                  // optional
	CancellationDateTime *DateTime          `json:"CancellationDateTime"`   // optional
	BookingDate_         string             `json:"booking_date"`           // optional, xs:date
	BookingRate          string             `json:"booking_rate"`           // optional
	BookingSiteConfNum   string             `json:"booking_site_conf_num"`  // optional
	BookingSiteName      string             `json:"booking_site_name"`      // optional
	BookingSitePhone     string             `json:"booking_site_phone"`     // optional
	BookingSiteUrl       string             `json:"booking_site_url"`       // optional
	RecordLocator        string             `json:"record_locator"`         // optional
	SupplierConfNum      string             `json:"supplier_conf_num"`      // optional
	SupplierContact      string             `json:"supplier_contact"`       // optional
	SupplierEmailAddress string             `json:"supplier_email_address"` // optional
	SupplierName         string             `json:"supplier_name"`          // optional
	SupplierPhone        string             `json:"supplier_phone"`         // optional
	SupplierUrl          string             `json:"supplier_url"`           // optional
	IsPurchased_         *string            `json:"is_purchased"`           // optional
	Notes                string             `json:"notes"`                  // optional
	Restrictions         string             `json:"restrictions"`           // optional
	TotalCost            string             `json:"total_cost"`             // optional
	StartDateTime        *DateTime          `json:"StartDateTime"`          // optional
	EndTime              string             `json:"end_time"`               // optional, xs:time
	Address              *Address           `json:"Address"`                // optional
	Participant          *TravelerPtrVector `json:"Participant"`            // optional
	DetailTypeCode       string             `json:"detail_type_code"`       // optional
	LocationName         string             `json:"location_name"`          // optional
}

func (o *ActivityObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *ActivityObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *ActivityObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *ActivityObject) BookingDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate_)
}

func (o *ActivityObject) IsPurchased() (bool, os.Error) {
	if o.IsPurchased_ == nil {
		return false, os.NewError("IsPurchased not specified")
	}
	return strconv.Atob(*o.IsPurchased_)
}

// Note Detail Types
const (
	NoteDetailTypeArticle = "A"
)

// NoteObject contains information about notes added by the traveler
type NoteObject struct {
	Id_               *string         `json:"id"`                 // optional, read-only
	TripId_           *string         `json:"trip_id"`            // optional
	IsClientTraveler_ *string         `json:"is_client_traveler"` // optional, read-only
	RelativeUrl       string          `json:"relative_url"`       // optional, read-only
	DisplayName       string          `json:"display_name"`       // optional
	Image             *ImagePtrVector `json:"Image"`              // optional
	DateTime          *DateTime       `json:"DateTime"`           // optional
	Address           *Address        `json:"Address"`            // optional
	DetailTypeCode    string          `json:"detail_type_code"`   // optional
	Source            string          `json:"source"`             // optional
	Text              string          `json:"text"`               // optional
	Url               string          `json:"url"`                // optional
	Notes             string          `json:"notes"`              // optional
}

func (o *NoteObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *NoteObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *NoteObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// MapObject contains addresses to show on a map
type MapObject struct {
	Id_               *string         `json:"id"`                 // optional, read-only
	TripId_           *string         `json:"trip_id"`            // optional
	IsClientTraveler_ *string         `json:"is_client_traveler"` // optional, read-only
	RelativeUrl       string          `json:"relative_url"`       // optional, read-only
	DisplayName       string          `json:"display_name"`       // optional
	Image             *ImagePtrVector `json:"Image"`              // optional
	DateTime          *DateTime       `json:"DateTime"`           // optional
	Address           *Address        `json:"Address"`            // optional
}

func (o *MapObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *MapObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *MapObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// DirectionsObject contains addresses to show directions for on the trip
type DirectionsObject struct {
	Id_               *string         `json:"id"`                 // optional, read-only
	TripId_           *string         `json:"trip_id"`            // optional
	IsClientTraveler_ *string         `json:"is_client_traveler"` // optional, read-only
	RelativeUrl       string          `json:"relative_url"`       // optional, read-only
	DisplayName       string          `json:"display_name"`       // optional
	Image             *ImagePtrVector `json:"Image"`              // optional
	DateTime          *DateTime       `json:"DateTime"`           // optional
	StartAddress      *Address        `json:"StartAddress"`       // optional
	EndAddress        *Address        `json:"EndAddress"`         // optional
}

func (o *DirectionsObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *DirectionsObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *DirectionsObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// WeatherObject contains information about the weather at a particular destination
// Weather is read-only
type WeatherObject struct {
	Id_                 *string         `json:"id"`                   // optional, read-only
	TripId_             *string         `json:"trip_id"`              // optional
	IsClientTraveler_   *string         `json:"is_client_traveler"`   // optional, read-only
	RelativeUrl         string          `json:"relative_url"`         // optional, read-only
	DisplayName         string          `json:"display_name"`         // optional
	Image               *ImagePtrVector `json:"Image"`                // optional
	Date_               string          `json:"date"`                 // optional, read-only, xs:date
	Location            string          `json:"location"`             // optional, read-only
	AvgHighTempC_       *string         `json:"avg_high_temp_c"`      // optional, read-only
	AvgLowTempC_        *string         `json:"avg_low_temp_c"`       // optional, read-only
	AvgWindSpeedKn_     *string         `json:"avg_wind_speed_kn"`    // optional, read-only
	AvgPrecipitationCm_ *string         `json:"avg_precipitation_cm"` // optional, read-only
	AvgSnowDepthCm_     *string         `json:"avg_snow_depth_cm"`    // optional, read-only
}

func (o *WeatherObject) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *WeatherObject) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *WeatherObject) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

// returns a time.Time object for StartDate
// Note: This won't have proper time zone information
func (w *WeatherObject) Date() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", w.Date_)
}

func (w *WeatherObject) AvgHighTempC() (float64, os.Error) {
	if w.AvgHighTempC_ == nil {
		return 0.0, os.NewError("AvgHighTempC not specified")
	}
	return strconv.Atof64(*w.AvgHighTempC_)
}

func (w *WeatherObject) AvgLowTempC() (float64, os.Error) {
	if w.AvgLowTempC_ == nil {
		return 0.0, os.NewError("AvgLowTempC not specified")
	}
	return strconv.Atof64(*w.AvgLowTempC_)
}

func (w *WeatherObject) AvgWindSpeedKn() (float64, os.Error) {
	if w.AvgWindSpeedKn_ == nil {
		return 0.0, os.NewError("AvgWindSpeedKn not specified")
	}
	return strconv.Atof64(*w.AvgWindSpeedKn_)
}

func (w *WeatherObject) AvgPrecipitationCm() (float64, os.Error) {
	if w.AvgPrecipitationCm_ == nil {
		return 0.0, os.NewError("AvgPrecipitationCm not specified")
	}
	return strconv.Atof64(*w.AvgPrecipitationCm_)
}

func (w *WeatherObject) AvgSnowDepthCm() (float64, os.Error) {
	if w.AvgSnowDepthCm_ == nil {
		return 0.0, os.NewError("AvgSnowDepthCm not specified")
	}
	return strconv.Atof64(*w.AvgSnowDepthCm_)
}
