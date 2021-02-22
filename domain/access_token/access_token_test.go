package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.NotNil(t, at, "Brand new access token should not be nil")

	assert.False(t, at.IsExpired(), "Brand new access token should not be expired")

	assert.NotEqualValues(t, at.AccessToken, "", "New access token should not have defined access token id")

	assert.EqualValues(t, at.UserId, 0, "New access token should not have an associated user id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.IsExpired(), "Empty access token should not be expired")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Empty access token should not be expired")

	at.Expires = time.Now().UTC().Add(-24 * time.Hour + time.Second).Unix()
	assert.True(t, at.IsExpired(), "access token expiring 24 hours 1 second from now should be expired")
}