package crypto

type CryptoError struct {
	s string
}

func (e *CryptoError) Error() string {
	return "CryptoError: " + e.s
}
