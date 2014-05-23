package telapi

import (
	"errors"
	"strings"
)

func (helper TelapiHelper) CarrierLookup(phone_number string) (string, error) {
	if phone_number == "" {
		return "", errors.New("Missing required voicemail sid.")
	}

	data := map[string]string{
		"PhoneNumber": phone_number,
	}

	response, err := helper.PostRequest("/Lookups/Carrier", data)

	if err != nil {
		return "", err
	}

	response_list := response["carrier_lookups"]
	carrierList := response_list.([]interface{})
	carrierData := carrierList[0].(map[string]interface{})

	carrier := carrierData["network"].(string)

	return strings.Split(carrier, " ")[0], nil

}

func (helper TelapiHelper) BnaLookup(phone_number string) (map[string]interface{}, error) {
	if phone_number == "" {
		return nil, errors.New("Missing required voicemail sid.")
	}

	data := map[string]string{
		"PhoneNumber": phone_number,
	}

	response, err := helper.PostRequest("/Lookups/Bna", data)

	if err != nil {
		return nil, err
	}

	response_list := response["bna_lookups"]
	bna_list := response_list.([]interface{})
	bna_data := bna_list[0].(map[string]interface{})

	// carrier := bnaData["network"].(string)

	return bna_data, nil

}

/*
BNA RETURNS AS

map[
	date_created:Thu, 22 May 2014 16:14:01 -0000
	price:0.00
	api_version:v2
	city:DETROIT
	state:MI
	phone_number:+4698694768
	sid:BL6c8890844ce10910008c4bf49ffd9420
	date_updated:Thu, 22 May 2014 16:14:02 -0000
	uri:/v2/Accounts/AC1d530461c32a4840a1a19183d0a0bb8c/BNA/BL6c8890844ce10910008c4bf49ffd9420
	account_sid:AC1d530461c32a4840a1a19183d0a0bb8c
	country_code:US]

*/

func (helper TelapiHelper) CnamLookup(phone_number string) (map[string]interface{}, error) {
	if phone_number == "" {
		return nil, errors.New("Missing required voicemail sid.")
	}

	data := map[string]string{
		"PhoneNumber": phone_number,
	}

	response, err := helper.PostRequest("/Lookups/Cnam", data)

	if err != nil {
		return nil, err
	}

	response_list := response["cnam_lookups"]
	cnam_list := response_list.([]interface{})
	cnam_data := cnam_list[0].(map[string]interface{})

	// carrier := bnaData["network"].(string)

	return cnam_data, nil

}
