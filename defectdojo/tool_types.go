package defectdojo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ToolType struct {
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ToolTypes struct {
	Count    *int        `json:"count,omitempty"`
	Next     *string     `json:"next,omitempty"`
	Previous *string     `json:"previous,omitempty"`
	Results  []*ToolType `json:"results,omitempty"`
}

type ToolTypesOptions struct {
	Limit       int
	Offset      int
	ID          int
	Name        string
	Description string
}

func (o *ToolTypesOptions) ToString() string {
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
		if o.ID > 0 {
			opts = append(opts, fmt.Sprintf("id=%d", o.ID))
		}
		if len(o.Name) > 0 {
			opts = append(opts, fmt.Sprintf("username=%s", o.Name))
		}
		if len(o.Description) > 0 {
			opts = append(opts, fmt.Sprintf("first_name=%s", o.Description))
		}
		optsString += strings.Join(opts, "&")
	}
	return optsString
}

func (c *Client) ToolTypesList(ctx context.Context, options *ToolTypesOptions) (*ToolTypes, error) {
	path := fmt.Sprintf("%s/tool_types/%s", c.BaseURL, options.ToString())

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := ToolTypes{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ToolTypesRead(ctx context.Context, id int) (*ToolType, error) {
	path := fmt.Sprintf("%s/tool_types/%d/", c.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ToolType)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ToolTypesCreate(ctx context.Context, u *ToolType) (*ToolType, error) {
	path := fmt.Sprintf("%s/tool_types/", c.BaseURL)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ToolType)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ToolTypesUpdate(ctx context.Context, id int, u *ToolType) (*ToolType, error) {
	path := fmt.Sprintf("%s/tool_types/%d/", c.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ToolType)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ToolTypesPartialUpdate(ctx context.Context, id int, u *ToolType) (*ToolType, error) {
	path := fmt.Sprintf("%s/tool_types/%d/", c.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ToolType)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ToolTypesDelete(ctx context.Context, id int) (*ToolType, error) {
	path := fmt.Sprintf("%s/tool_types/%d/", c.BaseURL, id)

	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ToolType)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
