package tripit

import (
	"time"
	"fmt"
	"os"
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
	Code              int      "code"                // read-only
	DetailedErrorCode *float64 "detailed_error_code" // optional, read-only
	Description       string   "description"         // read-only
	EntityType        string   "entity_type"         // read-only
	Timestamp         string   "timestamp"           // read-only, xs:datetime
}

// returns a time.Time object for the Timestamp
func (e *Error) Time() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, e.Timestamp)
}

// Returns a string containing the error information
func (e *Error) String() string {
	return fmt.Sprintf("TripIt Error %d: %s", e.Code, e.Description)
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

// Represents a TripIt API Response
type Response struct {
	Timestamp        int                "timestamp"
	NumBytes         int                "num_bytes"
	Error            []Error            "Error"            // optional
	Warning          []Warning          "Warning"          // optional
	Trip             []Trip             "Trip"             // optional
	ActivityObject   []ActivityObject   "ActivityObject"   // optional
	AirObject        []AirObject        "AirObject"        // optional
	CarObject        []CarObject        "CarObject"        // optional
	CruiseObject     []CruiseObject     "CruiseObject"     // optional
	DirectionsObject []DirectionsObject "DirectionsObject" // optional
	LodgingObject    []LodgingObject    "LodgingObject"    // optional
	MapObject        []MapObject        "MapObject"        // optional
	NoteObject       []NoteObject       "NoteObject"       // optional
	RailObject       []RailObject       "RailObject"       // optional
	RestaurantObject []RestaurantObject "RestaurantObject" // optional
	TransportObject  []TransportObject  "TransportObject"  // optional
	WeatherObject    []WeatherObject    "WeatherObject"    // optional
	PointsProgram    []PointsProgram    "PointsProgram"    // optional
	Profile          []Profile          "Profile"          // optional
	// @TODO need to add invitee stuff
}

// returns a time.Time object for the Timestamp
func (r *Response) Time() (*time.Time, os.Error) {
	return time.SecondsToUTC(int64(r.Timestamp)), nil
}

/*
   	For create, use either:
	- address for single-line addresses.
	- addr1, addr2, city, state, zip, and country for multi-line addresses.
	Multi-line address will be ignored if single-line address is present.
	See documentation for more information.
*/
type Address struct {
	Address   string   "address"   // optional
	Addr1     string   "addr1"     // optional
	Addr2     string   "addr2"     // optional
	City      string   "city"      // optional
	State     string   "state"     // optional
	Zip       string   "zip"       // optional
	Country   string   "country"   // optional
	Latitude  *float64 "latitude"  // optional, read-only
	Longitude *float64 "longitude" // optional, read-only
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
	FlightStatus               *int      "flight_status"              // optional, read-only
	IsConnectionAtRisk         *bool     "is_connection_at_risk"      // optional, read-only
	DepartureTerminal          string    "departure_terminal"         // optional, read-only
	DepartureGate              string    "departure_gate"             // optional, read-only
	ArrivalTerminal            string    "arrival_terminal"           // optional, read-only
	ArrivalGate                string    "arrival_gate"               // optional, read-only
	LayoverMinutes             string    "layover_minutes"            // optional, read-only
	BaggageClaim               string    "baggage_claim"              // optional, read-only
	DivertedAirportCode        string    "diverted_airport_code"      // optional, read-only
	LastModified               int       "last_modified"              // read-only
}

// returns a time.Time object for LastModified
func (fs *FlightStatus) Time() (*time.Time, os.Error) {
	return time.SecondsToUTC(int64(fs.LastModified)), nil
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
	DateFld   string "date"       // optional, xs:date
	TimeFld   string "time"       // optional, xs:time
	Timezone  string "timezone"   // optional, read-only
	UtcOffset string "utc_offset" // optional, read-only
}

func (dt DateTime) Time() (*time.Time, os.Error) {
	if dt.UtcOffset == "" {
		return time.Parse(time.RFC3339, fmt.Sprintf("%sT%sZ", dt.DateFld, dt.TimeFld))
	}
	return time.Parse(time.RFC3339, fmt.Sprintf("%sT%s%s", dt.DateFld, dt.TimeFld, dt.UtcOffset))
}

func (dt *DateTime) SetTime(t *time.Time) {
	dt.DateFld = t.Format("2006-01-02")
	dt.TimeFld = t.Format("15:04:05")
	dt.UtcOffset = t.Format("-07:00")
	dt.Timezone = t.Format("MST")
}

// All PointsProgram elements are read-only
type PointsProgram struct {
	Id                  uint                      "id"                    // read-only
	Name                string                    "name optional"         // read-only
	AccountNumber       string                    "account_number"        // optional, read-only
	AccountLogin        string                    "account_login"         // optional, read-only
	Balance             string                    "balance"               // optional, read-only
	EliteStatus         string                    "elite_status"          // optional, read-only
	EliteNextStatus     string                    "elite_next_status"     // optional, read-only
	EliteYtdQualify     string                    "elite_ytd_qualify"     // optional, read-only
	EliteNeedToEarn     string                    "elite_need_to_earn"    // optional, read-only
	LastModified        int                       "last_modified"         // read-only
	TotalNumActivities  int                       "total_num_activities"  // read-only
	TotalNumExpirations int                       "total_num_expirations" // read-only
	ErrorMessage        string                    "error_message"         // optional, read-only
	Activity            []PointsProgramActivity   "Activity"              // optional, read-only
	Expiration          []PointsProgramExpiration "Expiration"            // optional, read-only
}

// returns a time.Time object for LastModified
func (pp *PointsProgram) Time() (*time.Time, os.Error) {
	return time.SecondsToUTC(int64(pp.LastModified)), nil
}

// All PointsProgramActivity elements are read-only
type PointsProgramActivity struct {
	Date        string "date"        // read-only, xs:date
	Description string "description" // optional, read-only
	Base        string "base"        // optional, read-only
	Bonus       string "bonus"       // optional, read-only
	Total       string "total"       // optional, read-only
}

// returns a time.Time object for Date
// Note: This won't have proper time zone information
func (pa *PointsProgramActivity) Time() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", pa.Date)
}

// All PointsProgramExpiration elements are read-only
type PointsProgramExpiration struct {
	Date   string "date"   // read-only, xs:date
	Amount string "amount" // optional, read-only
}

// returns a time.Time object for Date
// Note: This won't have proper time zone information
func (pe *PointsProgramExpiration) Time() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", pe.Date)
}

type TripShare struct {
	TripId            uint "trip_id"
	IsTraveler        bool "is_traveler"
	IsReadOnly        bool "is_read_only"
	IsSentWithDetails bool "is_sent_with_details"
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
	Ref                   string                "ref"                   // read-only
	ProfileEmailAddresses []ProfileEmailAddress "ProfileEmailAddresses" // optional, read-only
	GroupMemberships      []Group               "GroupMemberships"      // optional, read-only
	IsClient              bool                  "is_client"             // read-only
	IsPro                 bool                  "is_pro"                // read-only
	ScreenNanem           string                "screen_name"           // read-only
	PublicDisplayName     string                "public_display_name"   // read-only
	ProfileUrl            string                "profile_url"           // read-only
	HomeCity              string                "home_city"             // optional, read-only
	Company               string                "company"               // optional, read-only
	AboutMeInfo           string                "about_me_info"         // optional, read-only
	PhotoUrl              string                "photo_url"             // optional, read-only
	ActivityFeedUrl       string                "activity_feed_url"     // optional, read-only
	AlertsFeedUrl         string                "alerts_feed_url"       // optional, read-only
	IcalUrl               string                "ical_url"              // optional, read-only
}

// All ProfileEmailAddress elements are read-only
type ProfileEmailAddress struct {
	Address      string "address"        // read-only
	IsAutoImport bool   "is_auto_import" // read-only
	IsConfirmed  bool   "is_confirmed"   // read-only
	IsPrimary    bool   "is_primary"     // read-only
}

// All Group elements are read-only
type Group struct {
	DisplayName string "display_name" // read-only
	Url         string "url"          // read-only
}

// All Invitee elements are read-only
type Invitee struct {
	IsReadOnly bool   "is_read_only" // read-only
	IsTraveler bool   "is_traveler"  // read-only
	ProfileRef string "profile_ref"  // read-only, Use the profile_ref attribute to reference a Profile
}

// All TripCrsRemark elements are read-only
type TripCrsRemark struct {
	RecordLocator string "record_locator" // read-only
	Notes         string "notes"          // read-only
}

// All ClosenessMatch elements are read-only
type ClosenessMatch struct {
	ProfileRef string "profile_ref" // read-only, Use the profile_ref attribute to reference a Profile
}

type Trip struct {
	ClosenessMatches       []ClosenessMatch "ClosenessMatches"         // optional, ClosenessMatches are read-only
	Invitees               []Invitee        "Invitees"                 // optional, TripInvitees are read-only
	TripCrsRemarks         []TripCrsRemark  "TripCrsRemarks"           // optional, TripCrsRemarks are read-only
	Id                     *uint            "id"                       // optional, id is a read-only field
	RelativeUrl            string           "relative_url"             // optional, relative_url is a read-only field
	StartDate              string           "start_date"               // optional, xs:date
	EndDate                string           "end_date"                 // optional, xs:date
	Description            string           "description"              // optional
	DisplayName            string           "display_name"             // optional
	ImageUrl               string           "image_url"                // optional
	IsPrivate              *bool            "is_private"               // optional
	PrimaryLocation        string           "primary_location"         // optional
	PrimaryLocationAddress *Address         "primary_location_address" // optional, PrimaryLocationAddress is a read-only field
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

type Object struct {
	Id               *uint   "id"                 // optional, read-only
	TripId           *uint   "trip_id"            // optional
	IsClientTraveler *bool   "is_client_traveler" // optional, read-only
	RelativeUrl      string  "relative_url"       // optional, read-only
	DisplayName      string  "display_name"       // optional
	Image            []Image "Image"              // optional
}

type ReservationObject struct {
	Object
	CancellationDateTime *DateTime "CancellationDateTime"   // optional
	BookingDate          string    "booking_date"           // optional, xs:date
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
	IsPurchased          *bool     "is_purchased"           // optional
	Notes                string    "notes"                  // optional
	Restrictions         string    "restrictions"           // optional
	TotalCost            string    "total_cost"             // optional
}

// returns a time.Time object for BookingDate
// Note: This won't have proper time zone information
func (r *ReservationObject) BookingTime() (*time.Time, os.Error) {
	return time.Parse("2006-01-02", r.BookingDate)
}

type AirObject struct {
	ReservationObject
	Segment  []AirSegment "Segment"
	Traveler []Traveler   "Traveler" // optional
}

type AirSegment struct {
	Status                *int      "Status"                  // optional
	StartDateTime         *DateTime "StartDateTime"           // optional
	EndDateTime           *DateTime "EndDateTime"             // optional
	StartAirportCode      string    "start_airport_code"      // optional
	StartAirportLatitude  *float64  "start_airport_latitude"  // optional, read-only
	StartAirportLongitude *float64  "start_airport_longitude" // optional, read-only
	StartCityName         string    "start_city_name"         // optional
	StartGate             string    "start_gate"              // optional
	StartTerminal         string    "start_terminal"          // optional
	EndAirportCode        string    "end_airport_code"        // optional
	EndAirportLatitude    *float64  "end_airport_latitude"    // optional, read-only
	EndAirportLongitude   *float64  "end_airport_longitude"   // optional, read-only
	EndCityName           string    "end_city_name"           // optional
	EndGate               string    "end_gate"                // optional
	EndTerminal           string    "end_terminal"            // optional
	MarketingAirline      string    "marketing_airline"       // optional
	MarketingAirlineCode  string    "marketing_airline_code"  // optional, read-only
	MarketingFlightNumber string    "marketing_flight_number" // optional
	OperatingAirline      string    "operating_airline"       // optional
	OperatingAirlineCode  string    "operating_airline_code"  // optional, read-only
	OperatingFlightNumber string    "operating_flight_number" // optional
	AlternativeFlightsUrl string    "alternate_flights_url"   // optional, read-only
	Aircraft              string    "aircraft"                // optional
	AircraftDisplayName   string    "aircraft_display_name"   // optional, read-only
	Distance              string    "distance"                // optional
	Duration              string    "duration"                // optional
	Entertainment         string    "entertainment"           // optional
	Meal                  string    "meal"                    // optional
	Notes                 string    "notes"                   // optional
	OntimePerc            string    "ontime_perc"             // optional
	Seats                 string    "seats"                   // optional
	ServiceClass          string    "service_class"           // optional
	Stops                 string    "stops"                   // optional
	BaggageClaim          string    "baggage_claim"           // optional
	CheckInUrl            string    "check_in_url"            // optional
	ConflictResolutionUrl string    "conflict_resolution_url" // optional, read-only
	IsHidden              *bool     "is_hidden"               // optional, read-only
	Id                    *uint     "id"                      // optional, read-only
}

// hotel cancellation remarks should be in restrictions
// hotel room description should be in notes
// hotel average daily rate should be in booking_rate
type LodgingObject struct {
	ReservationObject
	StartDateTime *DateTime  "StartDateTime" // optional
	EndDateTime   *DateTime  "EndDateTime"   // optional
	Address       *Address   "Address"       // optional
	Guest         []Traveler "Guest"         // optional
	NumberGuests  string     "number_guests" // optional
	NumberRooms   string     "numer_rooms"   // optional
	RoomType      string     "room_type"     // optional
}

// car cancellation remarks should be in restrictions
// car pickup instructions should be in notes
// car daily rate should be in booking_rate
type CarObject struct {
	ReservationObject
	StartDateTime        *DateTime  "StartDateTime"        // optional
	EndDateTime          *DateTime  "EndDateTime"          // optional
	StartLocationAddress *Address   "StartLocationAddress" // optional
	EndLocationAddress   *Address   "EndLocationAddress"   // optional
	Driver               []Traveler "Driver"               // optional
	StartLocationHours   string     "start_location_hours" // optional
	StartLocationName    string     "start_location_name"  // optional
	StartLocationPhone   string     "start_location_phone" // optional
	EndLocationHours     string     "end_location_hours"   // optional
	EndLocationName      string     "end_location_name"    // optional
	EndLocationPhone     string     "end_location_phone"   // optional
	CarDescription       string     "car_description"      // optional
	CarType              string     "car_type"             // optional
	MileageCharges       string     "mileage_charges"      // optional
}

type RailObject struct {
	ReservationObject
	Segment  []RailSegment "Segment"
	Traveler []Traveler    "Traveler" // optional
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
	Id                  *uint     "id"                  // optional, read-only
}

// Transport Detail Types
const (
	TransportDetailTypeFerry                = "F"
	TransportDetailTypeGroundTransportation = "G"
)

type TransportObject struct {
	ReservationObject
	Segment  []TransportSegment "Segment"
	Traveler []Traveler         "Traveler" // optional
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
	Id                   *uint     "id"                   // optional, read-only
}

// Cruise Detail Types
const (
	CruiseDetailTypePortOfCall = "P"
)

type CruiseObject struct {
	ReservationObject
	Segment     []CruiseSegment "Segment"
	Traveler    []Traveler      "Traveler"     // optional
	CabinNumber string          "cabin_number" // optional
	CabinType   string          "cabin_type"   // optional
	Dining      string          "dining"       // optional
	ShipName    string          "ship_name"    // optional
}

type CruiseSegment struct {
	StartDateTime   *DateTime "StartDateTime"    // optional
	EndDateTime     *DateTime "EndDateTime"      // optional
	LocationAddress *Address  "LocationAddress"  // optional
	LocationName    string    "location_name"    // optional
	DetailTypeCode  string    "detail_type_code" // optional
	Id              *uint     "id"               // optional, read-only
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
	StartDateTime  *DateTime  "StartDateTime"    // optional
	EndTime        string     "end_time"         // optional, xs:time
	Address        *Address   "Address"          // optional
	Participant    []Traveler "Participant"      // optional
	DetailTypeCode string     "detail_type_code" // optional
	LocationName   string     "location_name"    // optional
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
	Date               string   "date"                 // optional, read-only, xs:date
	Location           string   "location"             // optional, read-only
	AvgHighTempC       *float64 "avg_high_temp_c"      // optional, read-only
	AvgLowTempC        *float64 "avg_low_temp_c"       // optional, read-only
	AvgWindSpeedKn     *float64 "avg_wind_speed_kn"    // optional, read-only
	AvgPrecipitationCm *float64 "avg_precipitation_cm" // optional, read-only
	AvgSnowDepthCm     *float64 "avg_snow_depth_cm"    // optional, read-only
}
