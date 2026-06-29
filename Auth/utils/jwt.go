package utils

import (
	"errors"
	"net/http"
	"os"
	"time"

	"auth/models"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a new JWT token for a user
func GenerateJWT(user models.User, w http.ResponseWriter) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	if len(secretKey) == 0 {
		return "", errors.New("JWT_SECRET environment variable is not set")
	}
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		// expiresAt
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		// issuedAt
		"iat": time.Now().Unix(),
		// issuer
		"iss": "auth-service",
	}
	// 1, generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	// 2. Set as cookie (good for browser clients)
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt-token",
		Value:    signed,
		Path:     "/",
		MaxAge:   24 * 60 * 60, // matches the 24h token expiry
		HttpOnly: true,         // not accessible via JS — mitigates XSS
		// Secure:   true,         // only sent over HTTPS
		SameSite: http.SameSiteLaxMode,
	})

	return signed, err
}
