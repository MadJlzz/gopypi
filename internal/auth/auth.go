package auth

type Authentifier interface {
	authenticate(user, password string) error
}

type ApiKey struct {

}

func (ak *ApiKey) authenticate(user, password string) error {
	return nil
}