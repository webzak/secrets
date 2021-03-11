package secrets

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvWithoutNameTransform(t *testing.T) {
	s := NewEnvStorage("", false)
	s.Set("foo", "bar")
	assert.Equal(t, "bar", os.Getenv("foo"))
	secret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
}

func TestEnvWithPrefix(t *testing.T) {
	s := NewEnvStorage("my_", false)
	s.Set("foo", "bar")
	assert.Equal(t, "bar", os.Getenv("my_foo"))
	secret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
}

func TestEnvWithUppercase(t *testing.T) {
	s := NewEnvStorage("", true)
	s.Set("foo", "bar")
	assert.Equal(t, "bar", os.Getenv("FOO"))
	secret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
}

func TestEnvWithPrefixAngUppercase(t *testing.T) {
	s := NewEnvStorage("my_", true)
	s.Set("foo", "bar")
	assert.Equal(t, "bar", os.Getenv("MY_FOO"))
	secret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", secret)
}

func TestEnvAbsent(t *testing.T) {
	s := NewEnvStorage("", true)
	secret, err := s.Get("absent")
	assert.True(t, IsNotFound(err))
	assert.Equal(t, "", secret)
}
