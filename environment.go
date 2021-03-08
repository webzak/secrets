package secrets

import (
	"os"
	"strings"
)

// EnvironmentStorage stores secrets in environment variables
type EnvironmentStorage struct {
	prefix    string
	uppercase bool
}

// InitMemoryStorage creates new environment storage and sets it as package storage
func InitEnvironmentStorage(prefix string, uppercase bool) *EnvironmentStorage {
	es := &EnvironmentStorage{prefix, uppercase}
	storage = es
	return es
}

// Get reads secret from storage
func (es *EnvironmentStorage) Get(name string) (string, error) {
	secret, ok := os.LookupEnv(es.makeEnvName(name))
	if !ok {
		return "", &SecretNotFound{name}
	}
	return secret, nil
}

// Set sets the secret value in memory storage
func (es *EnvironmentStorage) Set(name, secret string) error {
	return os.Setenv(es.makeEnvName(name), secret)
}

func (es *EnvironmentStorage) makeEnvName(name string) string {
	en := es.prefix + name
	if es.uppercase {
		en = strings.ToUpper(en)
	}
	return en
}
