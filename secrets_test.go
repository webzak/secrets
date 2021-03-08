package secrets

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestNotInitialized(t *testing.T) {
	storage = nil
	_, err := Get("foo")
	assert.Equal(t, &StorageIsNotInitialized{}, err)
}

func TestGet(t *testing.T) {
	s := InitMemoryStorage()
	s.Set("foo", "bar")
	secret, err := Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
	secret, err = Get("absent")
	assert.Equal(t, &SecretNotFound{"absent"}, err)
	assert.Equal(t, "", secret)
}
