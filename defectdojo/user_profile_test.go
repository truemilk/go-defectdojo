package defectdojo

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestUserProfileService_List(t *testing.T) {
	response := `{
		"user": {
			"id": 1,
			"username": "admin",
			"first_name": "Admin",
			"last_name": "User",
			"email": "admin@example.com",
			"last_login": "2022-01-01T12:00:00Z",
			"is_active": true,
			"is_staff": true,
			"is_superuser": true
		},
		"user_contact_info": {
			"id": 1,
			"title": "Administrator",
			"phone_number": "+1234567890",
			"block_execution": false,
			"force_password_reset": false,
			"user": 1
		},
		"global_role": {
			"id": 1,
			"user": 1,
			"role": 4
		},
		"dojo_group_member": [
			{
				"id": 1,
				"group": 1,
				"user": 1,
				"role": 4
			}
		],
		"product_type_member": [],
		"product_member": []
	}`

	expected := UserProfile{
		User: &User{
			ID:          Int(1),
			Username:    Str("admin"),
			FirstName:   Str("Admin"),
			LastName:    Str("User"),
			Email:       Str("admin@example.com"),
			LastLogin:   Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
			IsActive:    Bool(true),
			IsStaff:     Bool(true),
			IsSuperuser: Bool(true),
		},
		UserContactInfo: &UserContactInfo{
			Id:                 Int(1),
			Title:              Str("Administrator"),
			PhoneNumber:        Str("+1234567890"),
			BlockExecution:     Bool(false),
			ForcePasswordReset: Bool(false),
			User:               Int(1),
		},
		GlobalRole: &struct {
			Id    *int `json:"id,omitempty"`
			User  *int `json:"user,omitempty"`
			Group *int `json:"group,omitempty"`
			Role  *int `json:"role,omitempty"`
		}{
			Id:   Int(1),
			User: Int(1),
			Role: Int(4),
		},
		DojoGroupMember: &[]struct {
			Id    *int `json:"id,omitempty"`
			Group *int `json:"group,omitempty"`
			User  *int `json:"user,omitempty"`
			Role  *int `json:"role,omitempty"`
		}{
			{
				Id:    Int(1),
				Group: Int(1),
				User:  Int(1),
				Role:  Int(4),
			},
		},
		ProductTypeMember: &[]struct {
			Id          *int `json:"id,omitempty"`
			ProductType *int `json:"product_type,omitempty"`
			User        *int `json:"user,omitempty"`
			Role        *int `json:"role,omitempty"`
		}{},
		ProductMember: &[]struct {
			Id      *int `json:"id,omitempty"`
			Product *int `json:"product,omitempty"`
			User    *int `json:"user,omitempty"`
			Role    *int `json:"role,omitempty"`
		}{},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/user_profile/") {
			t.Errorf("Expected /user_profile/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.UserProfile.List(context.Background())
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}
