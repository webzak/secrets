package secrets

// Storage provides the interface for reading secrets
type Storage interface {
	Get(name string) (string, error)
}
