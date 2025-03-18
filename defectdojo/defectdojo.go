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
	mediaTypeJson = "application/json"
)

type Client struct {
	BaseURL    *url.URL
	Token      string
	HTTPClient *http.Client

	ApiTokenAuth     *ApiTokenAuthService
	DojoGroups       *DojoGroupsService
	Engagements      *EngagementsService
	Findings         *FindingsService
	ImportScan       *ImportScanService
	ReImportScan     *ReImportScanService
	Notes            *NotesService
	ProductTypes     *ProductTypesService
	Products         *ProductsService
	Technologies     *TechnologiesService
	TestTypes        *TestTypesService
	ToolTypes        *ToolTypesService
	UserContactInfos *UserContactInfosService
	UserProfile      *UserProfileService
	Users            *UsersService
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

	baseurl, err := url.Parse(dojourl + "/api/v2")
	if err != nil {
		return nil, fmt.Errorf("NewDojoClient: cannot parse URL: %w", err)
	}

	c := &Client{
		BaseURL:    baseurl,
		Token:      token,
		HTTPClient: httpClient,
	}

	c.ApiTokenAuth = &ApiTokenAuthService{client: c}
	c.DojoGroups = &DojoGroupsService{client: c}
	c.Engagements = &EngagementsService{client: c}
	c.Findings = &FindingsService{client: c}
	c.ImportScan = &ImportScanService{client: c}
	c.ReImportScan = &ReImportScanService{client: c}
	c.Notes = &NotesService{client: c}
	c.ProductTypes = &ProductTypesService{client: c}
	c.Products = &ProductsService{client: c}
	c.Technologies = &TechnologiesService{client: c}
	c.TestTypes = &TestTypesService{client: c}
	c.ToolTypes = &ToolTypesService{client: c}
	c.UserContactInfos = &UserContactInfosService{client: c}
	c.UserProfile = &UserProfileService{client: c}
	c.Users = &UsersService{client: c}

	return c, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", mediaTypeJson)

	if len(c.Token) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.Token))
	}

	if len(req.Header.Get("Content-Type")) == 0 {
		req.Header.Set("Content-Type", mediaTypeJson)
	}

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
