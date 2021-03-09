package secrets

import (
	"os"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestEnvironmentWithoutNameTransform(t *testing.T) {
	s := InitEnvironmentStorage("", false)
	s.Set("foo", "bar")
	assert.Equal(t, "bar", os.Getenv("foo"))
	secret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
}

func TestEnvironmentWithPrefix(t *testing.T) {
	s := InitEnvironmentStorage("my_", false)
	s.Set("foo", "bar")
	assert.Equal(t, "bar", os.Getenv("my_foo"))
	secret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
}

func TestEnvironmentWithUppercase(t *testing.T) {
	s := InitEnvironmentStorage("", true)
	s.Set("foo", "bar")
	assert.Equal(t, "bar", os.Getenv("FOO"))
	secret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
}

func TestEnvironmentWithPrefixAngUppercase(t *testing.T) {
	s := InitEnvironmentStorage("my_", true)
	s.Set("foo", "bar")
	assert.Equal(t, "bar", os.Getenv("MY_FOO"))
	secret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
}

func TestEnvironmentAbsent(t *testing.T) {
	s := InitEnvironmentStorage("", true)
	secret, err := s.Get("absent")
	assert.True(t, IsNotFound(err))
	assert.Equal(t, "", secret)
}
