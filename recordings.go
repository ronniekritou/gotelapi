package telapi

import (
	"encoding/json"
)

type RecordingData struct {
	Recordings []Recording
}

type Recording struct {
	Sid          string `json:"sid"`
	CallSid      string `json:"call_sid"`
	Duration     int    `json:"duration"`
	DateCreated  string `json:"date_created"`
	DateUpdated  string `json:"date_updated"`
	ApiVersion   string `json:"api_version"`
	RecordingUrl string `json:"recording_url"`
	Uri          string `json:"uri"`
}

func (helper TelapiHelper) ListRecordings(params map[string]string) ([]Recording, error) {

	resp, err := helper.GetRequestWithParamsAdded("/Recordings", params)
	if err != nil {
		return nil, err
	}

	recordingData := &RecordingData{}
	if err = json.Unmarshal(*resp, &recordingData); err != nil {
		return nil, err
	}
	return recordingData.Recordings, nil
}
