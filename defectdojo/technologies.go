package defectdojo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type TechnologiesService struct {
	client *Client
}

type Technology struct {
	Id           *int       `json:"id,omitempty"`
	Tags         *[]string  `json:"tags,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Confidence   *int       `json:"confidence,omitempty"`
	Version      *string    `json:"version,omitempty"`
	Icon         *string    `json:"icon,omitempty"`
	Website      *string    `json:"website,omitempty"`
	WebsiteFound *string    `json:"website_found,omitempty"`
	Created      *time.Time `json:"created,omitempty"`
	Product      *int       `json:"product,omitempty"`
	User         *int       `json:"user,omitempty"`
}

type Technologies struct {
	Count    *int          `json:"count,omitempty"`
	Next     *string       `json:"next,omitempty"`
	Previous *string       `json:"previous,omitempty"`
	Results  *[]Technology `json:"results,omitempty"`
}

type TechnologiesOptions struct {
	Limit  int
	Offset int
	ID     int
}

func (o *TechnologiesOptions) ToString() string {
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
		optsString += strings.Join(opts, "&")
	}
	return optsString
}

func (c *TechnologiesService) List(ctx context.Context, options *TechnologiesOptions) (*Technologies, error) {
	path := fmt.Sprintf("%s/technologies/%s", c.client.BaseURL, options.ToString())

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Technologies{}
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *TechnologiesService) Read(ctx context.Context, id int) (*Technology, error) {
	path := fmt.Sprintf("%s/technologies/%d/", c.client.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Technology)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *TechnologiesService) Create(ctx context.Context, u *Technology) (*Technology, error) {
	path := fmt.Sprintf("%s/technologies/", c.client.BaseURL)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Technology)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *TechnologiesService) Delete(ctx context.Context, id int) (*Technology, error) {
	path := fmt.Sprintf("%s/technologies/%d/", c.client.BaseURL, id)

	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Technology)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
