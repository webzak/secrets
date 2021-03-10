package secrets

import (
	"os"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestEncryptedEnvironment(t *testing.T) {
	s, err := NewEncryptedEnvironmentStorage("masterpass", "", false)
	assert.Nil(t, err)
	secret := "baronetwothreefour five"
	s.Set("foo", secret)
	assert.NotEqual(t, "bar", os.Getenv("foo"))
	t.Logf(os.Getenv("foo"))
	readSecret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, secret, readSecret)
}

func TestEncryptedEnvironmentAbsent(t *testing.T) {
	s, err := NewEncryptedEnvironmentStorage("masterpass", "", true)
	assert.Nil(t, err)
	secret, err := s.Get("absent")
	assert.True(t, IsNotFound(err))
	assert.Equal(t, "", secret)
}
