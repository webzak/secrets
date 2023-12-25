package secrets

import (
	"encoding/json"
	"fmt"

	"github.com/webzak/secrets/crypto"
)

// EncMapStorage stores secrets in encrypted JSON with a plan string map structure
type EncMapStorage struct {
	cipher crypto.Cipher
	data   map[string]string
}

// NewEncryptedEnvironment Storage creates new environment storage
func NewEncMapStorage(master, dump string) (*EncMapStorage, error) {
	cipher, err := crypto.NewAesGcmCypher(master)
	if err != nil {
		return nil, err
	}
	s := &EncMapStorage{cipher, map[string]string{}}
	// decrypt data if not empty
	if dump != "" {
		bct, err := crypto.B64ToByte(dump)
		if err != nil {
			return nil, err
		}
		b, err := s.cipher.Decrypt(bct)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(b, &s.data); err != nil {
			return nil, err
		}
	}
	return s, nil
}

// Get reads secret from storage
func (s *EncMapStorage) Get(name string) (string, error) {
	v, ok := s.data[name]
	if !ok {
		return "", fmt.Errorf("%w:%s", ErrSecretNotFound, name)
	}
	return v, nil
}

// Set the secret value
func (s *EncMapStorage) Set(name, secret string) {
	s.data[name] = secret
}

func (s *EncMapStorage) Dump() (string, error) {
	if len(s.data) == 0 {
		return "", nil
	}
	b, err := json.Marshal(&s.data)
	if err != nil {
		return "", err
	}
	ct, err := s.cipher.Encrypt(b)
	if err != nil {
		return "", err
	}
	return crypto.ByteToB64(ct), nil
}
