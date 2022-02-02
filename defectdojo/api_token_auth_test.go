package defectdojo

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestApiTokenAuthService_Create(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	object := AuthToken{
		Username: nil,
		Password: nil,
		Token:    String("token"),
	}
	response, _ := httpmock.NewJsonResponder(200, object)
	httpmock.RegisterResponder("POST", "server/api/v2/api-token-auth/", response)

	dj, err := NewDojoClient("server", "", nil)
	assert.Nil(t, err)
	ctx := context.Background()

	r, err := dj.ApiTokenAuth.Create(ctx, &AuthToken{
		Username: String("username"),
		Password: String("password"),
	})
	assert.Nil(t, err)
	b, err := json.Marshal(r)
	assert.Nil(t, err)

	actual := string(b)
	assert.Equal(t, `{"token":"token"}`, actual)
}
