package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrNoToken          = errors.New("no token found")
	ErrMalformedHeader  = errors.New("malformed Authorization header")
	ErrInvalidToken     = errors.New("invalid token")
	ErrMissingSecretKey = errors.New("JWT_SECRET environment variable is not set")
)


var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// extractToken reads the token from the Authorization header or cookie.
func extractToken(r *http.Request) (string, error) {
	if h := r.Header.Get("Authorization"); h != "" {
		parts := strings.SplitN(h, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
			return strings.TrimSpace(parts[1]), nil
		}
		return "", ErrMalformedHeader
	}

	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		return "", ErrNoToken
	}
	return cookie.Value, nil
}

func parseAndValidate(tokenString string) (jwt.MapClaims, error) {
	if tokenString == "" {
		return nil, ErrNoToken
	}
	if len(jwtSecret) == 0 {
		return nil, ErrMissingSecretKey
	}

	token, err := jwt.Parse(
		tokenString,
		func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}


func writeAuthError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrMissingSecretKey):
		http.Error(w, "internal server error", http.StatusInternalServerError)
	case errors.Is(err, ErrNoToken), errors.Is(err, ErrMalformedHeader), errors.Is(err, ErrInvalidToken):
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	default:
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}
}

type claimsContextKey struct{}

// ClaimsFromContext retrieves the JWT claims stored by the Auth middleware.
// Handlers downstream call this to access the authenticated user's claims.
func ClaimsFromContext(ctx context.Context) (jwt.MapClaims, bool) {
	claims, ok := ctx.Value(claimsContextKey{}).(jwt.MapClaims)
	return claims, ok
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := extractToken(r)
		if err != nil {
			writeAuthError(w, err)
			return
		}

		claims, err := parseAndValidate(tokenString)
		if err != nil {
			writeAuthError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), claimsContextKey{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}