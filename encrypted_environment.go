package secrets

import (
	"github.com/webzak/secrets/crypto"
)

// EncryptedEnvironmentStorage stores secrets in environment variables
type EncryptedEnvironmentStorage struct {
	cipher crypto.Cipher
	es     *EnvironmentStorage
}

// NewEncryptedEnvironment Storage creates new environment storage
func NewEncryptedEnvironmentStorage(master, prefix string, uppercase bool) (*EncryptedEnvironmentStorage, error) {
	cipher, err := crypto.NewAesGcmCypher(master)
	if err != nil {
		return nil, err
	}
	es := &EnvironmentStorage{prefix, uppercase}
	ees := &EncryptedEnvironmentStorage{cipher, es}
	storage = ees
	return ees, nil
}

// InitEncryptedMemoryStorage creates new environment storage and sets it as package storage
func InitEncryptedEnvironmentStorage(master, prefix string, uppercase bool) (*EncryptedEnvironmentStorage, error) {
	ees, err := NewEncryptedEnvironmentStorage(master, prefix, uppercase)
	if err != nil {
		return nil, err
	}
	storage = ees
	return ees, nil
}

// creates MustInitMemoryStorage new environment storage and sets it as package storage
func MustInitEncryptedEnvironmentStorage(master, prefix string, uppercase bool) *EncryptedEnvironmentStorage {
	ees, err := InitEncryptedEnvironmentStorage(master, prefix, uppercase)
	if err != nil {
		panic(err)
	}
	return ees
}

// Get reads secret from storage
func (ees *EncryptedEnvironmentStorage) Get(name string) (string, error) {
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

// Set sets the secret value in memory storage
func (ees *EncryptedEnvironmentStorage) Set(name, secret string) error {
	ct, err := ees.cipher.Encrypt([]byte(secret))
	if err != nil {
		return err
	}
	return ees.es.Set(name, crypto.ByteToB64(ct))
}
