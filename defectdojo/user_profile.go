package defectdojo

import (
	"context"
	"fmt"
	"net/http"
)

type UserProfileService struct {
	client *Client
}

type UserProfile struct {
	User            *User            `json:"user,omitempty"`
	UserContactInfo *UserContactInfo `json:"user_contact_info,omitempty"`
	GlobalRole      *struct {
		Id    *int `json:"id,omitempty"`
		User  *int `json:"user,omitempty"`
		Group *int `json:"group,omitempty"`
		Role  *int `json:"role,omitempty"`
	} `json:"global_role,omitempty"`
	DojoGroupMember []*struct {
		Id    *int `json:"id,omitempty"`
		Group *int `json:"group,omitempty"`
		User  *int `json:"user,omitempty"`
		Role  *int `json:"role,omitempty"`
	} `json:"dojo_group_member,omitempty"`
	ProductTypeMember []*struct {
		Id          *int `json:"id,omitempty"`
		ProductType *int `json:"product_type,omitempty"`
		User        *int `json:"user,omitempty"`
		Role        *int `json:"role,omitempty"`
	} `json:"product_type_member,omitempty"`
	ProductMember []*struct {
		Id      *int `json:"id,omitempty"`
		Product *int `json:"product,omitempty"`
		User    *int `json:"user,omitempty"`
		Role    *int `json:"role,omitempty"`
	} `json:"product_member,omitempty"`
}

func (c *UserProfileService) List(ctx context.Context) (*UserProfile, error) {
	path := fmt.Sprintf("%s/user_profile/", c.client.BaseURL)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(UserProfile)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
