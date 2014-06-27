package telapi

import (
	"errors"
)

func (helper TelapiHelper) SendSMS(to string, from string, body string) error {

	if to == "" || from == "" {
		return errors.New("Missing required To or From.")
	}

	data := map[string]string{
		"To":            to,
		"From":          from,
		"Body":          body,
		"AllowMultiple": "true",
	}

	if _, err := helper.PostRequest("/SMS/Messages", data); err != nil {
		return err
	}

	return nil
}
