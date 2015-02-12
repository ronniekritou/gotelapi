package telapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ConferenceData struct {
	Conference Conference
}

type Conference struct {
	Sid                     string
	DataCreated             string
	DateUpdated             string
	FriendlyName            string
	ActiveParticipantsCount int
	Uri                     string
	Status                  string
}

func (helper TelapiHelper) FindConferenceByFriendlyName(friendlyName string) (*Conference, error) {
	if friendlyName == "" {
		return nil, errors.New("Missing friendlyName!")
	}

	resp, err := helper.PostRequest(fmt.Sprintf("/Conferences/%s", friendlyName), nil)

	if err != nil {
		return nil, err
	}

	conference := new(ConferenceData)

	if err = json.Unmarshal(*resp, &conference); err != nil {
		return nil, err
	}

	return &conference.Conference, nil

}
