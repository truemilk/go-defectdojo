package defectdojo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDojoClient(t *testing.T) {

	t.Run("Empty URL", func(t *testing.T) {
		_, err := NewDojoClient("", "", nil)

		if assert.Error(t, err) {
			assert.Equal(t, err, fmt.Errorf("NewDojoClient: cannot create client, URL string is empty"))
		}
	})
}
