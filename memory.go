package secrets

import "fmt"

// MemoryStorage does not persist secrets data
type MemoryStorage struct {
	data map[string]string
}

// NewMemoryStorage creates new memory storage
func NewMemoryStorage() *MemoryStorage {
	ms := &MemoryStorage{}
	ms.data = make(map[string]string)
	return ms
}

// Get reads secret from storage
func (ms *MemoryStorage) Get(name string) (string, error) {
	secret, ok := ms.data[name]
	if !ok {
		return "", fmt.Errorf("%w : %s", ErrSecretNotFound, name)
	}
	return secret, nil
}

// Set sets the secret value in memory storage
func (ms *MemoryStorage) Set(name, secret string) {
	ms.data[name] = secret
}
