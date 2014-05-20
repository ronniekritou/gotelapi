package telapi

import (
	"net/url"
)

func DataMapToUrlValues(data map[string]string) *url.Values {
	data_values := &url.Values{}

	for key, value := range data {
		data_values.Add(key, value)
	}

	return data_values
}
