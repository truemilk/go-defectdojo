package defectdojo

import (
	"context"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestUsersService_Read(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "server/api/v2/users/123/",
		httpmock.NewStringResponder(200, `{"id": 123,"username": "string",
		"first_name": "string","last_name": "string","email": "user@example.com",
		"last_login": "2022-02-03T14:34:15.085Z","is_active": true,"is_staff": true,
		"is_superuser": true,"password": "string"}`))

	dj, err := NewDojoClient("server", "token", nil)
	assert.Nil(t, err)

	ctx := context.Background()
	r, err := dj.Users.Read(ctx, 123)
	assert.Nil(t, err)

	assert.Equal(t, &User{
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
	}, r)
}
