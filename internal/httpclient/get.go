package httpclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Get(url string) ([]string, error) {
	res, err := http.Get(url + "/v2/_catalog")
	if err != nil {
		return nil, errors.New("Unable to reach registry.")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Unable to parse response.")
	}

	var parsed_body map[string][]string
	err = json.Unmarshal(body, &parsed_body)
	if err != nil {
		return nil, errors.New("Error parsing response body")
	}

	return parsed_body["repositories"], nil
}
