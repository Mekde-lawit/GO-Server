package services

import (
	"errors"
	
	"jwt/models"
	repository "jwt/repositories"

	"github.com/gin-gonic/gin"
)


func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthoized to access this resource!")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userID string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userID {
		err = errors.New("Unauthorized to access this resource!")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

func GetUserByID(c *gin.Context, id string) (*models.User, error) {

    if err := MatchUserTypeToUid(c, id); err != nil {
        return nil, err
    }

    return repository.GetUserByID(id)
}

func GetAllUsers(c *gin.Context) ([]*models.User, error) {

	if err := CheckUserType(c, "ADMIN"); err != nil {
		return nil, err
	}
	return repository.GetAllUsers()
}