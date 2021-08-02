package auth

import (
	"encoding/base64"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type AuthenticationMiddlewareFunc func(method Authenticator) mux.MiddlewareFunc

type authenticationMiddleware struct {
	logger          *zap.SugaredLogger
	isAuthenticated bool
}

func NewAuthenticationMiddleware(logger *zap.SugaredLogger) *authenticationMiddleware {
	return &authenticationMiddleware{
		logger:       logger,
	}
}

// HandleCloudIAPAuthentication checks if a user that connects through Cloud IAP is authenticated.
//func (am *authenticationMiddleware) HandleCloudIAPAuthentication(next http.Handler, authenticator Authenticator) http.Handler {
//	am.isAuthenticated = false
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		authHeader := r.Header.Get("Cookie")
//		if strings.Contains(authHeader, "GCP_IAP_UID") {
//			ok := authenticator.authenticate(authHeader, "")
//			if !ok {
//				http.Error(w, "You are not allowed to see this page.", http.StatusForbidden)
//				return
//			}
//			am.isAuthenticated = true
//		}
//		// Call the next handler, which can be another middleware in the chain, or the final handler.
//		next.ServeHTTP(w, r)
//	})
//}

func (am *authenticationMiddleware) HandleBasicAuthentication(authenticator Authenticator) mux.MiddlewareFunc {
	am.isAuthenticated = false
	return func(next http.Handler) http.Handler {
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
				user, password := func() (string, string) {
					s := strings.SplitN(string(decodedBytes), ":", 2)
					return s[0], s[1]
				}()
				am.logger.Infof("%s, %s", user, password)
				ok := authenticator.authenticate()
				if !ok {
					http.Error(w, "You are not allowed to see this page.", http.StatusForbidden)
					return
				}
				am.isAuthenticated = true
			}
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
}

// HandleNoAuthentication checks if the client is actually authenticated, if not reject it.
func (am *authenticationMiddleware) HandleNoAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !am.isAuthenticated {
			http.Error(w, "You are not allowed to see this page.", http.StatusForbidden)
			return
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}