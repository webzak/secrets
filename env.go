package secrets

import (
	"os"
	"strings"
)

// EnvStorage stores secrets in environment variables
type EnvStorage struct {
	prefix    string
	uppercase bool
}

// NewMemoryStorage creates new environment storage
func NewEnvStorage(prefix string, uppercase bool) *EnvStorage {
	es := &EnvStorage{prefix, uppercase}
	storage = es
	return es
}

// InitMemoryStorage creates new environment storage and sets it as package storage
func InitEnvStorage(prefix string, uppercase bool) *EnvStorage {
	es := NewEnvStorage(prefix, uppercase)
	storage = es
	return es
}

// Get reads secret from storage
func (es *EnvStorage) Get(name string) (string, error) {
	secret, ok := os.LookupEnv(es.makeEnvName(name))
	if !ok {
		return "", &SecretNotFound{name}
	}
	return secret, nil
}

// Set sets the secret value in memory storage
func (es *EnvStorage) Set(name, secret string) error {
	return os.Setenv(es.makeEnvName(name), secret)
}

func (es *EnvStorage) makeEnvName(name string) string {
	en := es.prefix + name
	if es.uppercase {
		en = strings.ToUpper(en)
	}
	return en
}
