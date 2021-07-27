package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseResponseBodyToMap parsed http response body to map
func ParseResponseBodyToMap(res *http.Response) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
