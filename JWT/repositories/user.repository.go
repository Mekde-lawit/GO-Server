package repository

import (
	"context"
	"time"

	"jwt/configs"
	"jwt/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection = configs.OpenCollection("users")


func GetUserByID(id string) (*models.User, error) {

    ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
    defer cancel()

    var user models.User

    err := userCollection.FindOne(
        ctx,
        bson.M{"user_id": id},
    ).Decode(&user)

    if err != nil {
        return nil, err
    }

    return &user, nil
}

func GetAllUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func EmailExists(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}