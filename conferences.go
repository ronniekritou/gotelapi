package telapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ConferenceData struct {
	Conferences []Conference
}

type Conference struct {
	Sid                     string
	DataCreated             string
	DateUpdated             string
	FriendlyName            string
	ActiveParticipantsCount int `json:"active_participants_count"`
	Uri                     string
	Status                  string
}

func (helper TelapiHelper) FindConferenceByFriendlyName(friendlyName string) (*Conference, error) {
	if friendlyName == "" {
		return nil, errors.New("Missing friendlyName!")
	}

	fmt.Println(friendlyName)
	data := map[string]string{
		"FriendlyName": friendlyName,
	}

	resp, err := helper.GetRequestWithParamsAdded(fmt.Sprintf("/Conferences"), data)

	if err != nil {
		return nil, err
	}

	conference := new(ConferenceData)

	if err = json.Unmarshal(*resp, &conference); err != nil {
		return nil, err
	}

	if len(conference.Conferences) < 1 {
		return nil, errors.New("Could no locate conference by friendlyname")
	}

	return &conference.Conferences[0], nil

}
