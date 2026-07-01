package utils

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func GenerateAllTokens(email string, firstName string, lastName string, userType *string, userID string) (signedToken string, signedRefreshToken string, err error) {

}