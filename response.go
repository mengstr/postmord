package postmord

import (
	"time"
)

type Response struct {
	TrackingInformation TrackingInformation `json:"TrackingInformationResponse"`
}

type ParamValues struct {
	Param string `json:"param"`
	Value string `json:"value"`
}

type Faults struct {
	FaultCode       string        `json:"faultCode"`
	ExplanationText string        `json:"explanationText"`
	ParamValues     []ParamValues `json:"paramValues"`
}

type CompositeFault struct {
	Faults []Faults `json:"faults"`
}

type Service struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Address struct {
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Country     string `json:"country"`
	PostCode    string `json:"postCode"`
}

type Consignor struct {
	Name       string  `json:"name"`
	Issuercode string  `json:"issuercode"`
	Address    Address `json:"address"`
}

type Consignee struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type Contact struct {
	ContactName string `json:"contactName"`
	Phone       string `json:"phone"`
	MobilePhone string `json:"mobilePhone"`
	Email       string `json:"email"`
}

type ReturnParty struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Contact Contact `json:"contact"`
}

type PickupParty struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Contact Contact `json:"contact"`
}

type CollectionParty struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Contact Contact `json:"contact"`
}

type StatusText struct {
	Header                 string `json:"header"`
	Body                   string `json:"body"`
	EstimatedTimeOfArrival string `json:"estimatedTimeOfArrival"`
}

type Coordinate struct {
	SrID     string `json:"srId"`
	Northing string `json:"northing"`
	Easting  string `json:"easting"`
}

type OpeningHour struct {
	OpenFrom  string `json:"openFrom"`
	OpenTo    string `json:"openTo"`
	OpenFrom2 string `json:"openFrom2"`
	OpenTo2   string `json:"openTo2"`
	Monday    bool   `json:"monday"`
	Tuesday   bool   `json:"tuesday"`
	Wednesday bool   `json:"wednesday"`
	Thursday  bool   `json:"thursday"`
	Friday    bool   `json:"friday"`
	Saturday  bool   `json:"saturday"`
	Sunday    bool   `json:"sunday"`
}

type DeliveryPoint struct {
	Name             string        `json:"name"`
	LocationDetail   string        `json:"locationDetail"`
	Address          Address       `json:"address"`
	Contact          Contact       `json:"contact"`
	Coordinate       []Coordinate  `json:"coordinate"`
	OpeningHour      []OpeningHour `json:"openingHour"`
	DisplayName      string        `json:"displayName"`
	LocationID       string        `json:"locationId"`
	ServicePointType string        `json:"servicePointType"`
}

type DestinationDeliveryPoint struct {
	Name             string        `json:"name"`
	LocationDetail   string        `json:"locationDetail"`
	Address          Address       `json:"address"`
	Contact          Contact       `json:"contact"`
	Coordinate       []Coordinate  `json:"coordinate"`
	OpeningHour      []OpeningHour `json:"openingHour"`
	DisplayName      string        `json:"displayName"`
	LocationID       string        `json:"locationId"`
	ServicePointType string        `json:"servicePointType"`
}

type TotalWeight struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type TotalVolume struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type AssessedWeight struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type AssessedVolume struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type SplitStatuses struct {
	NoItemsWithStatus int    `json:"noItemsWithStatus"`
	NoItems           int    `json:"noItems"`
	StatusDescription string `json:"statusDescription"`
	Status            string `json:"status"`
}

type ShipmentReferences struct {
	Value string `json:"value"`
	Type  string `json:"type"`
	Name  string `json:"name"`
}

type AdditionalServices struct {
	Code      string `json:"code"`
	GroupCode string `json:"groupCode"`
	Name      string `json:"name"`
}

type Acceptor struct {
	SignatureReference string `json:"signatureReference"`
	Name               string `json:"name"`
}

type Weight struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type Length struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type Height struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type Width struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type Volume struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type StatedMeasurement struct {
	Weight Weight `json:"weight"`
	Length Length `json:"length"`
	Height Height `json:"height"`
	Width  Width  `json:"width"`
	Volume Volume `json:"volume"`
}

type AssessedMeasurement struct {
	Weight Weight `json:"weight"`
	Length Length `json:"length"`
	Height Height `json:"height"`
	Width  Width  `json:"width"`
	Volume Volume `json:"volume"`
}

type Location struct {
	Name         string `json:"name"`
	CountryCode  string `json:"countryCode"`
	Country      string `json:"country"`
	LocationID   string `json:"locationId"`
	DisplayName  string `json:"displayName"`
	Postcode     string `json:"postcode"`
	City         string `json:"city"`
	LocationType string `json:"locationType"`
}

type GeoLocation struct {
	GeoNorthing        int    `json:"geoNorthing"`
	GeoEasting         int    `json:"geoEasting"`
	GeoReferenceSystem string `json:"geoReferenceSystem"`
	GeoPostalCode      string `json:"geoPostalCode"`
	GeoCity            string `json:"geoCity"`
	GeoCountryCode     string `json:"geoCountryCode"`
}

type Events struct {
	EventTime          time.Time   `json:"eventTime"`
	EventCode          string      `json:"eventCode"`
	Location           Location    `json:"location"`
	GeoLocation        GeoLocation `json:"geoLocation"`
	Status             string      `json:"status"`
	EventDescription   string      `json:"eventDescription"`
	LocalDeviationDode string      `json:"localDeviationDode"`
}

type References struct {
	Value string `json:"value"`
	Type  string `json:"type"`
	Name  string `json:"name"`
}

type FreeText struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type Items struct {
	ItemID                 string              `json:"itemId"`
	EstimatedTimeOfArrival time.Time           `json:"estimatedTimeOfArrival"`
	DropOffDate            time.Time           `json:"dropOffDate"`
	DeliveryDate           time.Time           `json:"deliveryDate"`
	ReturnDate             time.Time           `json:"returnDate"`
	TypeOfItem             string              `json:"typeOfItem"`
	TypeOfItemName         string              `json:"typeOfItemName"`
	TypeOfItemActual       string              `json:"typeOfItemActual"`
	TypeOfItemActualName   string              `json:"typeOfItemActualName"`
	AdditionalInformation  string              `json:"additionalInformation"`
	NoItems                int                 `json:"noItems"`
	NumberOfPallets        string              `json:"numberOfPallets"`
	Status                 string              `json:"status"`
	StatusText             StatusText          `json:"statusText"`
	Acceptor               Acceptor            `json:"acceptor"`
	StatedMeasurement      StatedMeasurement   `json:"statedMeasurement"`
	AssessedMeasurement    AssessedMeasurement `json:"assessedMeasurement"`
	Events                 []Events            `json:"events"`
	References             []References        `json:"references"`
	PreviousItemStates     []string            `json:"previousItemStates"`
	FreeText               []FreeText          `json:"freeText"`
}

type Shipments struct {
	ShipmentID               string                   `json:"shipmentId"`
	URI                      string                   `json:"uri"`
	AssessedNumberOfItems    int                      `json:"assessedNumberOfItems"`
	CashOnDeliveryText       string                   `json:"cashOnDeliveryText"`
	DeliveryDate             time.Time                `json:"deliveryDate"`
	ReturnDate               time.Time                `json:"returnDate"`
	EstimatedTimeOfArrival   time.Time                `json:"estimatedTimeOfArrival"`
	NumberOfPallets          string                   `json:"numberOfPallets"`
	FlexChangePossible       bool                     `json:"flexChangePossible"`
	Service                  Service                  `json:"service"`
	Consignor                Consignor                `json:"consignor"`
	Consignee                Consignee                `json:"consignee"`
	ReturnParty              ReturnParty              `json:"returnParty"`
	PickupParty              PickupParty              `json:"pickupParty"`
	CollectionParty          CollectionParty          `json:"collectionParty"`
	StatusText               StatusText               `json:"statusText"`
	Status                   string                   `json:"status"`
	DeliveryPoint            DeliveryPoint            `json:"deliveryPoint"`
	DestinationDeliveryPoint DestinationDeliveryPoint `json:"destinationDeliveryPoint"`
	TotalWeight              TotalWeight              `json:"totalWeight"`
	TotalVolume              TotalVolume              `json:"totalVolume"`
	AssessedWeight           AssessedWeight           `json:"assessedWeight"`
	AssessedVolume           AssessedVolume           `json:"assessedVolume"`
	SplitStatuses            []SplitStatuses          `json:"splitStatuses"`
	ShipmentReferences       []ShipmentReferences     `json:"shipmentReferences"`
	AdditionalServices       []AdditionalServices     `json:"additionalServices"`
	HarmonizedVersion        int                      `json:"harmonizedVersion"`
	Items                    []Items                  `json:"items"`
}

type TrackingInformation struct {
	CompositeFault CompositeFault `json:"compositeFault"`
	Shipments      []Shipments    `json:"shipments"`
}
