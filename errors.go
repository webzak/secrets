package secrets

type StorageIsNotInitialized struct{}

func (e *StorageIsNotInitialized) Error() string {
	return "secrets package requires the storage initialization before use"
}

type SecretNotFound struct {
	name string
}

func (e *SecretNotFound) Error() string {
	return "secret not found: " + e.name
}
