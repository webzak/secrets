package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotInitialized(t *testing.T) {
	storage = nil
	_, err := Get("foo")
	assert.True(t, IsNotInitialized(err))
}

func TestGet(t *testing.T) {
	s := InitMemoryStorage()
	s.Set("foo", "bar")
	secret, err := Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
	secret, err = Get("absent")
	assert.True(t, IsNotFound(err))
	assert.Equal(t, "", secret)
}
