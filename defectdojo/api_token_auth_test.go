package defectdojo

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestApiTokenAuthService_Create(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
}
