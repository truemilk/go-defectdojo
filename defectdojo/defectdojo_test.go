package defectdojo

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewDojoClient(t *testing.T) {

	t.Run("empty URL", func(t *testing.T) {

		_, err := NewDojoClient("", "", nil)
		if cmp.Equal(err, nil) {
			t.Errorf("expected an error with an empty URL")
		}
	})
}

func TestClient_sendRequest(t *testing.T) {

	t.Run("server unavailable", func(t *testing.T) {

		c, _ := NewDojoClient("unavailable", "token", nil)

		req, err := http.NewRequest(http.MethodGet, "", nil)
		if !cmp.Equal(err, nil) {
			t.Errorf("ERR")
		}

		var res errorResponse
		err = c.sendRequest(req, &res)
		if cmp.Equal(err, nil) {
			t.Errorf("expected an error with server unavailable")
		}
	})

}
