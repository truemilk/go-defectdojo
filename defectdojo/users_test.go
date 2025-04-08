package defectdojo

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestUsersService_List(t *testing.T) {

	response := `{"count": 0,"next": "string","previous": "string","results": [{"id": 0,"username": "string",
        "first_name": "string","last_name": "string","email": "user@example.com","last_login": "2022-02-04T20:04:09.950Z",
        "is_active": true,"is_staff": true,"is_superuser": true,"password": "string"}]}`

	expected := Users{
		Count:    Int(0),
		Next:     Str("string"),
		Previous: Str("string"),
		Results: &[]User{
			{
				ID:          Int(0),
				Username:    Str("string"),
				FirstName:   Str("string"),
				LastName:    Str("string"),
				Email:       Str("user@example.com"),
				LastLogin:   Date(time.Date(2022, time.February, 4, 20, 4, 9, 950000000, time.UTC)),
				IsActive:    Bool(true),
				IsStaff:     Bool(true),
				IsSuperuser: Bool(true),
				Password:    Str("string"),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Users.List(context.Background(), nil)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestUsersService_Read(t *testing.T) {

	response := `{"id": 123,"username": "string",
		"first_name": "string","last_name": "string","email": "user@example.com",
		"last_login": "2022-02-03T14:34:15.085Z","is_active": true,"is_staff": true,
		"is_superuser": true,"password": "string"}`

	expected := User{
		ID:          Int(123),
		Username:    Str("string"),
		FirstName:   Str("string"),
		LastName:    Str("string"),
		Email:       Str("user@example.com"),
		LastLogin:   Date(time.Date(2022, time.February, 3, 14, 34, 15, 85000000, time.UTC)),
		IsActive:    Bool(true),
		IsStaff:     Bool(true),
		IsSuperuser: Bool(true),
		Password:    Str("string"),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Users.Read(context.Background(), 123)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}
package defectdojo

import (
	"context"
	"testing"
	"time"
)

func TestUsersService_List_TableDriven(t *testing.T) {
	tests := []struct {
		name          string
		responseCode  int
		responseBody  string
		expectError   bool
	}{
		{
			name:         "valid response",
			responseCode: http.StatusOK,
			responseBody: `{"count":1,"results":[{"id": 123, "username": "test"}]}`,
			expectError:  false,
		},
		{
			name:         "malformed JSON",
			responseCode: http.StatusOK,
			responseBody: `invalid json`,
			expectError:  true,
		},
		{
			name:         "HTTP error",
			responseCode: http.StatusInternalServerError,
			responseBody: `{"detail": "server error"}`,
			expectError:  true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ts := NewTestServer(tc.responseCode, tc.responseBody)
			defer ts.Close()
			client, _ := NewDojoClient(ts.URL, "dummy", nil)
			_, err := client.Users.List(context.Background(), nil)
			if tc.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestContextCancellation(t *testing.T) {
	ts := NewTestServer(http.StatusOK, `{}`)
	defer ts.Close()
	client, _ := NewDojoClient(ts.URL, "dummy", nil)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	_, err := client.Users.List(ctx, nil)
	if err == nil {
		t.Errorf("expected context deadline exceeded error")
	}
}
