package auth

import (
	"encoding/base64"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strings"
)

type authenticationMiddleware struct {
	logger *zap.SugaredLogger
	apiKey string
}

func NewAuthenticationMiddleware(logger *zap.SugaredLogger) *authenticationMiddleware {
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		logger.Fatalf("API_KEY variable should be set. ")
	}
	return &authenticationMiddleware{
		logger: logger,
		apiKey: apiKey,
	}
}

func (am *authenticationMiddleware) HandleBasicAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if strings.Contains(authHeader, "Basic") {
			encodedPart := strings.Split(authHeader, " ")[1]
			decodedBytes, err := base64.StdEncoding.DecodeString(encodedPart)
			if err != nil {
				am.logger.Errorf("decoding authorization with basic authentication failed. got: %v", err)
				http.Error(w, "You are not allowed to see this page.", http.StatusForbidden)
				return
			}
			user, _ := func() (string, string) {
				s := strings.SplitN(string(decodedBytes), ":", 2)
				return s[0], s[1]
			}()
			if user == am.apiKey {
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "You are not allowed to see this page.", http.StatusForbidden)
	})
}
