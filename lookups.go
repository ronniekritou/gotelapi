package telapi

import (
	"encoding/json"
	"errors"
)

type CarrierLookupData struct {
	CarrierLookups []CarrierLookup `json:"carrier_lookups"`
}

type CarrierLookup struct {
	Sid         string
	DataCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	AccountSid  string `json:"account_sid"`
	PhoneNumber string `json:"phone_number"`
	Network     string
	Mobile      bool    `json:"mobile"`
	CarrierId   float64 `json:"carrier_id"`
	CountryCode string  `json:"country_code"`
	Mnc         string
	Mcc         string
	Price       string
	ApiVersion  string
	Uri         string
}

type CnamLookupData struct {
	CnamLookups []CnamLookup `json:"cnam_lookups"`
}

type CnamLookup struct {
	Sid         string
	DataCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	AccountSid  string `json:"account_sid"`
	PhoneNumber string `json:"phone_number"`
	Body        string
	Price       string
	ApiVersion  string
	Uri         string
}

type BnaLookupData struct {
	BnaLookups []BnaLookup `json:"bna_lookups"`
}

type BnaLookup struct {
	Sid         string
	DataCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	AccountSid  string `json:"account_sid"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string
	City        string
	State       string
	ZipCode     string `json:"zip_code"`
	CountryCode string `json:"country_code"`
	Price       string
	ApiVersion  string
	Uri         string
}

func (helper TelapiHelper) CarrierLookup(phone_number string) (*CarrierLookup, error) {
	if phone_number == "" {
		return nil, errors.New("Missing required voicemail sid.")
	}

	data := map[string]string{
		"PhoneNumber": phone_number,
	}

	resp, err := helper.PostRequest("/Lookups/Carrier", data)

	if err != nil {
		return nil, err
	}

	carrier := new(CarrierLookupData)

	if err = json.Unmarshal(*resp, &carrier); err != nil {
		return nil, err
	}

	return &carrier.CarrierLookups[0], nil

}

func (helper TelapiHelper) BnaLookup(phone_number string) (*BnaLookup, error) {
	if phone_number == "" {
		return nil, errors.New("Missing required voicemail sid.")
	}

	data := map[string]string{
		"PhoneNumber": phone_number,
	}

	resp, err := helper.PostRequest("/Lookups/Bna", data)

	if err != nil {
		return nil, err
	}

	bna := new(BnaLookupData)

	if err = json.Unmarshal(*resp, &bna); err != nil {
		return nil, err
	}

	return &bna.BnaLookups[0], nil

}

func (helper TelapiHelper) CnamLookup(phone_number string) (*CnamLookup, error) {
	if phone_number == "" {
		return nil, errors.New("Missing required voicemail sid.")
	}

	data := map[string]string{
		"PhoneNumber": phone_number,
	}

	resp, err := helper.PostRequest("/Lookups/Cnam", data)

	if err != nil {
		return nil, err
	}

	cnam := new(CnamLookupData)

	if err = json.Unmarshal(*resp, &cnam); err != nil {
		return nil, err
	}

	return &cnam.CnamLookups[0], nil

}
