package telapi

import (
	"errors"
)

func (helper TelapiHelper) TranscribeRecording(vm_sid string) (map[string]interface{}, error) {
	if vm_sid == "" {
		return nil, errors.New("Missing required voicemail sid.")
	}

	response, err := helper.PostRequest("/Recordings/"+vm_sid+"/Transcriptions", nil)

	if err != nil {
		return nil, err
	}

	return response, nil

}
