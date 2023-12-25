package secrets

import (
	"fmt"
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
	return &EnvStorage{prefix, uppercase}
}

// Get reads secret from storage
func (es *EnvStorage) Get(name string) (string, error) {
	secret, ok := os.LookupEnv(es.makeEnvName(name))
	if !ok {
		return "", fmt.Errorf("%w:%s", ErrSecretNotFound, name)
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
