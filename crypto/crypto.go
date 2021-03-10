package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

// Cipher provides the interface to encrypt and decrypt the secrets
type Cipher interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

// AesGcmCypher implements the AES GCM
type AesGcmCypher struct {
	adead cipher.AEAD
}

// NewAesGcmCypher creates the intance of AesGcmCypher
func NewAesGcmCypher(key string) (*AesGcmCypher, error) {
	bkey := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(bkey[:])
	if err != nil {
		return nil, &CryptoError{err.Error()}
	}
	adead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, &CryptoError{err.Error()}
	}
	return &AesGcmCypher{adead}, nil
}

// Encrypt the data
func (c *AesGcmCypher) Encrypt(pt []byte) ([]byte, error) {
	nonce := make([]byte, c.adead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, &CryptoError{err.Error()}
	}
	return c.adead.Seal(nonce, nonce, pt, nil), nil
}

// Decrypt the data
func (c *AesGcmCypher) Decrypt(ct []byte) ([]byte, error) {
	n := c.adead.NonceSize()
	if len(ct) <= n {
		return nil, &CryptoError{"ct is too short"}
	}
	pt, err := c.adead.Open(nil, ct[:n], ct[n:], nil)
	if err != nil {
		return nil, &CryptoError{err.Error()}
	}
	return pt, nil

}

// ByteToB64 encodes to bytes to base64 string
func ByteToB64(key []byte) string {
	return base64.RawStdEncoding.EncodeToString(key)
}

// B64ToBytes converst base64 string to bytes
func B64ToByte(s string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(s)
}
