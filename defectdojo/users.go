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

type User struct {
	ID          *int       `json:"id,omitempty"`
	Username    *string    `json:"username,omitempty"`
	FirstName   *string    `json:"first_name,omitempty"`
	LastName    *string    `json:"last_name,omitempty"`
	Email       *string    `json:"email,omitempty"`
	LastLogin   *time.Time `json:"last_login,omitempty"`
	IsActive    *bool      `json:"is_active,omitempty"`
	IsStaff     *bool      `json:"is_staff,omitempty"`
	IsSuperuser *bool      `json:"is_superuser,omitempty"`
	Password    *string    `json:"password,omitempty"`
}

type Users struct {
	Count    *int    `json:"count,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []*User `json:"results,omitempty"`
}

type UsersOptions struct {
	Limit     int
	Offset    int
	ID        int
	Username  string
	FirstName string
	LastName  string
	Email     string
}

func (o *UsersOptions) ToString() string {
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
		if len(o.Username) > 0 {
			opts = append(opts, fmt.Sprintf("username=%s", o.Username))
		}
		if len(o.FirstName) > 0 {
			opts = append(opts, fmt.Sprintf("first_name=%s", o.FirstName))
		}
		if len(o.LastName) > 0 {
			opts = append(opts, fmt.Sprintf("last_name=%s", o.LastName))
		}
		if len(o.Email) > 0 {
			opts = append(opts, fmt.Sprintf("email=%s", o.Email))
		}
		optsString += strings.Join(opts, "&")
	}
	return optsString
}

func (c *Client) UsersList(ctx context.Context, options *UsersOptions) (*Users, error) {
	path := fmt.Sprintf("%s/users/%s", c.BaseURL, options.ToString())

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Users{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UsersRead(ctx context.Context, id int) (*User, error) {
	path := fmt.Sprintf("%s/users/%d/", c.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(User)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UsersCreate(ctx context.Context, u *User) (*User, error) {
	path := fmt.Sprintf("%s/users/", c.BaseURL)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(User)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UsersUpdate(ctx context.Context, id int, u *User) (*User, error) {
	path := fmt.Sprintf("%s/users/%d/", c.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(User)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UsersPartialUpdate(ctx context.Context, id int, u *User) (*User, error) {
	path := fmt.Sprintf("%s/users/%d/", c.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(User)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UsersDelete(ctx context.Context, id int) (*User, error) {
	path := fmt.Sprintf("%s/users/%d/", c.BaseURL, id)

	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(User)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
