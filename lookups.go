package telapi

import (
	"encoding/json"
	"errors"
)

type CarrierLookup struct {
	Sid         string
	DataCreated string
	DateUpdated string
	AccountSid  string
	PhoneNumber string
	Network     string
	Mobile      bool `json:",string"`
	CarrierId   float64
	CountryCode string
	Mnc         string
	Mcc         string
	Price       string
	ApiVersion  string
	Uri         string
}

type CnamLookup struct {
	Sid         string
	DataCreated string
	DateUpdated string
	AccountSid  string
	PhoneNumber string
	Body        string
	Price       string
	ApiVersion  string
	Uri         string
}

type BnaLookup struct {
	Sid         string
	DataCreated string
	DateUpdated string
	AccountSid  string
	PhoneNumber string
	FirstName   string
	LastName    string
	Address     string
	City        string
	State       string
	ZipCode     string
	CountryCode string
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

	//Lets unmarshal our response
	var f interface{}
	err = json.Unmarshal(*resp, &f)
	if err != nil {
		return nil, err
	}

	//Since it returns us a map with an array as the first element, we have to parse it out
	data_map := f.(map[string]interface{})
	response_list := data_map["carrier_lookups"]
	carrierList := response_list.([]interface{}) //array of interfaces
	carrier_data, err := json.Marshal(carrierList[0].(map[string]interface{}))
	// make it back into bytes so we can apply attributes

	if err != nil {
		return nil, err
	}

	carrier := new(CarrierLookup)

	if err = json.Unmarshal(carrier_data, &carrier); err != nil {
		return nil, err
	}
	return carrier, nil

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

	//Lets unmarshal our response
	var f interface{}
	err = json.Unmarshal(*resp, &f)
	if err != nil {
		return nil, err
	}

	//Since it returns us a map with an array as the first element, we have to parse it out
	data_map := f.(map[string]interface{})
	response_list := data_map["bna_lookups"]
	bnaList := response_list.([]interface{}) //array of interfaces
	bna_data, err := json.Marshal(bnaList[0].(map[string]interface{}))
	// make it back into bytes so we can apply attributes

	bna := new(BnaLookup)

	if err = json.Unmarshal(bna_data, &bna); err != nil {
		return nil, err
	}
	return bna, nil
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

	//Lets unmarshal our response
	var f interface{}
	err = json.Unmarshal(*resp, &f)
	if err != nil {
		return nil, err
	}

	//Since it returns us a map with an array as the first element, we have to parse it out
	data_map := f.(map[string]interface{})
	response_list := data_map["cnam_lookups"]
	cnamList := response_list.([]interface{}) //array of interfaces
	cnam_data, err := json.Marshal(cnamList[0].(map[string]interface{}))
	// make it back into bytes so we can apply attributes

	cnam := new(CnamLookup)

	if err = json.Unmarshal(cnam_data, &cnam); err != nil {
		return nil, err
	}
	return cnam, nil

}
