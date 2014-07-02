package telapi

import (
	"encoding/json"
	"errors"
)

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
