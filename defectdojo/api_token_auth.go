package defectdojo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiTokenAuthService struct {
	client *Client
}

type AuthToken struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Token    *string `json:"token,omitempty"`
}

func (c *ApiTokenAuthService) Create(ctx context.Context, u *AuthToken) (*AuthToken, error) {
	path := fmt.Sprintf("%s/api-token-auth/", c.client.BaseURL)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(AuthToken)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
