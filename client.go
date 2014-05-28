package telapi

import (
	"errors"
)

type TelapiHelper struct {
	sid        string
	auth_token string
}

func CreateClient(sid string, auth_token string) (TelapiHelper, error) {

	if sid == "" || auth_token == "" {
		return TelapiHelper{"", ""}, errors.New("Missing sid or auth token.")
	}

	telapi_helper := TelapiHelper{sid, auth_token}

	resp, err := telapi_helper.PostRequest("", nil)
	if err != nil {
		return TelapiHelper{"", ""}, err
	}

	if resp == nil {
		return TelapiHelper{"", ""}, errors.New("Should have recieved a response.")
	}

	return telapi_helper, nil
}
