package telapi

import (
	"errors"
)

func (helper TelapiHelper) TranscribeRecording(vm_sid string, callback_url string) (map[string]interface{}, error) {
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

	return response, nil

}

func (helper TelapiHelper) TranscribeAudioUrl(audio_url string, callback_url string) (map[string]interface{}, error) {
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

	return response, nil

}

func (helper TelapiHelper) ViewTranscription(tr_sid string) (map[string]interface{}, error) {
	if tr_sid == "" {
		return nil, errors.New("Missing required transcription sid.")
	}

	response, err := helper.GetRequest("/Transcriptions/"+tr_sid, nil)

	if err != nil {
		return nil, err

	}

	return response, nil

}
