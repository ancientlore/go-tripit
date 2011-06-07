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
	Invitation       []Invitation      "Invitation"       // optional
	Trip             *Trip             "Trip"             // optional
	ActivityObject   *ActivityObject   "ActivityObject"   // optional
	AirObject        *AirObject        "AirObject"        // optional
	CarObject        *CarObject        "CarObject"        // optional
	CruiseObject     *CruiseObject     "CruiseObject"     // optional
	DirectionsObject *DirectionsObject "DirectionsObject" // optional
	LodgingObject    *LodgingObject    "LodgingObject"    // optional
	MapObject        *MapObject        "MapObject"        // optional
	NoteObject       *NoteObject       "NoteObject"       // optional
	RailObject       *RailObject       "RailObject"       // optional
	RestaurantObject *RestaurantObject "RestaurantObject" // optional
	TransportObject  *TransportObject  "TransportObject"  // optional
}

// Error is returned from TripIt on error conditions
type Error struct {
	Code_              string  "code"                // read-only
	DetailedErrorCode_ *string "detailed_error_code" // optional, read-only
	Description        string  "description"         // read-only
	EntityType         string  "entity_type"         // read-only
	Timestamp_         string  "timestamp"           // read-only, xs:datetime
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
func (e *Error) Timestamp() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, e.Timestamp_)
}

// Returns a string containing the error information
func (e *Error) String() string {
	return fmt.Sprintf("TripIt Error %s: %s", e.Code_, e.Description)
}

// Warning is returned from TripIt to indicate warning conditions
type Warning struct {
	Description string "description" // read-only
	EntityType  string "entity_type" // read-only
	Timestamp   string "timestamp"   // read-only, xs:datetime
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
	Timestamp_       string                     "timestamp"
	NumBytes_        string                     "num_bytes"
	Error            *ErrorVector               "Error"            // optional
	Warning          *WarningVector             "Warning"          // optional
	Trip             *TripPtrVector             "Trip"             // optional
	ActivityObject   *ActivityObjectPtrVector   "ActivityObject"   // optional
	AirObject        *AirObjectPtrVector        "AirObject"        // optional
	CarObject        *CarObjectPtrVector        "CarObject"        // optional
	CruiseObject     *CruiseObjectPtrVector     "CruiseObject"     // optional
	DirectionsObject *DirectionsObjectPtrVector "DirectionsObject" // optional
	LodgingObject    *LodgingObjectPtrVector    "LodgingObject"    // optional
	MapObject        *MapObjectPtrVector        "MapObject"        // optional
	NoteObject       *NoteObjectPtrVector       "NoteObject"       // optional
	RailObject       *RailObjectPtrVector       "RailObject"       // optional
	RestaurantObject *RestaurantObjectPtrVector "RestaurantObject" // optional
	TransportObject  *TransportObjectPtrVector  "TransportObject"  // optional
	WeatherObject    *WeatherObjectVector       "WeatherObject"    // optional
	PointsProgram    *PointsProgramVector       "PointsProgram"    // optional
	Profile          *ProfileVector             "Profile"          // optional
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
	Address    string  "address"   // optional
	Addr1      string  "addr1"     // optional
	Addr2      string  "addr2"     // optional
	City       string  "city"      // optional
	State      string  "state"     // optional
	Zip        string  "zip"       // optional
	Country    string  "country"   // optional
	Latitude_  *string "latitude"  // optional, read-only
	Longitude_ *string "longitude" // optional, read-only
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
	FirstName                string "first_name"                 // optional
	MiddleName               string "middle_name"                // optional
	LastName                 string "last_name"                  // optional
	FrequentTravelerNum      string "frequent_traveler_num"      // optional
	FrequentTravelerSupplier string "frequent_traveler_supplier" // optional
	MealPreference           string "meal_preference"            // optional
	SeatPreference           string "seat_preference"            // optional
	TicketNum                string "ticket_num optional"
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
	ScheduledDepartureDateTime *DateTime "ScheduledDepartureDateTime" // optional, read-only
	EstimatedDepartureDateTime *DateTime "EstimatedDepartureDateTime" // optional, read-only
	ScheduledArrivalDateTime   *DateTime "ScheduledArrivalDateTime"   // optional, read-only
	EstimatedArrivalDateTime   *DateTime "EstimatedArrivalDateTime"   // optional, read-only
	FlightStatus_              *string   "flight_status"              // optional, read-only
	IsConnectionAtRisk_        *string   "is_connection_at_risk"      // optional, read-only
	DepartureTerminal          string    "departure_terminal"         // optional, read-only
	DepartureGate              string    "departure_gate"             // optional, read-only
	ArrivalTerminal            string    "arrival_terminal"           // optional, read-only
	ArrivalGate                string    "arrival_gate"               // optional, read-only
	LayoverMinutes             string    "layover_minutes"            // optional, read-only
	BaggageClaim               string    "baggage_claim"              // optional, read-only
	DivertedAirportCode        string    "diverted_airport_code"      // optional, read-only
	LastModified_              string    "last_modified"              // read-only
}

func (fs *FlightStatus) FlightStatus() (int, os.Error) {
	if fs.FlightStatus_ == nil {
		return 0, os.NewError("Flight status not specified")
	}
	return strconv.Atoi(*fs.FlightStatus_)
}

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

type Image struct {
	Caption string "caption" // optional
	Url     string "url"
}

// Stores date and time zone information, for example:
// {
//	 "date":"2009-11-10",
//   "time":"14:00:00",
//    "timezone":"America\/Los_Angeles",
//    "utc_offset":"-08:00"
// }
type DateTime struct {
	Date_     string "date"       // optional, xs:date
	Time_     string "time"       // optional, xs:time
	Timezone  string "timezone"   // optional, read-only
	UtcOffset string "utc_offset" // optional, read-only
}

func (dt DateTime) DateTime() (*time.Time, os.Error) {
	if dt.UtcOffset == "" {
		return time.Parse(time.RFC3339, fmt.Sprintf("%sT%sZ", dt.Date_, dt.Time_))
	}
	return time.Parse(time.RFC3339, fmt.Sprintf("%sT%s%s", dt.Date_, dt.Time_, dt.UtcOffset))
}

func (dt *DateTime) SetDateTime(t *time.Time) {
	dt.Date_ = t.Format("2006-01-02")
	dt.Time_ = t.Format("15:04:05")
	dt.UtcOffset = t.Format("-07:00")
	dt.Timezone = t.Format("MST")
}

// All PointsProgram elements are read-only
type PointsProgram struct {
	Id_                  string                         "id"                    // read-only
	Name                 string                         "name optional"         // read-only
	AccountNumber        string                         "account_number"        // optional, read-only
	AccountLogin         string                         "account_login"         // optional, read-only
	Balance              string                         "balance"               // optional, read-only
	EliteStatus          string                         "elite_status"          // optional, read-only
	EliteNextStatus      string                         "elite_next_status"     // optional, read-only
	EliteYtdQualify      string                         "elite_ytd_qualify"     // optional, read-only
	EliteNeedToEarn      string                         "elite_need_to_earn"    // optional, read-only
	LastModified_        string                         "last_modified"         // read-only
	TotalNumActivities_  string                         "total_num_activities"  // read-only
	TotalNumExpirations_ string                         "total_num_expirations" // read-only
	ErrorMessage         string                         "error_message"         // optional, read-only
	Activity             *PointsProgramActivityVector   "Activity"              // optional, read-only
	Expiration           *PointsProgramExpirationVector "Expiration"            // optional, read-only
}

func (pp *PointsProgram) Id() (uint, os.Error) {
	return strconv.Atoui(pp.Id_)
}

func (pp *PointsProgram) TotalNumActivities() (int, os.Error) {
	return strconv.Atoi(pp.TotalNumActivities_)
}

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

// All PointsProgramActivity elements are read-only
type PointsProgramActivity struct {
	Date_       string "date"        // read-only, xs:date
	Description string "description" // optional, read-only
	Base        string "base"        // optional, read-only
	Bonus       string "bonus"       // optional, read-only
	Total       string "total"       // optional, read-only
}

// returns a time.Time object for Date
// Note: This won't have proper time zone information
func (pa *PointsProgramActivity) Date() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", pa.Date_)
}

// All PointsProgramExpiration elements are read-only
type PointsProgramExpiration struct {
	Date_  string "date"   // read-only, xs:date
	Amount string "amount" // optional, read-only
}

// returns a time.Time object for Date
// Note: This won't have proper time zone information
func (pe *PointsProgramExpiration) Date() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", pe.Date_)
}

type TripShare struct {
	TripId_            string "trip_id"
	IsTraveler_        string "is_traveler"
	IsReadOnly_        string "is_read_only"
	IsSentWithDetails_ string "is_sent_with_details"
}

func (ts *TripShare) TripId() (uint, os.Error) {
	return strconv.Atoui(ts.TripId_)
}

func (ts *TripShare) IsTraveler() (bool, os.Error) {
	return strconv.Atob(ts.IsTraveler_)
}

func (ts *TripShare) IsReadOnly() (bool, os.Error) {
	return strconv.Atob(ts.IsReadOnly_)
}

func (ts *TripShare) IsSentWithDetails() (bool, os.Error) {
	return strconv.Atob(ts.IsSentWithDetails_)
}

type ConnectionRequest struct {

}

type Invitation struct {
	EmailAddresses    []string           "EmailAddresses"
	TripShare         *TripShare         "TripShare"         // optional
	ConnectionRequest *ConnectionRequest "ConnectionRequest" // optional
	Message           string             "message"           // optional
}

// All Profile elements are read-only
type Profile struct {
	Attributes            ProfileAttributes      "_attributes"           // read-only
	ProfileEmailAddresses *ProfileEmailAddresses "ProfileEmailAddresses" // optional, read-only
	GroupMemberships      *GroupMemberships      "GroupMemberships"      // optional, read-only
	IsClient_             string                 "is_client"             // read-only
	IsPro_                string                 "is_pro"                // read-only
	ScreenName            string                 "screen_name"           // read-only
	PublicDisplayName     string                 "public_display_name"   // read-only
	ProfileUrl            string                 "profile_url"           // read-only
	HomeCity              string                 "home_city"             // optional, read-only
	Company               string                 "company"               // optional, read-only
	AboutMeInfo           string                 "about_me_info"         // optional, read-only
	PhotoUrl              string                 "photo_url"             // optional, read-only
	ActivityFeedUrl       string                 "activity_feed_url"     // optional, read-only
	AlertsFeedUrl         string                 "alerts_feed_url"       // optional, read-only
	IcalUrl               string                 "ical_url"              // optional, read-only
}

type ProfileEmailAddresses struct {
	ProfileEmailAddress *ProfileEmailAddressVector "ProfileEmailAddress"
}

type GroupMemberships struct {
	Group *GroupVector "Group" // optional, read-only
}

type ProfileAttributes struct {
	Ref string "ref" // read-only
}

// Returns whether the profile is a client
func (p *Profile) IsClient() (bool, os.Error) {
	return strconv.Atob(p.IsClient_)
}

// Returns whether the profile has TripIt pro
func (p *Profile) IsPro() (bool, os.Error) {
	return strconv.Atob(p.IsPro_)
}

// All ProfileEmailAddress elements are read-only
type ProfileEmailAddress struct {
	Address       string "address"        // read-only
	IsAutoImport_ string "is_auto_import" // read-only
	IsConfirmed_  string "is_confirmed"   // read-only
	IsPrimary_    string "is_primary"     // read-only
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

// All Group elements are read-only
type Group struct {
	DisplayName string "display_name" // read-only
	Url         string "url"          // read-only
}

// All Invitee elements are read-only
type Invitee struct {
	IsReadOnly_ string            "is_read_only" // read-only
	IsTraveler_ string            "is_traveler"  // read-only
	Attributes  InviteeAttributes "_attributes"  // read-only, Use the profile_ref attribute to reference a Profile
}

func (i *Invitee) IsReadOnly() (bool, os.Error) {
	return strconv.Atob(i.IsReadOnly_)
}

func (i *Invitee) IsTraveler() (bool, os.Error) {
	return strconv.Atob(i.IsTraveler_)
}

type InviteeAttributes struct {
	ProfileRef string "profile_ref" // read-only, used to reference a profile
}

// All TripCrsRemark elements are read-only
type TripCrsRemark struct {
	RecordLocator string "record_locator" // read-only
	Notes         string "notes"          // read-only
}

// All ClosenessMatch elements are read-only
type ClosenessMatch struct {
	Attributes ClosenessMatchAttributes "_attributes" // read-only, Use the profile_ref attribute to reference a Profile
}

type ClosenessMatchAttributes struct {
	ProfileRef string "profile_ref" // read-only, Use the profile_ref attribute to reference a Profile
}

type Trip struct {
	ClosenessMatches       *ClosenessMatches "ClosenessMatches"         // optional, ClosenessMatches are read-only
	TripInvitees           *TripInvitees     "TripInvitees"             // optional, TripInvitees are read-only
	TripCrsRemarks         *TripCrsRemarks   "TripCrsRemarks"           // optional, TripCrsRemarks are read-only
	Id_                    *string           "id"                       // optional, id is a read-only field
	RelativeUrl            string            "relative_url"             // optional, relative_url is a read-only field
	StartDate_             string            "start_date"               // optional, xs:date
	EndDate_               string            "end_date"                 // optional, xs:date
	Description            string            "description"              // optional
	DisplayName            string            "display_name"             // optional
	ImageUrl               string            "image_url"                // optional
	IsPrivate_             *string           "is_private"               // optional
	PrimaryLocation        string            "primary_location"         // optional
	PrimaryLocationAddress *Address          "primary_location_address" // optional, PrimaryLocationAddress is a read-only field
}

type TripInvitees struct {
	Invitee *InviteeVector "Invitee" // optional, TripInvitees are read-only
}

type ClosenessMatches struct {
	ClosenessMatch *ClosenessMatchVector "Match" // optional, ClosenessMatches are read-only
}

type TripCrsRemarks struct {
	TripCrsRemark *TripCrsRemarkVector "TripCrsRemark" // optional, TripCrsRemarks are read-only
}

func (t *Trip) Id() (uint, os.Error) {
	if t.Id_ == nil {
		return 0, os.NewError("Id field is not specified")
	}
	return strconv.Atoui(*t.Id_)
}

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

type Object struct {
	Id_               *string         "id"                 // optional, read-only
	TripId_           *string         "trip_id"            // optional
	IsClientTraveler_ *string         "is_client_traveler" // optional, read-only
	RelativeUrl       string          "relative_url"       // optional, read-only
	DisplayName       string          "display_name"       // optional
	Image             *ImagePtrVector "Image"              // optional
}

func (o *Object) Id() (uint, os.Error) {
	if o.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*o.Id_)
}

func (o *Object) TripId() (uint, os.Error) {
	if o.TripId_ == nil {
		return 0, os.NewError("TripId not specified")
	}
	return strconv.Atoui(*o.TripId_)
}

func (o *Object) IsClientTraveler() (bool, os.Error) {
	if o.IsClientTraveler_ == nil {
		return false, os.NewError("IsClientTraveler not specified")
	}
	return strconv.Atob(*o.IsClientTraveler_)
}

type ReservationObject struct {
	Object
	CancellationDateTime *DateTime "CancellationDateTime"   // optional
	BookingDate_         string    "booking_date"           // optional, xs:date
	BookingRate          string    "booking_rate"           // optional
	BookingSiteConfNum   string    "booking_site_conf_num"  // optional
	BookingSiteName      string    "booking_site_name"      // optional
	BookingSitePhone     string    "booking_site_phone"     // optional
	BookingSiteUrl       string    "booking_site_url"       // optional
	RecordLocator        string    "record_locator"         // optional
	SupplierConfNum      string    "supplier_conf_num"      // optional
	SupplierContact      string    "supplier_contact"       // optional
	SupplierEmailAddress string    "supplier_email_address" // optional
	SupplierName         string    "supplier_name"          // optional
	SupplierPhone        string    "supplier_phone"         // optional
	SupplierUrl          string    "supplier_url"           // optional
	IsPurchased_         *string   "is_purchased"           // optional
	Notes                string    "notes"                  // optional
	Restrictions         string    "restrictions"           // optional
	TotalCost            string    "total_cost"             // optional
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *ReservationObject) BookingDate() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate_)
}

func (o *ReservationObject) IsPurchased() (bool, os.Error) {
	if o.IsPurchased_ == nil {
		return false, os.NewError("IsPurchased not specified")
	}
	return strconv.Atob(*o.IsPurchased_)
}

type AirObject struct {
	ReservationObject
	Segment  *AirSegmentPtrVector "Segment"
	Traveler *TravelerPtrVector   "Traveler" // optional
}

type AirSegment struct {
	Status                 *FlightStatus "Status"                  // optional
	StartDateTime          *DateTime     "StartDateTime"           // optional
	EndDateTime            *DateTime     "EndDateTime"             // optional
	StartAirportCode       string        "start_airport_code"      // optional
	StartAirportLatitude_  *string       "start_airport_latitude"  // optional, read-only
	StartAirportLongitude_ *string       "start_airport_longitude" // optional, read-only
	StartCityName          string        "start_city_name"         // optional
	StartGate              string        "start_gate"              // optional
	StartTerminal          string        "start_terminal"          // optional
	EndAirportCode         string        "end_airport_code"        // optional
	EndAirportLatitude_    *string       "end_airport_latitude"    // optional, read-only
	EndAirportLongitude_   *string       "end_airport_longitude"   // optional, read-only
	EndCityName            string        "end_city_name"           // optional
	EndGate                string        "end_gate"                // optional
	EndTerminal            string        "end_terminal"            // optional
	MarketingAirline       string        "marketing_airline"       // optional
	MarketingAirlineCode   string        "marketing_airline_code"  // optional, read-only
	MarketingFlightNumber  string        "marketing_flight_number" // optional
	OperatingAirline       string        "operating_airline"       // optional
	OperatingAirlineCode   string        "operating_airline_code"  // optional, read-only
	OperatingFlightNumber  string        "operating_flight_number" // optional
	AlternativeFlightsUrl  string        "alternate_flights_url"   // optional, read-only
	Aircraft               string        "aircraft"                // optional
	AircraftDisplayName    string        "aircraft_display_name"   // optional, read-only
	Distance               string        "distance"                // optional
	Duration               string        "duration"                // optional
	Entertainment          string        "entertainment"           // optional
	Meal                   string        "meal"                    // optional
	Notes                  string        "notes"                   // optional
	OntimePerc             string        "ontime_perc"             // optional
	Seats                  string        "seats"                   // optional
	ServiceClass           string        "service_class"           // optional
	Stops                  string        "stops"                   // optional
	BaggageClaim           string        "baggage_claim"           // optional
	CheckInUrl             string        "check_in_url"            // optional
	ConflictResolutionUrl  string        "conflict_resolution_url" // optional, read-only
	IsHidden_              *string       "is_hidden"               // optional, read-only
	Id_                    *string       "id"                      // optional, read-only
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

// hotel cancellation remarks should be in restrictions
// hotel room description should be in notes
// hotel average daily rate should be in booking_rate
type LodgingObject struct {
	ReservationObject
	StartDateTime *DateTime          "StartDateTime" // optional
	EndDateTime   *DateTime          "EndDateTime"   // optional
	Address       *Address           "Address"       // optional
	Guest         *TravelerPtrVector "Guest"         // optional
	NumberGuests  string             "number_guests" // optional
	NumberRooms   string             "numer_rooms"   // optional
	RoomType      string             "room_type"     // optional
}

// car cancellation remarks should be in restrictions
// car pickup instructions should be in notes
// car daily rate should be in booking_rate
type CarObject struct {
	ReservationObject
	StartDateTime        *DateTime          "StartDateTime"        // optional
	EndDateTime          *DateTime          "EndDateTime"          // optional
	StartLocationAddress *Address           "StartLocationAddress" // optional
	EndLocationAddress   *Address           "EndLocationAddress"   // optional
	Driver               *TravelerPtrVector "Driver"               // optional
	StartLocationHours   string             "start_location_hours" // optional
	StartLocationName    string             "start_location_name"  // optional
	StartLocationPhone   string             "start_location_phone" // optional
	EndLocationHours     string             "end_location_hours"   // optional
	EndLocationName      string             "end_location_name"    // optional
	EndLocationPhone     string             "end_location_phone"   // optional
	CarDescription       string             "car_description"      // optional
	CarType              string             "car_type"             // optional
	MileageCharges       string             "mileage_charges"      // optional
}

type RailObject struct {
	ReservationObject
	Segment  *RailSegmentPtrVector "Segment"
	Traveler *TravelerPtrVector    "Traveler" // optional
}

type RailSegment struct {
	StartDateTime       *DateTime "StartDateTime"       // optional
	EndDateTime         *DateTime "EndDateTime"         // optional
	StartStationAddress *Address  "StartStationAddress" // optional
	EndStationAddress   *Address  "EndStationAddress"   // optional
	StartStationName    string    "start_station_name"  // optional
	EndStationName      string    "end_station_name"    // optional
	CarrierName         string    "carrier_name"        // optional
	CoachNumber         string    "coach_number"        // optional
	ConfirmationNum     string    "confirmation_num"    // optional
	Seats               string    "seats"               // optional
	ServiceClass        string    "service_class"       // optional
	TrainNumber         string    "train_number"        // optional
	TrainType           string    "train_type"          // optional
	Id_                 *string   "id"                  // optional, read-only
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

type TransportObject struct {
	ReservationObject
	Segment  *TransportSegmentPtrVector "Segment"
	Traveler *TravelerPtrVector         "Traveler" // optional
}

type TransportSegment struct {
	StartDateTime        *DateTime "StartDateTime"        // optional
	EndDateTime          *DateTime "EndDateTime"          // optional
	StartLocationAddress *Address  "StartLocationAddress" // optional
	EndLocationAddress   *Address  "EndLocationAddress"   // optional
	StartLocationName    string    "start_location_name"  // optional
	EndLocationName      string    "end_location_name"    // optional
	DetailTypeCode       string    "detail_type_code"     // optional
	CarrierName          string    "carrier_name"         // optional
	ConfirmationNum      string    "confirmation_num"     // optional
	NumberPassengers     string    "number_passengers"    // optional
	VehicleDescription   string    "vehicle_description"  // optional
	Id_                  *string   "id"                   // optional, read-only
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

type CruiseObject struct {
	ReservationObject
	Segment     *CruiseSegmentPtrVector "Segment"
	Traveler    *TravelerPtrVector      "Traveler"     // optional
	CabinNumber string                  "cabin_number" // optional
	CabinType   string                  "cabin_type"   // optional
	Dining      string                  "dining"       // optional
	ShipName    string                  "ship_name"    // optional
}

type CruiseSegment struct {
	StartDateTime   *DateTime "StartDateTime"    // optional
	EndDateTime     *DateTime "EndDateTime"      // optional
	LocationAddress *Address  "LocationAddress"  // optional
	LocationName    string    "location_name"    // optional
	DetailTypeCode  string    "detail_type_code" // optional
	Id_             *string   "id"               // optional, read-only
}

func (r *CruiseSegment) Id() (uint, os.Error) {
	if r.Id_ == nil {
		return 0, os.NewError("Id not specified")
	}
	return strconv.Atoui(*r.Id_)
}

// restaurant name should be in supplier_name
// restaurant notes should be in notes
type RestaurantObject struct {
	ReservationObject
	DateTime          *DateTime "DateTime"          // optional
	Address           *Address  "Address"           // optional
	ReservationHolder *Traveler "ReservationHolder" // optional
	Cuisine           string    "cuisine"           // optional
	DressCode         string    "dress_code"        // optional
	Hours             string    "hours"             // optional
	NumberPatrons     string    "number_patrons"    // optional
	PriceRange        string    "price_range"       // optional
}

// Activity Detail Types
const (
	ActivityDetailTypeConcert = "C"
	ActivityDetailTypeTheatre = "H"
	ActivityDetailTypeMeeting = "M"
	ActivityDetailTypeTour    = "T"
)

type ActivityObject struct {
	ReservationObject
	StartDateTime  *DateTime          "StartDateTime"    // optional
	EndTime        string             "end_time"         // optional, xs:time
	Address        *Address           "Address"          // optional
	Participant    *TravelerPtrVector "Participant"      // optional
	DetailTypeCode string             "detail_type_code" // optional
	LocationName   string             "location_name"    // optional
}

// Note Detail Types
const (
	NoteDetailTypeArticle = "A"
)

type NoteObject struct {
	Object
	DateTime       *DateTime "DateTime"         // optional
	Address        *Address  "Address"          // optional
	DetailTypeCode string    "detail_type_code" // optional
	Source         string    "source"           // optional
	Text           string    "text"             // optional
	Url            string    "url"              // optional
	Notes          string    "notes"            // optional
}


type MapObject struct {
	Object
	DateTime *DateTime "DateTime" // optional
	Address  *Address  "Address"  // optional
}

type DirectionsObject struct {
	Object
	DateTime     *DateTime "DateTime"     // optional
	StartAddress *Address  "StartAddress" // optional
	EndAddress   *Address  "EndAddress"   // optional
}

// Weather is read-only
type WeatherObject struct {
	Object
	Date_               string  "date"                 // optional, read-only, xs:date
	Location            string  "location"             // optional, read-only
	AvgHighTempC_       *string "avg_high_temp_c"      // optional, read-only
	AvgLowTempC_        *string "avg_low_temp_c"       // optional, read-only
	AvgWindSpeedKn_     *string "avg_wind_speed_kn"    // optional, read-only
	AvgPrecipitationCm_ *string "avg_precipitation_cm" // optional, read-only
	AvgSnowDepthCm_     *string "avg_snow_depth_cm"    // optional, read-only
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
