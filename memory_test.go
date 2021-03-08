package secrets

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestMemory(t *testing.T) {
	s := InitMemoryStorage()
	s.Set("foo", "bar")
	secret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
}

func TestMemoryAbsent(t *testing.T) {
	s := InitMemoryStorage()
	secret, err := s.Get("absent")
	assert.Equal(t, &SecretNotFound{"absent"}, err)
	assert.Equal(t, "", secret)
}
