package secrets

import (
	"github.com/webzak/secrets/crypto"
)

// EncryptedEnvironmentStorage stores secrets in environment variables
type EncEnvStorage struct {
	cipher crypto.Cipher
	es     *EnvStorage
}

// NewEncryptedEnvironment Storage creates new environment storage
func NewEncEnvStorage(master, prefix string, uppercase bool) (*EncEnvStorage, error) {
	cipher, err := crypto.NewAesGcmCypher(master)
	if err != nil {
		return nil, err
	}
	es := &EnvStorage{prefix, uppercase}
	ees := &EncEnvStorage{cipher, es}
	storage = ees
	return ees, nil
}

// InitEncryptedMemoryStorage creates new environment storage and sets it as package storage
func InitEncEnvStorage(master, prefix string, uppercase bool) (*EncEnvStorage, error) {
	ees, err := NewEncEnvStorage(master, prefix, uppercase)
	if err != nil {
		return nil, err
	}
	storage = ees
	return ees, nil
}

// creates MustInitMemoryStorage new environment storage and sets it as package storage
func MustInitEncEnvStorage(master, prefix string, uppercase bool) *EncEnvStorage {
	ees, err := InitEncEnvStorage(master, prefix, uppercase)
	if err != nil {
		panic(err)
	}
	return ees
}

// Get reads secret from storage
func (ees *EncEnvStorage) Get(name string) (string, error) {
	ct, err := ees.es.Get(name)
	if err != nil {
		return "", err
	}
	bct, err := crypto.B64ToByte(ct)
	if err != nil {
		return "", err
	}
	secret, err := ees.cipher.Decrypt(bct)
	if err != nil {
		return "", err
	}
	return string(secret), nil
}

// Set the secret value
func (ees *EncEnvStorage) Set(name, secret string) error {
	ct, err := ees.cipher.Encrypt([]byte(secret))
	if err != nil {
		return err
	}
	return ees.es.Set(name, crypto.ByteToB64(ct))
}

// Prepare secret value name and encrypted value
func (ees *EncEnvStorage) Prepare(name, secret string) (string, string, error) {
	ename := ees.es.makeEnvName(name)
	ct, err := ees.cipher.Encrypt([]byte(secret))
	if err != nil {
		return "", "", err
	}
	return ename, crypto.ByteToB64(ct), nil
}
