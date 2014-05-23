package telapi

import (
	"errors"
)

func (helper TelapiHelper) ViewCall(call_sid string) (map[string]interface{}, error) {
	if call_sid == "" {
		return nil, errors.New("Missing required call sid.")
	}

	response, err := helper.GetRequest("/Calls/"+call_sid, nil)

	if err != nil {
		return nil, err
	}

	return response, nil

}
