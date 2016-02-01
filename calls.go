package telapi

import (
	"encoding/json"
	"errors"
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

type CallOptions struct { //Needs to be added with more values
	HideCallerId bool
}

func (callOptions CallOptions) ToMap() map[string]string { //needs to be updated as the struct is updated
	callOptionsMap := map[string]string{}

	callOptionsMap["HideCallerId"] = fmt.Sprintf("%v", callOptions.HideCallerId)

	return callOptionsMap
}

//minimal caller, optional map being passed but not implemented
func (helper TelapiHelper) MakeCall(from, to, url string, options *CallOptions) (*Call, error) {
	if from == "" || to == "" || url == "" {
		return nil, errors.New("Missing needed From, To, or Url")
	}

	data := map[string]string{
		"To":   to,
		"From": from,
		"Url":  url,
	}

	if options != nil {

		dataMap := options.ToMap()

		for k, v := range dataMap {
			data[k] = v
		}
	}

	resp, err := helper.PostRequest("/Calls", data)

	if err != nil {
		return nil, err
	}

	call := new(Call)

	if err = json.Unmarshal(*resp, &call); err != nil {
		return nil, err
	}

	return call, nil

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

//Actions that may be done , on on-going calls.

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

func (helper TelapiHelper) PlayAudioCall(call_sid string, params map[string]string) error {
	if call_sid == "" {
		return errors.New("Missing required call sid.")
	}

	_, err := helper.PostRequest("/Calls/"+call_sid+"/Play", params)

	if err != nil {
		return err
	}

	return nil
}

func (helper TelapiHelper) InterruptCall(call_sid string, params map[string]string) error {
	if call_sid == "" {
		return errors.New("Missing required call sid.")
	}

	fmt.Println("HERE ?")

	_, err := helper.PostRequestv2("/Calls/"+call_sid, params)

	if err != nil {
		return err
	}

	fmt.Println("HERE2 ?", err)

	return nil
}
