package router

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var (
	nextSecret = os.Getenv("PENNYPINCHER_NEXTSECRET")
)

// authNMiddleware authenticates the request
func authNMiddleware(next http.Handler, logger *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := nextSecret
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString = extractTokenFromHeader(tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			logger.Info("Error: ", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			logger.Info("Claims: ", claims)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getAuthNMiddleware(logger *zap.SugaredLogger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return authNMiddleware(next, logger)
	}
}

// Extracts the token value from the Authorization header
func extractTokenFromHeader(header string) string {
	// Split the header value by whitespace
	split := strings.SplitN(header, " ", 2)

	if len(split) != 2 || strings.ToLower(split[0]) != "bearer" {
		log.Fatal("Invalid Authorization header format")
	}

	return split[1]
}
