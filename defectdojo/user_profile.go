package defectdojo

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type UserProfile struct {
	User *struct {
		Id          *int       `json:"id,omitempty"`
		Username    *string    `json:"username,omitempty"`
		FirstName   *string    `json:"first_name,omitempty"`
		LastName    *string    `json:"last_name,omitempty"`
		Email       *string    `json:"email,omitempty"`
		LastLogin   *time.Time `json:"last_login,omitempty"`
		IsActive    *bool      `json:"is_active,omitempty"`
		IsStaff     *bool      `json:"is_staff,omitempty"`
		IsSuperuser *bool      `json:"is_superuser,omitempty"`
		Password    *string    `json:"password,omitempty"`
	} `json:"user,omitempty"`
	UserContactInfo *struct {
		Id                 *int    `json:"id,omitempty"`
		Title              *string `json:"title,omitempty"`
		PhoneNumber        *string `json:"phone_number,omitempty"`
		CellNumber         *string `json:"cell_number,omitempty"`
		TwitterUsername    *string `json:"twitter_username,omitempty"`
		GithubUsername     *string `json:"github_username,omitempty"`
		SlackUsername      *string `json:"slack_username,omitempty"`
		SlackUserId        *string `json:"slack_user_id,omitempty"`
		BlockExecution     *bool   `json:"block_execution,omitempty"`
		ForcePasswordReset *bool   `json:"force_password_reset,omitempty"`
		User               *int    `json:"user,omitempty"`
	} `json:"user_contact_info,omitempty"`
	GlobalRole *struct {
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

func (c *Client) UserProfileList(ctx context.Context) (*UserProfile, error) {
	path := fmt.Sprintf("%s/user_profile/", c.BaseURL)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(UserProfile)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
