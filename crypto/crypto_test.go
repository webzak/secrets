package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesGcm(t *testing.T) {
	c, err := NewAesGcmCypher("masterkey")
	assert.Nil(t, err)

	pt := "one two three four five"
	ct, err := c.Encrypt([]byte(pt))
	assert.Nil(t, err)

	rpt, err := c.Decrypt(ct)
	assert.Nil(t, err)
	assert.Equal(t, pt, string(rpt))
}
