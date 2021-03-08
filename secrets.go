package secrets

//Storage provides the interface for reading secrets
type Storage interface {
	Get(name string) (string, error)
}

var storage Storage

// Get provides the proxy to Storage.Get of initialized package storage
func Get(name string) (string, error) {
	if storage == nil {
		return "", &StorageIsNotInitialized{}
	}
	return storage.Get(name)
}
