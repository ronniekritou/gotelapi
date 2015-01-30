package telapi

import (
	"encoding/json"
	"errors"
	// "strconv"
	"fmt"
)

type CallData struct {
	Call Call
}

type Call struct {
	Sid             string
	DataCreated     string
	DateUpdated     string
	ParentCallSid   string
	AccountSid      string
	From            string
	To              string
	PhoneNumberSid  string
	Status          string
	StartTime       string
	EndTime         string
	Price           string //Maybe should be float value is like 0.01000
	Direction       string
	AnsweredBy      string
	ApiVersion      string
	ForwardedFrom   string
	Duration        float64
	CallerIdBlocked bool
}

//minimal caller, optional map being passed but not implemented
func (helper TelapiHelper) MakeCall(from, to, url string, optional map[string]interface{}) (*Call, error) {
	if from == "" || to == "" || url == "" {
		return nil, errors.New("Missing needed From, To, or Url")
	}

	data := map[string]string{
		"To":   to,
		"From": from,
		"Url":  url,
	}

	if optional, ok := optional["HideCallerId"].(bool); ok {
		data["HideCallerId"] = fmt.Sprintf("%v", optional)
	}

	resp, err := helper.PostRequest("/Calls", data)

	if err != nil {
		return nil, err
	}

	call := new(CallData)

	if err = json.Unmarshal(*resp, &call); err != nil {
		return nil, err
	}

	return &call.Call, nil

}

func (helper TelapiHelper) ViewCall(call_sid string) (*Call, error) {
	if call_sid == "" {
		return nil, errors.New("Missing required call sid.")
	}

	response, err := helper.GetRequest("/Calls/"+call_sid, nil)

	if err != nil {
		return nil, err
	}

	call := new(Call)

	if err = json.Unmarshal(*response, &call); err != nil {
		return nil, err
	}
	return call, nil

}

func (helper TelapiHelper) RecordCall(call_sid string, params map[string]string) error {
	if call_sid == "" {
		return errors.New("Missing required call sid.")
	}

	params["Record"] = "true"

	_, err := helper.PostRequest("/Calls/"+call_sid+"/Recordings", params)

	if err != nil {
		return err
	}

	return nil
}
