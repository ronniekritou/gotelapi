package telapi

import (
	"errors"
	"net/http"
)

type TelapiHelper struct {
	Sid       string
	AuthToken string
	client    *http.Client
}

func CreateClient(sid string, auth_token string, client *http.Client) (TelapiHelper, error) {
	if client == nil {
		client = http.DefaultClient
	}

	if sid == "" || auth_token == "" {
		return TelapiHelper{"", "", client}, errors.New("missing sid or auth token")
	}

	telapi_helper := TelapiHelper{sid, auth_token, client}

	resp, err := telapi_helper.PostRequest("", nil)
	if err != nil {
		return TelapiHelper{"", "", client}, err
	}

	if resp == nil {
		return TelapiHelper{"", "", client}, errors.New("should have received a response")
	}

	return telapi_helper, nil
}
