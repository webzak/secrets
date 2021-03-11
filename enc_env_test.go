package secrets

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncEnv(t *testing.T) {
	s, err := NewEncEnvStorage("masterpass", "", false)
	assert.Nil(t, err)
	secret := "baronetwothreefour five"
	err = s.Set("foo", secret)
	assert.Nil(t, err)
	assert.NotEqual(t, "bar", os.Getenv("foo"))
	readSecret, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, secret, readSecret)
}

func TestEncEnvAbsent(t *testing.T) {
	s, err := NewEncEnvStorage("masterpass", "", true)
	assert.Nil(t, err)
	secret, err := s.Get("absent")
	assert.True(t, IsNotFound(err))
	assert.Equal(t, "", secret)
}

func TestEncEnvPrepare(t *testing.T) {
	s, err := NewEncEnvStorage("masterpass", "prf_", true)
	assert.Nil(t, err)
	name := "foo"
	secret := "hello"
	readName, encval, err := s.Prepare("foo", secret)
	assert.Nil(t, err)
	assert.Equal(t, "PRF_FOO", readName)
	s.es.Set(name, encval)
	readSecret, err := s.Get(name)
	assert.Nil(t, err)
	assert.Equal(t, secret, readSecret)

}
