package services

import (
	"errors"
	"jwt/models"
	"time"

	repository "jwt/repositories"
	"jwt/utils"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)
var validate = validator.New()

// hashes the user password
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
  return string(hashedPassword), nil
}

// checks if the provided password matches
func verifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}


func CreateUser(user models.User) (*models.User, error) {
	if err := validate.Struct(user); err != nil {
		return nil, err
	}
	exists, err := repository.EmailExists(*user.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("email already exists")
	}
	// Hash password
	hashedPassword, err := hashPassword(*user.Password); 
	if err != nil {
		return nil, err
	}
	user.Password = &hashedPassword
	user.Created_At, _ = time.Parse(time.Now().Format(time.RFC3339), time.RFC3339)	
	user.Updated_At, _ = time.Parse(time.Now().Format(time.RFC3339), time.RFC3339)
	user.ID = bson.NewObjectID()
	user.User_ID = user.ID.Hex()
	token, refreshToken, _ := utils.GenerateAllTokens(*user.Email, *user.First_Name, *user.Last_Name, user.User_Type, user.User_ID)
	user.Token = &token
	user.Refresh_Token = &refreshToken

	result, err := repository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return result, nil
 
}