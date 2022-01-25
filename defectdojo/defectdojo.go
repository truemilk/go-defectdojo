package defectdojo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	userAgent     = "go-defectdojo"
	MediaTypeJson = "application/json"
	//MediaTypeMultipart = "multipart/form-data"
)

type Client struct {
	BaseURL    *url.URL
	Token      string
	HTTPClient *http.Client
}

type errorResponse struct {
	Code        int      `json:"code,omitempty"`
	Detail      string   `json:"detail,omitempty"`
	Description []string `json:"description,omitempty"`
	Message     string   `json:"message,omitempty"`
}

func NewDojoClient(dojourl string, token string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if len(dojourl) == 0 {
		return nil, errors.New("NewDojoClient: cannot create client, URL string is empty")
	}
	if len(token) == 0 {
		return nil, errors.New("NewDojoClient: cannot create client, TOKEN string is empty")
	}

	baseurl, err := url.Parse(dojourl + "/api/v2")
	if err != nil {
		return nil, fmt.Errorf("NewDojoClient: cannot parse URL: %w", err)
	}

	return &Client{
		BaseURL:    baseurl,
		Token:      fmt.Sprintf("Token %s", token),
		HTTPClient: httpClient,
	}, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", MediaTypeJson)
	req.Header.Set("Accept", MediaTypeJson)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("sendRequest: cannot send request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		errorResp := errorResponse{
			Code: res.StatusCode,
		}
		if err = json.NewDecoder(res.Body).Decode(&errorResp); err == nil {
			return fmt.Errorf("sendRequest: API error: %v", errorResp)
		}
		return fmt.Errorf("sendRequest: unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return fmt.Errorf("sendRequest: cannot decode reponse: %w", err)
	}

	return nil
}