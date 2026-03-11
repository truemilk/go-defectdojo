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
		_, _ = fmt.Fprintln(w, response)
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
		_, _ = fmt.Fprintln(w, response)
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

func TestUsersService_Create(t *testing.T) {
	response := `{
		"id": 456,
		"username": "newuser",
		"first_name": "New",
		"last_name": "User",
		"email": "new@example.com",
		"last_login": "2022-03-01T10:00:00Z",
		"is_active": true,
		"is_staff": false,
		"is_superuser": false,
		"password": "hashed"
	}`

	expected := User{
		ID:          Int(456),
		Username:    Str("newuser"),
		FirstName:   Str("New"),
		LastName:    Str("User"),
		Email:       Str("new@example.com"),
		LastLogin:   Date(time.Date(2022, time.March, 1, 10, 0, 0, 0, time.UTC)),
		IsActive:    Bool(true),
		IsStaff:     Bool(false),
		IsSuperuser: Bool(false),
		Password:    Str("hashed"),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/users/") {
			t.Errorf("Expected /users/ in path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Users.Create(context.Background(), &User{
		Username:  Str("newuser"),
		FirstName: Str("New"),
		LastName:  Str("User"),
		Email:     Str("new@example.com"),
		Password:  Str("password123"),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestUsersService_Update(t *testing.T) {
	response := `{
		"id": 789,
		"username": "updateduser",
		"first_name": "Updated",
		"last_name": "User",
		"email": "updated@example.com",
		"last_login": "2022-04-01T10:00:00Z",
		"is_active": true,
		"is_staff": true,
		"is_superuser": false,
		"password": "hashed"
	}`

	expected := User{
		ID:          Int(789),
		Username:    Str("updateduser"),
		FirstName:   Str("Updated"),
		LastName:    Str("User"),
		Email:       Str("updated@example.com"),
		LastLogin:   Date(time.Date(2022, time.April, 1, 10, 0, 0, 0, time.UTC)),
		IsActive:    Bool(true),
		IsStaff:     Bool(true),
		IsSuperuser: Bool(false),
		Password:    Str("hashed"),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/users/789/") {
			t.Errorf("Expected /users/789/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Users.Update(context.Background(), 789, &User{
		Username:  Str("updateduser"),
		FirstName: Str("Updated"),
		LastName:  Str("User"),
		Email:     Str("updated@example.com"),
		IsStaff:   Bool(true),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestUsersService_PartialUpdate(t *testing.T) {
	response := `{
		"id": 321,
		"username": "partialuser",
		"first_name": "Partial",
		"last_name": "User",
		"email": "partial@example.com",
		"last_login": "2022-05-01T10:00:00Z",
		"is_active": false,
		"is_staff": false,
		"is_superuser": false,
		"password": "hashed"
	}`

	expected := User{
		ID:          Int(321),
		Username:    Str("partialuser"),
		FirstName:   Str("Partial"),
		LastName:    Str("User"),
		Email:       Str("partial@example.com"),
		LastLogin:   Date(time.Date(2022, time.May, 1, 10, 0, 0, 0, time.UTC)),
		IsActive:    Bool(false),
		IsStaff:     Bool(false),
		IsSuperuser: Bool(false),
		Password:    Str("hashed"),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/users/321/") {
			t.Errorf("Expected /users/321/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Users.PartialUpdate(context.Background(), 321, &User{
		IsActive: Bool(false),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestUsersService_Delete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/users/654/") {
			t.Errorf("Expected /users/654/ in path, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintln(w, "{}")
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Users.Delete(context.Background(), 654)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if actual == nil {
		t.Errorf("expected non-nil response")
	}
}

func TestUsersOptions_ToString(t *testing.T) {
	tests := []struct {
		name     string
		options  *UsersOptions
		expected string
	}{
		{
			name: "username only",
			options: &UsersOptions{
				Username: "admin",
			},
			expected: "?username=admin",
		},
		{
			name: "all fields",
			options: &UsersOptions{
				Limit:     10,
				Offset:    20,
				ID:        5,
				Username:  "admin",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@example.com",
			},
			expected: "?limit=10&offset=20&id=5&username=admin&first_name=John&last_name=Doe&email=john@example.com",
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
