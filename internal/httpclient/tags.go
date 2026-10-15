package httpclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type tagResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func Tags(url string, tag string) ([]string, error) {
	res, err := http.Get(url + "/v2/" + tag + "/tags/list")
	if err != nil {
		return nil, errors.New("Unable to reach registry.")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Unable to parse response.")
	}

	var parsed_body tagResponse
	err = json.Unmarshal(body, &parsed_body)
	if err != nil {
		return nil, errors.New("Error parsing response body")
	}

	return parsed_body.Tags, nil
}
