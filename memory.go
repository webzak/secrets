package secrets

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

// InitMemoryStorage creates new memory storage and sets it as package storage
func InitMemoryStorage() *MemoryStorage {
	ms := NewMemoryStorage()
	storage = ms
	return ms
}

// Get reads secret from storage
func (ms *MemoryStorage) Get(name string) (string, error) {
	secret, ok := ms.data[name]
	if !ok {
		return "", &SecretNotFound{name}
	}
	return secret, nil
}

// Set sets the secret value in memory storage
func (ms *MemoryStorage) Set(name, secret string) {
	ms.data[name] = secret
}
