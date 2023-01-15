package router

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/kr/pretty"
	"go.uber.org/zap"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

var (
	aud    = []string{os.Getenv("PENNYPINCHER_AUTH0AUDIENCE")}
	domain = os.Getenv("PENNYPINCHER_AUTH0DOMAIN")
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// HasScope checks whether our claims have a specific scope.
func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}

// authZMiddleware authorizes the request
func authZMiddleware(next http.Handler, logger *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issuerURL, err := url.Parse("https://" + domain + "/")
		if err != nil {
			log.Fatalf("Failed to parse the issuer url: %v", err)
		}
		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			aud,
			validator.WithCustomClaims(
				func() validator.CustomClaims {
					return &CustomClaims{}
				},
			),
			validator.WithAllowedClockSkew(time.Minute),
		)
		if err != nil {
			log.Fatalf("Failed to set up the jwt validator")
		}
		pretty.Print("____________________________________________________\n")
		pretty.Print(issuerURL.String())
		pretty.Print("____________________________________________________\n")
		pretty.Print(aud)
		pretty.Print("____________________________________________________\n")
		pretty.Print(validator.RS256)
		pretty.Print("____________________________________________________\n")
		pretty.Print("____________________________________________________\n")
		pretty.Print(err)

		if err != nil {
			log.Fatalf("failed to set up the validator: %v", err)
		}

		errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("Encountered error while validating JWT: %v", err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message":"Failed to validate JWT."}`))
		}

		middleware := jwtmiddleware.New(
			jwtValidator.ValidateToken,
			jwtmiddleware.WithErrorHandler(errorHandler),
		)

		// Apply middleware to the request and call the next handler
		middleware.CheckJWT(next).ServeHTTP(w, r)
	})
}

func getAuthZMiddleware(logger *zap.SugaredLogger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return authZMiddleware(next, logger)
	}
}
