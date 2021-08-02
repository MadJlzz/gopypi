package auth

import (
	"encoding/base64"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type authenticationMiddleware struct {
	logger       *zap.SugaredLogger
	authentifier Authentifier
}

func NewAuthenticationMiddleware(logger *zap.SugaredLogger, authentifier Authentifier) *authenticationMiddleware {
	return &authenticationMiddleware{
		logger:       logger,
		authentifier: authentifier,
	}
}

// HandleAuthentication FIXME: writing this logic like this break the navigation on the page for oauth2 authenticated clients.
func (am *authenticationMiddleware) HandleAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		encodedPart := strings.Split(authHeader, " ")[1]
		decodedBytes, err := base64.StdEncoding.DecodeString(encodedPart)
		if err != nil {
			am.logger.Errorf("decoding authorization with basic authentication failed. got: %v", err)
			http.Error(w, "You are not allowed to see this page.", http.StatusForbidden)
		}
		user, password := func() (string, string) {
			s := strings.SplitN(string(decodedBytes), ":", 2)
			return s[0], s[1]
		}()
		err = am.authentifier.authenticate(user, password)
		if err != nil {
			http.Error(w, "You are not allowed to see this page.", http.StatusForbidden)
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
