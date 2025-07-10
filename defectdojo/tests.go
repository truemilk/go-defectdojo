package defectdojo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type TestsService struct {
	client *Client
}

type Test struct {
	Id                   *int      `json:"id,omitempty"`
	EngagementId         *int      `json:"engagement,omitempty"`
	Notes                *[]Note   `json:"notes,omitempty"`
	Tags                 *[]string `json:"tags,omitempty"`
	ScanType             *string   `json:"scan_type,omitempty"`
	Title                *string   `json:"title,omitempty"`
	Description          *string   `json:"description,omitempty"`
	TargetStart          *string   `json:"target_start,omitempty"`
	TargetEnd            *string   `json:"target_end,omitempty"`
	PercentComplete      *int      `json:"percent_complete,omitempty"`
	BuildId              *string   `json:"build_id,omitempty"`
	BranchTag            *string   `json:"branch_tag,omitempty"`
	CommitHash           *string   `json:"commit_hash,omitempty"`
	Lead                 *int      `json:"lead,omitempty"`
	TestType             *int      `json:"test_type,omitempty"`
	ApiScanConfiguration *int      `json:"api_scan_configuration,omitempty"`
}

type Tests struct {
	Count    *int    `json:"count,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  *[]Test `json:"results,omitempty"`
}

type TestsOptions struct {
	Limit    int
	Offset   int
	Title    string
	Engagement int
	TestType int
}

func (o *TestsOptions) ToString() string {
	var opts []string
	var optsString string
	if o != nil {
		optsString += "?"
		if o.Limit > 0 {
			opts = append(opts, fmt.Sprintf("limit=%d", o.Limit))
		}
		if o.Offset > 0 {
			opts = append(opts, fmt.Sprintf("offset=%d", o.Offset))
		}
		if len(o.Title) > 0 {
			opts = append(opts, fmt.Sprintf("title=%s", o.Title))
		}
		if o.TestType > 0 {
			opts = append(opts, fmt.Sprintf("test_type=%d", o.TestType))
		}
		if o.Engagement > 0 {
			opts = append(opts, fmt.Sprintf("engagement=%d", o.Engagement))
		}
		optsString += strings.Join(opts, "&")
	}
	return optsString
}

func (c *TestsService) List(ctx context.Context, options *TestsOptions) (*Tests, error) {
	path := fmt.Sprintf("%s/tests/%s", c.client.BaseURL, options.ToString())

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Tests{}
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *TestsService) Read(ctx context.Context, id int) (*Test, error) {
	path := fmt.Sprintf("%s/tests/%d/", c.client.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Test)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *TestsService) Create(ctx context.Context, u *Test) (*Test, error) {
	path := fmt.Sprintf("%s/tests/", c.client.BaseURL)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Test)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *TestsService) Update(ctx context.Context, id int, u *Test) (*Test, error) {
	path := fmt.Sprintf("%s/tests/%d/", c.client.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Test)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *TestsService) PartialUpdate(ctx context.Context, id int, u *Test) (*Test, error) {
	path := fmt.Sprintf("%s/tests/%d/", c.client.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Test)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
