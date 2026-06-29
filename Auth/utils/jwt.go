package utils

import (
	"errors"
	"time"
	
	"github.com/golang-jwt/jwt/v5"
	"auth/models"
)

// Common JWT errors
var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
)

// GenerateJWT creates a new JWT token for a user
func GenerateJWT(user models.User, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
		"iss":     "auth-service",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateJWT parses and validates a JWT token
func ValidateJWT(tokenString string, secretKey string) (*jwt.Token, error) {
	if tokenString == "" {
		return nil, ErrInvalidToken
	}
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		// Check if token is expired
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, err
	}

	// Verify token is valid
	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return token, nil
}

// ExtractClaims extracts claims from a validated token
func ExtractClaims(token *jwt.Token) (jwt.MapClaims, error) {
	if token == nil {
		return nil, errors.New("token is nil")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims format")
	}

	return claims, nil
}

// GetUserIDFromToken extracts user ID from token claims
func GetUserIDFromToken(token *jwt.Token) (uint, error) {
	claims, err := ExtractClaims(token)
	if err != nil {
		return 0, err
	}

	userID, ok := claims["user_id"].(float64) // JWT numbers are float64 by default
	if !ok {
		return 0, errors.New("user_id not found or invalid type")
	}

	return uint(userID), nil
}

// GetEmailFromToken extracts email from token claims
func GetEmailFromToken(token *jwt.Token) (string, error) {
	claims, err := ExtractClaims(token)
	if err != nil {
		return "", err
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("email not found or invalid type")
	}

	return email, nil
}

// GetRoleFromToken extracts role from token claims
func GetRoleFromToken(token *jwt.Token) (string, error) {
	claims, err := ExtractClaims(token)
	if err != nil {
		return "", err
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", errors.New("role not found or invalid type")
	}

	return role, nil
}

// IsTokenExpired checks if a token is expired
func IsTokenExpired(token *jwt.Token) bool {
	claims, err := ExtractClaims(token)
	if err != nil {
		return true
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return true
	}

	expTime := time.Unix(int64(exp), 0)
	return time.Now().After(expTime)
}

// RefreshToken generates a new token with extended expiration
func RefreshToken(tokenString string, secretKey string) (string, error) {
	token, err := ValidateJWT(tokenString, secretKey)
	if err != nil && !errors.Is(err, ErrExpiredToken) {
		return "", err
	}

	claims, err := ExtractClaims(token)
	if err != nil {
		return "", err
	}

	// Create new claims with extended expiration
	newClaims := jwt.MapClaims{
		"user_id": claims["user_id"],
		"email":   claims["email"],
		"role":    claims["role"],
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
		"iss":     "auth-service",
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	return newToken.SignedString([]byte(secretKey))
}