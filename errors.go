package secrets

type StorageIsNotInitialized struct{}

func (e *StorageIsNotInitialized) Error() string {
	return "secrets package requires the storage initialization before use"
}

func IsNotInitialized(err error) bool {
	_, ok := err.(*StorageIsNotInitialized)
	return ok
}

type SecretNotFound struct {
	name string
}

func (e *SecretNotFound) Error() string {
	return "secret not found: " + e.name
}

func IsNotFound(err error) bool {
	_, ok := err.(*SecretNotFound)
	return ok
}
