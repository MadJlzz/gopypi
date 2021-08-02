package auth

type Authenticator interface {
	authenticate() bool
}

type ApiKey struct {
	Key string
}

func (ak *ApiKey) authenticate() bool {
	return true
}
