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

	response := `{"token":"token"}`

	expected := AuthToken{Token: String("token")}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "", nil)

	actual, err := dj.ApiTokenAuth.Create(context.Background(), &AuthToken{
		Username: String("username"),
		Password: String("password"),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}
