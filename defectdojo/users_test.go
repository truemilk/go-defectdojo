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

	t.Run("normal execution", func(t *testing.T) {
		response := `{"count": 0,"next": "string","previous": "string","results": [{"id": 0,"username": "string",
        "first_name": "string","last_name": "string","email": "user@example.com","last_login": "2022-02-04T20:04:09.950Z",
        "is_active": true,"is_staff": true,"is_superuser": true,"password": "string"}]}`

		expected := Users{
			Count:    Int(0),
			Next:     String("string"),
			Previous: String("string"),
			Results: []*User{
				{
					ID:          Int(0),
					Username:    String("string"),
					FirstName:   String("string"),
					LastName:    String("string"),
					Email:       String("user@example.com"),
					LastLogin:   Date(time.Date(2022, time.February, 4, 20, 4, 9, 950000000, time.UTC)),
					IsActive:    Bool(true),
					IsStaff:     Bool(true),
					IsSuperuser: Bool(true),
					Password:    String("string"),
				},
			},
		}

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, response)
		}))
		defer ts.Close()

		dj, err := NewDojoClient(ts.URL, "token", nil)
		if !cmp.Equal(err, nil) {
			t.Errorf("error")
		}

		actual, err := dj.Users.List(context.Background(), nil)
		if !cmp.Equal(err, nil) {
			t.Errorf("error")
		}

		if !cmp.Equal(actual, &expected) {
			t.Errorf("should have been equal, %+v, %+v", actual, &expected)
		}
	})
}

func TestUsersService_Read(t *testing.T) {

	t.Run("normal execution", func(t *testing.T) {
		response := `{"id": 123,"username": "string",
		"first_name": "string","last_name": "string","email": "user@example.com",
		"last_login": "2022-02-03T14:34:15.085Z","is_active": true,"is_staff": true,
		"is_superuser": true,"password": "string"}`

		expected := User{
			ID:          Int(123),
			Username:    String("string"),
			FirstName:   String("string"),
			LastName:    String("string"),
			Email:       String("user@example.com"),
			LastLogin:   Date(time.Date(2022, time.February, 3, 14, 34, 15, 85000000, time.UTC)),
			IsActive:    Bool(true),
			IsStaff:     Bool(true),
			IsSuperuser: Bool(true),
			Password:    String("string"),
		}

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, response)
		}))
		defer ts.Close()

		dj, err := NewDojoClient(ts.URL, "token", nil)
		if !cmp.Equal(err, nil) {
			t.Errorf("error")
		}

		actual, err := dj.Users.Read(context.Background(), 123)
		if !cmp.Equal(err, nil) {
			t.Errorf("error")
		}

		if !cmp.Equal(actual, &expected) {
			t.Errorf("should have been equal, %+v, %+v", actual, &expected)
		}
	})
}
