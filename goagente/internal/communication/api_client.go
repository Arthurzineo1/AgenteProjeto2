package communication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"net/http"
)

type APIClient struct {
	BaseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
	}
}

func (c *APIClient) GenericPost(endpoint string, payload interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	return resp, nil
}
