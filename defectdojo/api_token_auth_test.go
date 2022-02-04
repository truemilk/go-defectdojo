package defectdojo

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestApiTokenAuthService_Create(t *testing.T) {

	t.Run("get token", func(t *testing.T) {
		response := `{"token":"token"}`

		expected := AuthToken{Token: String("token")}

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, response)
		}))
		defer ts.Close()

		dj, err := NewDojoClient(ts.URL, "", nil)
		if !cmp.Equal(err, nil) {
			t.Errorf("error")
		}

		actual, err := dj.ApiTokenAuth.Create(context.Background(), &AuthToken{
			Username: String("username"),
			Password: String("password"),
		})
		if !cmp.Equal(err, nil) {
			t.Errorf("error")
		}

		if !cmp.Equal(actual, &expected) {
			t.Errorf("should be the same")
		}
	})

	t.Run("malformed json", func(t *testing.T) {
		response := `{"token":"token"`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, response)
		}))
		defer ts.Close()

		dj, err := NewDojoClient(ts.URL, "", nil)
		if !cmp.Equal(err, nil) {
			t.Errorf("error")
		}

		_, err = dj.ApiTokenAuth.Create(context.Background(), &AuthToken{
			Username: String("username"),
			Password: String("password"),
		})
		if cmp.Equal(err, nil) {
			t.Errorf("supposed to get an error")
		}
	})
}
