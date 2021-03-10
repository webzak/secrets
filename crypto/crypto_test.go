package crypto

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestAesGcm(t *testing.T) {
	c, err := NewAesGcmCypher("masterkey")
	assert.Nil(t, err)

	pt := "one two three four five"
	ct, err := c.Encrypt([]byte(pt))
	assert.Nil(t, err)
	t.Logf("%x", ct)

	rpt, err := c.Decrypt(ct)
	assert.Nil(t, err)
	t.Logf("%x", rpt)
}
