package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	User_ID       string        `bson:"user_id"`
	First_Name    *string       `bson:"first_name" validate:"required,min=2,max=100"`
	Last_Name     *string       `bson:"last_name" validate:"required,min=2,max=100"`
	Email         *string       `bson:"email" validate:"required,email"`
	Password      *string       `bson:"password" validate:"required,min=6"`
	Phone         *string       `bson:"phone" validate:"required"`
	User_Type     *string       `bson:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Token         *string       `bson:"token"`
	Refresh_Token *string       `bson:"refresh_token"`
	Created_At    time.Time     `bson:"created_at"`
	Updated_At    time.Time     `bson:"updated_at"`
}