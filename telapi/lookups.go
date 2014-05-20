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
