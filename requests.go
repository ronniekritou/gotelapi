package telapi

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (helper TelapiHelper) TelapiRequest(method string, urlStr string, params map[string]string) (*[]byte, error) {
	data := DataMapToUrlValues(params)

	if helper.Sid == "" || helper.AuthToken == "" {
		return nil, errors.New("we are missing ether the Sid or authtoken " + helper.Sid + ":" + helper.AuthToken)
	}

	if helper.client == nil {
		helper.client = http.DefaultClient
	}

	req, err := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(helper.Sid, helper.AuthToken)
	resp, err := helper.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("unexpected status code returned: %d \nError was: %s", resp.StatusCode, resp.Status))
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

func (helper TelapiHelper) PostRequestv2(uri string, param_data map[string]string) (*[]byte, error) {

	urlStr := fmt.Sprintf("https://api.telapi.com/v2/Accounts/%s%s.json", helper.Sid, uri)

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

func (helper TelapiHelper) GetRequestWithParamsAdded(uri string, param_data map[string]string) (*[]byte, error) {

	urlStr := fmt.Sprintf("https://api.telapi.com/v1/Accounts/%s%s.json?", helper.Sid, uri)

	first := true
	for k, v := range param_data {

		if first {
			first = false
			urlStr = urlStr + fmt.Sprintf("%s=%s", k, v)
		} else {
			urlStr = urlStr + fmt.Sprintf("&%s=%s", k, v)
		}
	}

	fmt.Println(urlStr)
	resp, err := helper.TelapiRequest("GET", urlStr, param_data)
	if err != nil {
		return nil, err
	}

	return resp, nil

}
