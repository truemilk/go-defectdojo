package defectdojo

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUserContactInfosService_List(t *testing.T) {
	response := `{
		"count": 1,
		"next": null,
		"previous": null,
		"results": [
			{
				"id": 1,
				"title": "Security Engineer",
				"phone_number": "+1234567890",
				"cell_number": "+0987654321",
				"twitter_username": "seceng",
				"github_username": "seceng",
				"slack_username": "sec.eng",
				"slack_user_id": "U12345",
				"block_execution": false,
				"force_password_reset": false,
				"user": 1
			}
		]
	}`

	expected := UserContactInfos{
		Count:    Int(1),
		Next:     nil,
		Previous: nil,
		Results: &[]UserContactInfo{
			{
				Id:                 Int(1),
				Title:              Str("Security Engineer"),
				PhoneNumber:        Str("+1234567890"),
				CellNumber:         Str("+0987654321"),
				TwitterUsername:    Str("seceng"),
				GithubUsername:     Str("seceng"),
				SlackUsername:      Str("sec.eng"),
				SlackUserID:        Str("U12345"),
				BlockExecution:     Bool(false),
				ForcePasswordReset: Bool(false),
				User:               Int(1),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/user_contact_infos/") {
			t.Errorf("Expected /user_contact_infos/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.UserContactInfos.List(context.Background(), &UserContactInfosOptions{
		Limit:  10,
		Offset: 0,
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestUserContactInfosService_Read(t *testing.T) {
	response := `{
		"id": 123,
		"title": "Lead Engineer",
		"phone_number": "+1111111111",
		"cell_number": "+2222222222",
		"twitter_username": "lead",
		"github_username": "lead",
		"slack_username": "lead.eng",
		"slack_user_id": "U99999",
		"block_execution": true,
		"force_password_reset": false,
		"user": 5
	}`

	expected := UserContactInfo{
		Id:                 Int(123),
		Title:              Str("Lead Engineer"),
		PhoneNumber:        Str("+1111111111"),
		CellNumber:         Str("+2222222222"),
		TwitterUsername:    Str("lead"),
		GithubUsername:     Str("lead"),
		SlackUsername:      Str("lead.eng"),
		SlackUserID:        Str("U99999"),
		BlockExecution:     Bool(true),
		ForcePasswordReset: Bool(false),
		User:               Int(5),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/user_contact_infos/123/") {
			t.Errorf("Expected /user_contact_infos/123/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.UserContactInfos.Read(context.Background(), 123)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestUserContactInfosService_Create(t *testing.T) {
	response := `{
		"id": 456,
		"title": "New Engineer",
		"phone_number": "+3333333333",
		"github_username": "neweng",
		"slack_username": "new.eng",
		"block_execution": false,
		"force_password_reset": true,
		"user": 10
	}`

	expected := UserContactInfo{
		Id:                 Int(456),
		Title:              Str("New Engineer"),
		PhoneNumber:        Str("+3333333333"),
		GithubUsername:     Str("neweng"),
		SlackUsername:      Str("new.eng"),
		BlockExecution:     Bool(false),
		ForcePasswordReset: Bool(true),
		User:               Int(10),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/user_contact_infos/") {
			t.Errorf("Expected /user_contact_infos/ in path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.UserContactInfos.Create(context.Background(), &UserContactInfo{
		Title:          Str("New Engineer"),
		PhoneNumber:    Str("+3333333333"),
		GithubUsername: Str("neweng"),
		User:           Int(10),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestUserContactInfosService_Update(t *testing.T) {
	response := `{
		"id": 789,
		"title": "Updated Engineer",
		"phone_number": "+4444444444",
		"github_username": "updated",
		"block_execution": false,
		"force_password_reset": false,
		"user": 10
	}`

	expected := UserContactInfo{
		Id:                 Int(789),
		Title:              Str("Updated Engineer"),
		PhoneNumber:        Str("+4444444444"),
		GithubUsername:     Str("updated"),
		BlockExecution:     Bool(false),
		ForcePasswordReset: Bool(false),
		User:               Int(10),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/user_contact_infos/789/") {
			t.Errorf("Expected /user_contact_infos/789/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.UserContactInfos.Update(context.Background(), 789, &UserContactInfo{
		Title:          Str("Updated Engineer"),
		PhoneNumber:    Str("+4444444444"),
		GithubUsername: Str("updated"),
		User:           Int(10),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestUserContactInfosService_PartialUpdate(t *testing.T) {
	response := `{
		"id": 321,
		"title": "Partial Update",
		"phone_number": "+5555555555",
		"block_execution": false,
		"force_password_reset": false,
		"user": 10
	}`

	expected := UserContactInfo{
		Id:                 Int(321),
		Title:              Str("Partial Update"),
		PhoneNumber:        Str("+5555555555"),
		BlockExecution:     Bool(false),
		ForcePasswordReset: Bool(false),
		User:               Int(10),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/user_contact_infos/321/") {
			t.Errorf("Expected /user_contact_infos/321/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.UserContactInfos.PartialUpdate(context.Background(), 321, &UserContactInfo{
		Title: Str("Partial Update"),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestUserContactInfosService_Delete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/user_contact_infos/654/") {
			t.Errorf("Expected /user_contact_infos/654/ in path, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintln(w, "{}")
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.UserContactInfos.Delete(context.Background(), 654)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if actual == nil {
		t.Errorf("expected non-nil response")
	}
}

func TestUserContactInfosOptions_ToString(t *testing.T) {
	tests := []struct {
		name     string
		options  *UserContactInfosOptions
		expected string
	}{
		{
			name: "user only",
			options: &UserContactInfosOptions{
				User: "5",
			},
			expected: "?user=5",
		},
		{
			name: "all emitted fields",
			options: &UserContactInfosOptions{
				Limit:              10,
				Offset:             20,
				User:               "5",
				PhoneNumber:        "+1234567890",
				CellNumber:         "+0987654321",
				TwitterUsername:    "tw",
				GithubUsername:     "gh",
				SlackUsername:      "sl",
				SlackUserID:        "U1",
				BlockExecution:     "false",
				ForcePasswordReset: "true",
				Prefetch:           "user",
			},
			expected: "?limit=10&offset=20&user=5&phone_number=+1234567890&cell_number=+0987654321&twitter_username=tw&github_username=gh&slack_username=sl&slack_user_id=U1&block_execution=false&force_password_reset=true&prefetch=user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.options.ToString()
			if actual != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, actual)
			}
		})
	}
}
