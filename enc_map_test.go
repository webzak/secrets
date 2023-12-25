package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncMapStorage(t *testing.T) {
	s, err := NewEncMapStorage("masterpass", "")
	assert.Nil(t, err)

	v, err := s.Get("foo")
	assert.NotNil(t, err)
	assert.Equal(t, "", v)

	d, err := s.Dump()
	assert.Nil(t, err)
	assert.Equal(t, "", d)

	s.Set("key1", "val1")
	s.Set("key2", "val2")

	d, err = s.Dump()
	assert.Nil(t, err)
	//t.Log(d)

	s2, err := NewEncMapStorage("masterpass", d)
	assert.Nil(t, err)

	v, err = s2.Get("key1")
	assert.Nil(t, err)
	assert.Equal(t, "val1", v)

	v, err = s2.Get("key2")
	assert.Nil(t, err)
	assert.Equal(t, "val2", v)
}
