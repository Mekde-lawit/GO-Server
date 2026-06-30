package controllers

import (
	"jwt/configs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection = configs.OpenCollection(configs.DBinstance(), "users")
var validate = validator.New()

func Login() {

}

func Signup() {

}
