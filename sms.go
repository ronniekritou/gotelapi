package telapi

import (
	"encoding/json"
	"errors"
)

type SmsMessage struct {
	Sid         string
	Body        string
	Status      string
	From        string
	To          string
	Direction   string
	DateUpdated string
	Price       string
	Uri         string
	AccountSid  string
	DateSent    string
	DateCreated string
	ApiVersion  string
}

func (helper TelapiHelper) SendSMS(to string, from string, body string) (*SmsMessage, error) {

	if to == "" || from == "" {
		return nil, errors.New("Missing required To or From.")
	}

	data := map[string]string{
		"To":            to,
		"From":          from,
		"Body":          body,
		"AllowMultiple": "true",
	}

	resp, err := helper.PostRequest("/SMS/Messages", data)
	if err != nil {
		return nil, err
	}

	//Lets unmarshal our response
	var f interface{}
	err = json.Unmarshal(*resp, &f)
	if err != nil {
		return nil, err
	}

	data_map := f.(map[string]interface{})
	response_list := data_map["sms_messages"]
	messageList := response_list.([]interface{}) //array of interfaces
	msg_data, err := json.Marshal(messageList[0].(map[string]interface{}))

	if err != nil {
		return nil, err
	}

	sms := new(SmsMessage)

	if err = json.Unmarshal(msg_data, &sms); err != nil {
		return nil, err
	}

	return sms, nil
}
