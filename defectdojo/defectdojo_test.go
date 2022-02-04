package defectdojo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewDojoClient(t *testing.T) {

	t.Run("Empty URL", func(t *testing.T) {
		_, err := NewDojoClient("", "", nil)

		if cmp.Equal(err, nil) {
			t.Errorf("supposed to get an error")
		}

	})
}
