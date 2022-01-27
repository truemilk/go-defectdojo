package defectdojo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type TestTypesService struct {
	client *Client
}

type TestType struct {
	Id          *int      `json:"id,omitempty"`
	Tags        []*string `json:"tags,omitempty"`
	Name        *string   `json:"name,omitempty"`
	StaticTool  *bool     `json:"static_tool,omitempty"`
	DynamicTool *bool     `json:"dynamic_tool,omitempty"`
	Active      *bool     `json:"active,omitempty"`
}

type TestTypes struct {
	Count    *int        `json:"count,omitempty"`
	Next     *string     `json:"next,omitempty"`
	Previous *string     `json:"previous,omitempty"`
	Results  []*TestType `json:"results,omitempty"`
}

type TestTypesOptions struct {
	Limit  int
	Offset int
	Name   string
}

func (o *TestTypesOptions) ToString() string {
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
		if len(o.Name) > 0 {
			opts = append(opts, fmt.Sprintf("name=%s", o.Name))
		}
		optsString += strings.Join(opts, "&")
	}
	return optsString
}

func (c *TestTypesService) List(ctx context.Context, options *TestTypesOptions) (*TestTypes, error) {
	path := fmt.Sprintf("%s/test_types/%s", c.client.BaseURL, options.ToString())

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := TestTypes{}
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *TestTypesService) Read(ctx context.Context, id int) (*TestType, error) {
	path := fmt.Sprintf("%s/test_types/%d/", c.client.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(TestType)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *TestTypesService) Create(ctx context.Context, u *TestType) (*TestType, error) {
	path := fmt.Sprintf("%s/test_types/", c.client.BaseURL)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(TestType)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *TestTypesService) Update(ctx context.Context, id int, u *TestType) (*TestType, error) {
	path := fmt.Sprintf("%s/test_types/%d/", c.client.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(TestType)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *TestTypesService) PartialUpdate(ctx context.Context, id int, u *TestType) (*TestType, error) {
	path := fmt.Sprintf("%s/test_types/%d/", c.client.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(TestType)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
