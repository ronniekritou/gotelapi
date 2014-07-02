package telapi

import (
	"encoding/json"
	"errors"
)

type Transcription struct {
	Sid                   string
	DataCreated           string
	DateUpdated           string
	ParentCallSid         string
	AccountSid            string
	Status                string
	Type                  string
	AudioUrl              string
	RecordingSid          string //Maybe should be float value is like 0.01000
	Duration              float64
	TranscriptionText     string
	ApiVersion            string
	Price                 string
	TranscriptionCallback string
	CallbackMethod        string
	Uri                   string
}

func (helper TelapiHelper) TranscribeRecording(vm_sid string, callback_url string) (*Transcription, error) {
	if vm_sid == "" {
		return nil, errors.New("Missing required voicemail sid.")
	}

	data := map[string]string{
		"TranscribeCallback": callback_url,
	}

	response, err := helper.PostRequest("/Recordings/"+vm_sid+"/Transcriptions", data)

	if err != nil {
		return nil, err
	}

	transcription := new(Transcription)

	if err = json.Unmarshal(*response, &transcription); err != nil {
		return nil, err
	}
	return transcription, nil

}

func (helper TelapiHelper) TranscribeAudioUrl(audio_url string, callback_url string) (*Transcription, error) {
	if audio_url == "" {
		return nil, errors.New("Missing required audio url.")
	}

	data := map[string]string{
		"AudioUrl":           audio_url,
		"Quality":            "auto",
		"TranscribeCallback": callback_url,
	}

	response, err := helper.PostRequest("/Transcriptions", data)

	if err != nil {
		return nil, err
	}

	transcription := new(Transcription)

	if err = json.Unmarshal(*response, &transcription); err != nil {
		return nil, err
	}
	return transcription, nil

}

func (helper TelapiHelper) ViewTranscription(tr_sid string) (*Transcription, error) {
	if tr_sid == "" {
		return nil, errors.New("Missing required transcription sid.")
	}

	response, err := helper.GetRequest("/Transcriptions/"+tr_sid, nil)

	if err != nil {
		return nil, err

	}

	transcription := new(Transcription)

	if err = json.Unmarshal(*response, &transcription); err != nil {
		return nil, err
	}
	return transcription, nil

}
