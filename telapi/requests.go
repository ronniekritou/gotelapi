package telapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	// "net/url"
)

func (helper TelapiHelper) TelapiRequest(method string, urlStr string, params map[string]string) (map[string]interface{}, error) {

	data := DataMapToUrlValues(params)

	client := &http.Client{}
	req, err := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(helper.sid, helper.auth_token)
	resp, err := client.Do(req)

	if resp.StatusCode == 404 {
		return nil, errors.New("Statuscode was 404 because : " + resp.Status)
	} else if resp.StatusCode != 200 {
		return nil, errors.New("Unexpected status code returned." + resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var f interface{}

	err = json.Unmarshal(bodyBytes, &f)
	if err != nil {
		return nil, err
	}

	m := f.(map[string]interface{})

	return m, nil

}

func (helper TelapiHelper) PostRequest(uri string, param_data map[string]string) (map[string]interface{}, error) {

	urlStr := fmt.Sprintf("https://api.telapi.com/v1/Accounts/%s%s.json", helper.sid, uri)

	resp, err := helper.TelapiRequest("POST", urlStr, param_data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (helper TelapiHelper) GetRequest(uri string, param_data map[string]string) (map[string]interface{}, error) {

	urlStr := fmt.Sprintf("https://api.telapi.com/v1/Accounts/%s%s.json", helper.sid, uri)

	resp, err := helper.TelapiRequest("GET", urlStr, param_data)
	if err != nil {
		return nil, err
	}

	return resp, nil

}
