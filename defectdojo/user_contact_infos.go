package defectdojo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type UserContactInfo struct {
	Id                 int    `json:"id,omitempty"`
	Title              string `json:"title,omitempty"`
	PhoneNumber        string `json:"phone_number,omitempty"`
	CellNumber         string `json:"cell_number,omitempty"`
	TwitterUsername    string `json:"twitter_username,omitempty"`
	GithubUsername     string `json:"github_username,omitempty"`
	SlackUsername      string `json:"slack_username,omitempty"`
	SlackUserID        string `json:"slack_user_id,omitempty"`
	BlockExecution     bool   `json:"block_execution,omitempty"`
	ForcePasswordReset bool   `json:"force_password_reset,omitempty"`
	User               int    `json:"user,omitempty"`
}

type UserContactInfos struct {
	Count    int               `json:"count,omitempty"`
	Next     string            `json:"next,omitempty"`
	Previous string            `json:"previous,omitempty"`
	Results  []UserContactInfo `json:"results,omitempty"`
	Prefetch struct {
		User map[string]User `json:"user,omitempty"`
	} `json:"prefetch,omitempty"`
}

type UserContactInfosOptions struct {
	Limit              int
	Offset             int
	User               string
	Title              string
	PhoneNumber        string
	CellNumber         string
	TwitterUsername    string
	GithubUsername     string
	SlackUsername      string
	SlackUserID        string
	BlockExecution     string
	ForcePasswordReset string
	Prefetch           string
}

func (o *UserContactInfosOptions) ToString() string {
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
		if len(o.User) > 0 {
			opts = append(opts, fmt.Sprintf("user=%s", o.User))
		}
		if len(o.PhoneNumber) > 0 {
			opts = append(opts, fmt.Sprintf("phone_number=%s", o.PhoneNumber))
		}
		if len(o.CellNumber) > 0 {
			opts = append(opts, fmt.Sprintf("cell_number=%s", o.CellNumber))
		}
		if len(o.TwitterUsername) > 0 {
			opts = append(opts, fmt.Sprintf("twitter_username=%s", o.TwitterUsername))
		}
		if len(o.GithubUsername) > 0 {
			opts = append(opts, fmt.Sprintf("github_username=%s", o.GithubUsername))
		}
		if len(o.SlackUsername) > 0 {
			opts = append(opts, fmt.Sprintf("slack_username=%s", o.SlackUsername))
		}
		if len(o.SlackUserID) > 0 {
			opts = append(opts, fmt.Sprintf("slack_user_id=%s", o.SlackUserID))
		}
		if len(o.BlockExecution) > 0 {
			opts = append(opts, fmt.Sprintf("block_execution=%s", o.BlockExecution))
		}
		if len(o.ForcePasswordReset) > 0 {
			opts = append(opts, fmt.Sprintf("force_password_reset=%s", o.ForcePasswordReset))
		}
		if len(o.Prefetch) > 0 {
			opts = append(opts, fmt.Sprintf("prefetch=%s", o.Prefetch))
		}
		optsString += strings.Join(opts, "&")
	}
	return optsString
}

func (c *Client) UserContactInfosList(ctx context.Context, options *UserContactInfosOptions) (*UserContactInfos, error) {
	path := fmt.Sprintf("%s/user_contact_infos/%s", c.BaseURL, options.ToString())

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := UserContactInfos{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UserContactInfosRead(ctx context.Context, id int) (*UserContactInfo, error) {
	path := fmt.Sprintf("%s/user_contact_infos/%d/", c.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(UserContactInfo)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UserContactInfosCreate(ctx context.Context, u *UserContactInfo) (*UserContactInfo, error) {
	path := fmt.Sprintf("%s/user_contact_infos/", c.BaseURL)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(UserContactInfo)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UserContactInfosUpdate(ctx context.Context, id int, u *UserContactInfo) (*UserContactInfo, error) {
	path := fmt.Sprintf("%s/user_contact_infos/%d/", c.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(UserContactInfo)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UserContactInfosPartialUpdate(ctx context.Context, id int, u *UserContactInfo) (*UserContactInfo, error) {
	path := fmt.Sprintf("%s/user_contact_infos/%d/", c.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(UserContactInfo)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UserContactInfosDelete(ctx context.Context, id int) (*UserContactInfo, error) {
	path := fmt.Sprintf("%s/user_contact_infos/%d/", c.BaseURL, id)

	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(UserContactInfo)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
