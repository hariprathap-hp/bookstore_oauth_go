package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours always")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "New Access Token should not be expired")
	assert.True(t, at.UserID == 0, "New Access Token should not have an used id associated")
	assert.EqualValues(t, "", at.AccessToken, "New Access Token should not have an Access Token Defined")
}

func TestAccessTokenExpirationTime(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "expiration time for an empty access token is expired by default")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring 3 hours from now should not be expired already")
}
