package telapi

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	// "net/url"
)

func (helper TelapiHelper) TelapiRequest(method string, urlStr string, params map[string]string) (*[]byte, error) {

	data := DataMapToUrlValues(params)

	var (
		resp       *http.Response
		maxRetries int
	)

	maxRetries = 1

	for i := 1; i <= maxRetries; i++ {
		client := &http.Client{}
		req, err := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
		if err != nil {
			return nil, err
		}

		if helper.Sid == "" || helper.AuthToken == "" {
			return nil, errors.New("We are missing ether the Sid or authtoken " + helper.Sid + ":" + helper.AuthToken)
		}

		req.SetBasicAuth(helper.Sid, helper.AuthToken)
		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			break
		}

		if maxRetries == i {
			return nil, errors.New("Unexpected status code returned." + resp.Status)
		}

	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &bodyBytes, nil

}

func (helper TelapiHelper) PostRequest(uri string, param_data map[string]string) (*[]byte, error) {

	urlStr := fmt.Sprintf("https://api.telapi.com/v1/Accounts/%s%s.json", helper.Sid, uri)

	resp, err := helper.TelapiRequest("POST", urlStr, param_data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (helper TelapiHelper) GetRequest(uri string, param_data map[string]string) (*[]byte, error) {

	urlStr := fmt.Sprintf("https://api.telapi.com/v1/Accounts/%s%s.json", helper.Sid, uri)

	resp, err := helper.TelapiRequest("GET", urlStr, param_data)
	if err != nil {
		return nil, err
	}

	return resp, nil

}
