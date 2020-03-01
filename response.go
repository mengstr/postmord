package postmord

import (
	"encoding/json"
	"encoding/xml"
	"github.com/go-yaml/yaml"
)

// TODO: Extract structs from this monster
type Response struct {
	Tracking *struct {
		CompositeFault *struct {
			Faults []*struct {
				FaultCode       string `json:"faultCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				ExplanationText string `json:"explanationText,omitempty" yaml:",omitempty" xml:",omitempty"`
				ParamValues     []*struct {
					Param string `json:"param,omitempty" yaml:",omitempty" xml:",omitempty"`
					Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"paramValues,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"faults,omitempty" yaml:",omitempty" xml:",omitempty"`
		} `json:"compositeFault,omitempty" yaml:",omitempty" xml:",omitempty"`
		Shipments []*struct {
			ShipmentID                   string `json:"shipmentId,omitempty" yaml:",omitempty" xml:",omitempty"`
			URI                          string `json:"uri,omitempty" yaml:",omitempty" xml:",omitempty"`
			AssessedNumberOfItems        int    `json:"assessedNumberOfItems,omitempty" yaml:",omitempty" xml:",omitempty"`
			CashOnDeliveryText           string `json:"cashOnDeliveryText,omitempty" yaml:",omitempty" xml:",omitempty"`
			DeliveryDate                 string `json:"deliveryDate,omitempty" yaml:",omitempty" xml:",omitempty"`
			OriginEstimatedTimeOfArrival string `json:"originEstimatedTimeOfArrival,omitempty" yaml:",omitempty" xml:",omitempty"`
			EstimatedTimeOfArrival       string `json:"estimatedTimeOfArrival,omitempty" yaml:",omitempty" xml:",omitempty"`
			RealTimeOfArrival            string `json:"realTimeOfArrival,omitempty" yaml:",omitempty" xml:",omitempty"`
			RequestedDeliveryDate        string `json:"requestedDeliveryDate,omitempty" yaml:",omitempty" xml:",omitempty"`
			RequestedProductionDate      string `json:"requestedProductionDate,omitempty" yaml:",omitempty" xml:",omitempty"`
			NotificationPhoneNumber      string `json:"notificationPhoneNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
			NotificationCode             string `json:"notificationCode,omitempty" yaml:",omitempty" xml:",omitempty"`
			CustomerNumber               string `json:"customerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
			NumberOfPallets              string `json:"numberOfPallets,omitempty" yaml:",omitempty" xml:",omitempty"`
			Service                      *struct {
				Code          string `json:"code,omitempty" yaml:",omitempty" xml:",omitempty"`
				SourceSystem  string `json:"sourceSystem,omitempty" yaml:",omitempty" xml:",omitempty"`
				Name          string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				ArticleNumber string `json:"articleNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"service,omitempty" yaml:",omitempty" xml:",omitempty"`
			Consignor *struct {
				Name       string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				Issuercode string `json:"issuercode,omitempty" yaml:",omitempty" xml:",omitempty"`
				Customer   *struct {
					ProductionCustomerNumber string `json:"productionCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
					ExternalCustomerNumber   string `json:"externalCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
					SapCustomerNumber        string `json:"sapCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"customer,omitempty" yaml:",omitempty" xml:",omitempty"`
				Address *struct {
					Street1     string `json:"street1,omitempty" yaml:",omitempty" xml:",omitempty"`
					Street2     string `json:"street2,omitempty" yaml:",omitempty" xml:",omitempty"`
					City        string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
					CountryCode string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					Country     string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
					PostCode    string `json:"postCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"address,omitempty" yaml:",omitempty" xml:",omitempty"`
				Contact *struct {
					ContactName string `json:"contactName,omitempty" yaml:",omitempty" xml:",omitempty"`
					Phone       string `json:"phone,omitempty" yaml:",omitempty" xml:",omitempty"`
					MobilePhone string `json:"mobilePhone,omitempty" yaml:",omitempty" xml:",omitempty"`
					Email       string `json:"email,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"contact,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"consignor,omitempty" yaml:",omitempty" xml:",omitempty"`
			Consignee *struct {
				Name    string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				Address *struct {
					Street1     string `json:"street1,omitempty" yaml:",omitempty" xml:",omitempty"`
					Street2     string `json:"street2,omitempty" yaml:",omitempty" xml:",omitempty"`
					City        string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
					CountryCode string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					Country     string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
					PostCode    string `json:"postCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"address,omitempty" yaml:",omitempty" xml:",omitempty"`
				Contact *struct {
					ContactName string `json:"contactName,omitempty" yaml:",omitempty" xml:",omitempty"`
					Phone       string `json:"phone,omitempty" yaml:",omitempty" xml:",omitempty"`
					MobilePhone string `json:"mobilePhone,omitempty" yaml:",omitempty" xml:",omitempty"`
					Email       string `json:"email,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"contact,omitempty" yaml:",omitempty" xml:",omitempty"`
				Customer *struct {
					ProductionCustomerNumber string `json:"productionCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
					ExternalCustomerNumber   string `json:"externalCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
					SapCustomerNumber        string `json:"sapCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"customer,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"consignee,omitempty" yaml:",omitempty" xml:",omitempty"`
			OriginalShipper *struct {
				Address *struct {
					Street1     string `json:"street1,omitempty" yaml:",omitempty" xml:",omitempty"`
					Street2     string `json:"street2,omitempty" yaml:",omitempty" xml:",omitempty"`
					City        string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
					CountryCode string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					Country     string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
					PostCode    string `json:"postCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"address,omitempty" yaml:",omitempty" xml:",omitempty"`
				Name     string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				Customer *struct {
					ProductionCustomerNumber string `json:"productionCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
					ExternalCustomerNumber   string `json:"externalCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
					SapCustomerNumber        string `json:"sapCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"customer,omitempty" yaml:",omitempty" xml:",omitempty"`
				Contact *struct {
					ContactName string `json:"contactName,omitempty" yaml:",omitempty" xml:",omitempty"`
					Phone       string `json:"phone,omitempty" yaml:",omitempty" xml:",omitempty"`
					MobilePhone string `json:"mobilePhone,omitempty" yaml:",omitempty" xml:",omitempty"`
					Email       string `json:"email,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"contact,omitempty" yaml:",omitempty" xml:",omitempty"`
				Issuercode string `json:"issuercode,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"originalShipper,omitempty" yaml:",omitempty" xml:",omitempty"`
			FreightPayer *struct {
				Name     string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				Customer *struct {
					ProductionCustomerNumber string `json:"productionCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
					ExternalCustomerNumber   string `json:"externalCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
					SapCustomerNumber        string `json:"sapCustomerNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"customer,omitempty" yaml:",omitempty" xml:",omitempty"`
				Contact *struct {
					ContactName string `json:"contactName,omitempty" yaml:",omitempty" xml:",omitempty"`
					Phone       string `json:"phone,omitempty" yaml:",omitempty" xml:",omitempty"`
					MobilePhone string `json:"mobilePhone,omitempty" yaml:",omitempty" xml:",omitempty"`
					Email       string `json:"email,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"contact,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"freightPayer,omitempty" yaml:",omitempty" xml:",omitempty"`
			ReturnParty *struct {
				Name    string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				Address *struct {
					Street1     string `json:"street1,omitempty" yaml:",omitempty" xml:",omitempty"`
					Street2     string `json:"street2,omitempty" yaml:",omitempty" xml:",omitempty"`
					City        string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
					CountryCode string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					Country     string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
					PostCode    string `json:"postCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"address,omitempty" yaml:",omitempty" xml:",omitempty"`
				Contact *struct {
					ContactName string `json:"contactName,omitempty" yaml:",omitempty" xml:",omitempty"`
					Phone       string `json:"phone,omitempty" yaml:",omitempty" xml:",omitempty"`
					MobilePhone string `json:"mobilePhone,omitempty" yaml:",omitempty" xml:",omitempty"`
					Email       string `json:"email,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"contact,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"returnParty,omitempty" yaml:",omitempty" xml:",omitempty"`
			PickupParty *struct {
				Name    string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				Address *struct {
					Street1     string `json:"street1,omitempty" yaml:",omitempty" xml:",omitempty"`
					Street2     string `json:"street2,omitempty" yaml:",omitempty" xml:",omitempty"`
					City        string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
					CountryCode string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					Country     string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
					PostCode    string `json:"postCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"address,omitempty" yaml:",omitempty" xml:",omitempty"`
				Contact *struct {
					ContactName string `json:"contactName,omitempty" yaml:",omitempty" xml:",omitempty"`
					Phone       string `json:"phone,omitempty" yaml:",omitempty" xml:",omitempty"`
					MobilePhone string `json:"mobilePhone,omitempty" yaml:",omitempty" xml:",omitempty"`
					Email       string `json:"email,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"contact,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"pickupParty,omitempty" yaml:",omitempty" xml:",omitempty"`
			CollectionParty *struct {
				Name    string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				Address *struct {
					Street1     string `json:"street1,omitempty" yaml:",omitempty" xml:",omitempty"`
					Street2     string `json:"street2,omitempty" yaml:",omitempty" xml:",omitempty"`
					City        string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
					CountryCode string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					Country     string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
					PostCode    string `json:"postCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"address,omitempty" yaml:",omitempty" xml:",omitempty"`
				Contact *struct {
					ContactName string `json:"contactName,omitempty" yaml:",omitempty" xml:",omitempty"`
					Phone       string `json:"phone,omitempty" yaml:",omitempty" xml:",omitempty"`
					MobilePhone string `json:"mobilePhone,omitempty" yaml:",omitempty" xml:",omitempty"`
					Email       string `json:"email,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"contact,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"collectionParty,omitempty" yaml:",omitempty" xml:",omitempty"`
			NotificationParty *struct {
				Name    string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				Address *struct {
					Street1     string `json:"street1,omitempty" yaml:",omitempty" xml:",omitempty"`
					Street2     string `json:"street2,omitempty" yaml:",omitempty" xml:",omitempty"`
					City        string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
					CountryCode string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					Country     string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
					PostCode    string `json:"postCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"address,omitempty" yaml:",omitempty" xml:",omitempty"`
				Contact *struct {
					ContactName string `json:"contactName,omitempty" yaml:",omitempty" xml:",omitempty"`
					Phone       string `json:"phone,omitempty" yaml:",omitempty" xml:",omitempty"`
					MobilePhone string `json:"mobilePhone,omitempty" yaml:",omitempty" xml:",omitempty"`
					Email       string `json:"email,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"contact,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"notificationParty,omitempty" yaml:",omitempty" xml:",omitempty"`
			CashOnDelivery *struct {
				Value    string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
				Currency string `json:"currency,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"cashOnDelivery,omitempty" yaml:",omitempty" xml:",omitempty"`
			StatusText *struct {
				Header                 string `json:"header,omitempty" yaml:",omitempty" xml:",omitempty"`
				Body                   string `json:"body,omitempty" yaml:",omitempty" xml:",omitempty"`
				EstimatedTimeOfArrival string `json:"estimatedTimeOfArrival,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"statusText,omitempty" yaml:",omitempty" xml:",omitempty"`
			Status                 string `json:"status,omitempty" yaml:",omitempty" xml:",omitempty"`
			RequestedDeliveryPoint *struct {
				Name           string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				Country        string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
				CountryCode    string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				LocationDetail string `json:"locationDetail,omitempty" yaml:",omitempty" xml:",omitempty"`
				Address        *struct {
					Street1     string `json:"street1,omitempty" yaml:",omitempty" xml:",omitempty"`
					Street2     string `json:"street2,omitempty" yaml:",omitempty" xml:",omitempty"`
					City        string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
					CountryCode string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					Country     string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
					PostCode    string `json:"postCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"address,omitempty" yaml:",omitempty" xml:",omitempty"`
				Contact *struct {
					ContactName string `json:"contactName,omitempty" yaml:",omitempty" xml:",omitempty"`
					Phone       string `json:"phone,omitempty" yaml:",omitempty" xml:",omitempty"`
					MobilePhone string `json:"mobilePhone,omitempty" yaml:",omitempty" xml:",omitempty"`
					Email       string `json:"email,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"contact,omitempty" yaml:",omitempty" xml:",omitempty"`
				Coordinate []*struct {
					SrID     string `json:"srId,omitempty" yaml:",omitempty" xml:",omitempty"`
					Northing string `json:"northing,omitempty" yaml:",omitempty" xml:",omitempty"`
					Easting  string `json:"easting,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"coordinate,omitempty" yaml:",omitempty" xml:",omitempty"`
				OpeningHour []*struct {
					OpenFrom  string `json:"openFrom,omitempty" yaml:",omitempty" xml:",omitempty"`
					OpenTo    string `json:"openTo,omitempty" yaml:",omitempty" xml:",omitempty"`
					OpenFrom2 string `json:"openFrom2,omitempty" yaml:",omitempty" xml:",omitempty"`
					OpenTo2   string `json:"openTo2,omitempty" yaml:",omitempty" xml:",omitempty"`
					Monday    bool   `json:"monday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Tuesday   bool   `json:"tuesday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Wednesday bool   `json:"wednesday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Thursday  bool   `json:"thursday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Friday    bool   `json:"friday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Saturday  bool   `json:"saturday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Sunday    bool   `json:"sunday,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"openingHour,omitempty" yaml:",omitempty" xml:",omitempty"`
				DisplayName      string `json:"displayName,omitempty" yaml:",omitempty" xml:",omitempty"`
				LocationID       string `json:"locationId,omitempty" yaml:",omitempty" xml:",omitempty"`
				Postcode         string `json:"postcode,omitempty" yaml:",omitempty" xml:",omitempty"`
				ServicePointType string `json:"servicePointType,omitempty" yaml:",omitempty" xml:",omitempty"`
				LocationType     string `json:"locationType,omitempty" yaml:",omitempty" xml:",omitempty"`
				City             string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"requestedDeliveryPoint,omitempty" yaml:",omitempty" xml:",omitempty"`
			DeliveryPoint *struct {
				Name           string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				Country        string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
				CountryCode    string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				LocationDetail string `json:"locationDetail,omitempty" yaml:",omitempty" xml:",omitempty"`
				Address        *struct {
					Street1     string `json:"street1,omitempty" yaml:",omitempty" xml:",omitempty"`
					Street2     string `json:"street2,omitempty" yaml:",omitempty" xml:",omitempty"`
					City        string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
					CountryCode string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					Country     string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
					PostCode    string `json:"postCode,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"address,omitempty" yaml:",omitempty" xml:",omitempty"`
				Contact *struct {
					ContactName string `json:"contactName,omitempty" yaml:",omitempty" xml:",omitempty"`
					Phone       string `json:"phone,omitempty" yaml:",omitempty" xml:",omitempty"`
					MobilePhone string `json:"mobilePhone,omitempty" yaml:",omitempty" xml:",omitempty"`
					Email       string `json:"email,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"contact,omitempty" yaml:",omitempty" xml:",omitempty"`
				Coordinate []*struct {
					SrID     string `json:"srId,omitempty" yaml:",omitempty" xml:",omitempty"`
					Northing string `json:"northing,omitempty" yaml:",omitempty" xml:",omitempty"`
					Easting  string `json:"easting,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"coordinate,omitempty" yaml:",omitempty" xml:",omitempty"`
				OpeningHour []*struct {
					OpenFrom  string `json:"openFrom,omitempty" yaml:",omitempty" xml:",omitempty"`
					OpenTo    string `json:"openTo,omitempty" yaml:",omitempty" xml:",omitempty"`
					OpenFrom2 string `json:"openFrom2,omitempty" yaml:",omitempty" xml:",omitempty"`
					OpenTo2   string `json:"openTo2,omitempty" yaml:",omitempty" xml:",omitempty"`
					Monday    bool   `json:"monday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Tuesday   bool   `json:"tuesday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Wednesday bool   `json:"wednesday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Thursday  bool   `json:"thursday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Friday    bool   `json:"friday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Saturday  bool   `json:"saturday,omitempty" yaml:",omitempty" xml:",omitempty"`
					Sunday    bool   `json:"sunday,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"openingHour,omitempty" yaml:",omitempty" xml:",omitempty"`
				DisplayName      string `json:"displayName,omitempty" yaml:",omitempty" xml:",omitempty"`
				LocationID       string `json:"locationId,omitempty" yaml:",omitempty" xml:",omitempty"`
				Postcode         string `json:"postcode,omitempty" yaml:",omitempty" xml:",omitempty"`
				ServicePointType string `json:"servicePointType,omitempty" yaml:",omitempty" xml:",omitempty"`
				LocationType     string `json:"locationType,omitempty" yaml:",omitempty" xml:",omitempty"`
				City             string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"deliveryPoint,omitempty" yaml:",omitempty" xml:",omitempty"`
			PaymentAccount *struct {
				AccountNumber string `json:"accountNumber,omitempty" yaml:",omitempty" xml:",omitempty"`
				AccountType   string `json:"accountType,omitempty" yaml:",omitempty" xml:",omitempty"`
				Bic           string `json:"bic,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"paymentAccount,omitempty" yaml:",omitempty" xml:",omitempty"`
			TotalWeight *struct {
				Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
				Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"totalWeight,omitempty" yaml:",omitempty" xml:",omitempty"`
			TotalVolume *struct {
				Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
				Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"totalVolume,omitempty" yaml:",omitempty" xml:",omitempty"`
			AssessedWeight *struct {
				Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
				Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"assessedWeight,omitempty" yaml:",omitempty" xml:",omitempty"`
			AssessedVolume *struct {
				Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
				Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"assessedVolume,omitempty" yaml:",omitempty" xml:",omitempty"`
			Items []*struct {
				ItemID                       string `json:"itemId,omitempty" yaml:",omitempty" xml:",omitempty"`
				OriginEstimatedTimeOfArrival string `json:"originEstimatedTimeOfArrival,omitempty" yaml:",omitempty" xml:",omitempty"`
				EstimatedTimeOfArrival       string `json:"estimatedTimeOfArrival,omitempty" yaml:",omitempty" xml:",omitempty"`
				RealTimeOfArrival            string `json:"realTimeOfArrival,omitempty" yaml:",omitempty" xml:",omitempty"`
				DropOffDate                  string `json:"dropOffDate,omitempty" yaml:",omitempty" xml:",omitempty"`
				DeliveryDate                 string `json:"deliveryDate,omitempty" yaml:",omitempty" xml:",omitempty"`
				ReturnDate                   string `json:"returnDate,omitempty" yaml:",omitempty" xml:",omitempty"`
				TypeOfItem                   string `json:"typeOfItem,omitempty" yaml:",omitempty" xml:",omitempty"`
				TypeOfItemName               string `json:"typeOfItemName,omitempty" yaml:",omitempty" xml:",omitempty"`
				TypeOfItemActual             string `json:"typeOfItemActual,omitempty" yaml:",omitempty" xml:",omitempty"`
				TypeOfItemActualName         string `json:"typeOfItemActualName,omitempty" yaml:",omitempty" xml:",omitempty"`
				AdditionalInformation        string `json:"additionalInformation,omitempty" yaml:",omitempty" xml:",omitempty"`
				NoItems                      string `json:"noItems,omitempty" yaml:",omitempty" xml:",omitempty"`
				NumberOfPallets              string `json:"numberOfPallets,omitempty" yaml:",omitempty" xml:",omitempty"`
				Signature                    string `json:"signature,omitempty" yaml:",omitempty" xml:",omitempty"`
				Status                       string `json:"status,omitempty" yaml:",omitempty" xml:",omitempty"`
				EventStatus                  string `json:"eventStatus,omitempty" yaml:",omitempty" xml:",omitempty"`
				StatusText                   *struct {
					Header                 string `json:"header,omitempty" yaml:",omitempty" xml:",omitempty"`
					Body                   string `json:"body,omitempty" yaml:",omitempty" xml:",omitempty"`
					EstimatedTimeOfArrival string `json:"estimatedTimeOfArrival,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"statusText,omitempty" yaml:",omitempty" xml:",omitempty"`
				Acceptor *struct {
					SignatureReference string `json:"signatureReference,omitempty" yaml:",omitempty" xml:",omitempty"`
					Name               string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"acceptor,omitempty" yaml:",omitempty" xml:",omitempty"`
				StatedMeasurement *struct {
					Weight *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"weight,omitempty" yaml:",omitempty" xml:",omitempty"`
					Length *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"length,omitempty" yaml:",omitempty" xml:",omitempty"`
					Height *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"height,omitempty" yaml:",omitempty" xml:",omitempty"`
					Width *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"width,omitempty" yaml:",omitempty" xml:",omitempty"`
					Volume *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"volume,omitempty" yaml:",omitempty" xml:",omitempty"`
					Compass *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"compass,omitempty" yaml:",omitempty" xml:",omitempty"`
					CompassPlusLength *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"compassPlusLength,omitempty" yaml:",omitempty" xml:",omitempty"`
					Bag *struct {
						Bag bool `json:"bag,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"bag,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"statedMeasurement,omitempty" yaml:",omitempty" xml:",omitempty"`
				AssessedMeasurement *struct {
					Weight *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"weight,omitempty" yaml:",omitempty" xml:",omitempty"`
					Length *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"length,omitempty" yaml:",omitempty" xml:",omitempty"`
					Height *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"height,omitempty" yaml:",omitempty" xml:",omitempty"`
					Width *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"width,omitempty" yaml:",omitempty" xml:",omitempty"`
					Volume *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"volume,omitempty" yaml:",omitempty" xml:",omitempty"`
					Compass *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"compass,omitempty" yaml:",omitempty" xml:",omitempty"`
					CompassPlusLength *struct {
						Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
						Unit  string `json:"unit,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"compassPlusLength,omitempty" yaml:",omitempty" xml:",omitempty"`
					Bag *struct {
						Bag bool `json:"bag,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"bag,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"assessedMeasurement,omitempty" yaml:",omitempty" xml:",omitempty"`
				Events []*struct {
					EventTime        string `json:"eventTime,omitempty" yaml:",omitempty" xml:",omitempty"`
					EventCode        string `json:"eventCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					Status           string `json:"status,omitempty" yaml:",omitempty" xml:",omitempty"`
					EventDescription string `json:"eventDescription,omitempty" yaml:",omitempty" xml:",omitempty"`
					Location         *struct {
						LocationID       string `json:"locationId,omitempty" yaml:",omitempty" xml:",omitempty"`
						DisplayName      string `json:"displayName,omitempty" yaml:",omitempty" xml:",omitempty"`
						Name             string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
						CountryCode      string `json:"countryCode,omitempty" yaml:",omitempty" xml:",omitempty"`
						Country          string `json:"country,omitempty" yaml:",omitempty" xml:",omitempty"`
						Postcode         string `json:"postcode,omitempty" yaml:",omitempty" xml:",omitempty"`
						City             string `json:"city,omitempty" yaml:",omitempty" xml:",omitempty"`
						ServicePointType string `json:"servicePointType,omitempty" yaml:",omitempty" xml:",omitempty"`
						LocationType     string `json:"locationType,omitempty" yaml:",omitempty" xml:",omitempty"`
					} `json:"location,omitempty" yaml:",omitempty" xml:",omitempty"`
					LocalEventCode string `json:"localEventCode,omitempty" yaml:",omitempty" xml:",omitempty"`
					ScanUserID     string `json:"scanUserId,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"events,omitempty" yaml:",omitempty" xml:",omitempty"`
				References []*struct {
					Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
					Type  string `json:"type,omitempty" yaml:",omitempty" xml:",omitempty"`
					Name  string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"references,omitempty" yaml:",omitempty" xml:",omitempty"`
				ItemRefID []*struct {
					ReferenceID string `json:"referenceId,omitempty" yaml:",omitempty" xml:",omitempty"`
					Type        string `json:"type,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"itemRefId" yaml:",omitempty" xml:",omitempty"`
				FreeText []*struct {
					Text string `json:"text,omitempty" yaml:",omitempty" xml:",omitempty"`
					Type string `json:"type,omitempty" yaml:",omitempty" xml:",omitempty"`
				} `json:"freeText" yaml:",omitempty" xml:",omitempty"`
			} `json:"items,omitempty" yaml:",omitempty" xml:",omitempty"`
			AdditionalServices []*struct {
				Code         string `json:"code,omitempty" yaml:",omitempty" xml:",omitempty"`
				SourceSystem string `json:"sourceSystem,omitempty" yaml:",omitempty" xml:",omitempty"`
				NameKey      string `json:"nameKey,omitempty" yaml:",omitempty" xml:",omitempty"`
				Name         string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"additionalServices,omitempty" yaml:",omitempty" xml:",omitempty"`
			SplitStatuses []struct {
				NoItemsWithStatus string `json:"noItemsWithStatus,omitempty" yaml:",omitempty" xml:",omitempty"`
				NoItems           string `json:"noItems,omitempty" yaml:",omitempty" xml:",omitempty"`
				StatusDescription string `json:"statusDescription,omitempty" yaml:",omitempty" xml:",omitempty"`
				Status            string `json:"status,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"splitStatuses" yaml:",omitempty" xml:",omitempty"`
			ShipmentReferences []*struct {
				Value string `json:"value,omitempty" yaml:",omitempty" xml:",omitempty"`
				Type  string `json:"type,omitempty" yaml:",omitempty" xml:",omitempty"`
				Name  string `json:"name,omitempty" yaml:",omitempty" xml:",omitempty"`
			} `json:"shipmentReferences,omitempty" yaml:",omitempty" xml:",omitempty"`
		} `json:"shipments,omitempty" yaml:",omitempty" xml:",omitempty"`
	} `json:"trackingInformationResponse,omitempty" yaml:",omitempty" xml:",omitempty"`
}

func (r *Response) JSON() ([]byte, error) {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}

func (r *Response) YAML() ([]byte, error) {
	b, err := yaml.Marshal(r)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}

func (r *Response) XML() ([]byte, error) {
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
