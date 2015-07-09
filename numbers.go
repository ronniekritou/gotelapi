package telapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

type IncomingPhoneNumbersList struct {
	FirstPageURI         string                `json:"first_page_uri"`
	End                  int                   `json:"end"`
	Total                int                   `json:"total"`
	PreviousPageURI      string                `json:"previous_page_uri"`
	NumPages             int                   `json:"num_pages"`
	IncomingPhoneNumbers []IncomingPhoneNumber `json:"incoming_phone_numbers"`
	URI                  string                `json:"uri"`
	PageSize             int                   `json:"page_size"`
	Start                int                   `json:"start"`
	NextPageURI          string                `json:"next_page_uri"`
	LastPageURI          string                `json:"last_page_uri"`
	Page                 int                   `json:"page"`
}

type IncomingPhoneNumber struct {
	DateUpdated         string `json:"date_updated"`
	VoiceURL            string `json:"voice_url"`
	VoiceFallbackMethod string `json:"voice_fallback_method"`
	Capabilities        struct {
		Voice string `json:"voice"`
		Sms   string `json:"sms"`
	} `json:"capabilities"`
	Sid                  string      `json:"sid"`
	HeartbeatMethod      string      `json:"heartbeat_method"`
	Type                 string      `json:"type"`
	StatusCallbackMethod string      `json:"status_callback_method"`
	VoiceFallbackURL     string      `json:"voice_fallback_url"`
	PhoneNumber          string      `json:"phone_number"`
	HangupCallback       interface{} `json:"hangup_callback"`
	HangupCallbackMethod string      `json:"hangup_callback_method"`
	HeartbeatURL         interface{} `json:"heartbeat_url"`
	SmsURL               string      `json:"sms_url"`
	VoiceMethod          string      `json:"voice_method"`
	VoiceCallerIDLookup  string      `json:"voice_caller_id_lookup"`
	FriendlyName         string      `json:"friendly_name"`
	URI                  string      `json:"uri"`
	SmsFallbackURL       string      `json:"sms_fallback_url"`
	AccountSid           string      `json:"account_sid"`
	SmsMethod            string      `json:"sms_method"`
	APIVersion           string      `json:"api_version"`
	SmsFallbackMethod    string      `json:"sms_fallback_method"`
	NextRenewalDate      string      `json:"next_renewal_date"`
	DateCreated          string      `json:"date_created"`
	StatusCallback       interface{} `json:"status_callback"`
}

func (helper TelapiHelper) GetAllIncomingNumbers() (*[]IncomingPhoneNumber, error) {

	resp, err := helper.GetRequest(fmt.Sprintf("/IncomingPhoneNumbers"), nil)

	if err != nil {
		return nil, err
	}

	incomingPhoneNumbers := []IncomingPhoneNumber{}

	// Original list of phone numbers we are going to use the data from this one to populate the actual list
	phoneNumberList := new(IncomingPhoneNumbersList)

	if err = json.Unmarshal(*resp, &phoneNumberList); err != nil {
		return nil, err
	}

	// Append our original list of numbers
	incomingPhoneNumbers = append(incomingPhoneNumbers, phoneNumberList.IncomingPhoneNumbers...)

	lastEquals := strings.LastIndex(phoneNumberList.LastPageURI, "=")

	lastPageSize := phoneNumberList.LastPageURI[lastEquals+1:]

	if phoneNumberList.NumPages > 0 {

		for i := 1; i < phoneNumberList.NumPages; i++ {

			data := map[string]string{
				"Page":     fmt.Sprintf("%d", i),
				"PageSize": fmt.Sprintf("%d", phoneNumberList.PageSize),
			}

			if i == phoneNumberList.NumPages-1 {
				fmt.Println("Using last page size", lastPageSize)
				data["pageSize"] = fmt.Sprintf("%s", lastPageSize)

			}

			// get the correct page number next if there is one
			resp, err := helper.GetRequestWithParamsAdded(fmt.Sprintf("/IncomingPhoneNumbers"), data)

			if err != nil {
				continue
			}

			// make it into a struct
			newPhoneNumberList := new(IncomingPhoneNumbersList)

			if err = json.Unmarshal(*resp, &newPhoneNumberList); err != nil {
				continue
			}

			incomingPhoneNumbers = append(incomingPhoneNumbers, newPhoneNumberList.IncomingPhoneNumbers...)

		}

	}

	return &incomingPhoneNumbers, nil

}
