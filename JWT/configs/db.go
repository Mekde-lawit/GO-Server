package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBinstance() *mongo.Client {

	InitEnv()

	mongoDB := os.Getenv("MONGODB_URL")

	client, err := mongo.Connect(options.Client().ApplyURI(mongoDB))
	if err != nil {
		log.Fatal("Error: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println("Connected to Mongodb!")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(collectionName string) *mongo.Collection {
	var collection *mongo.Collection = Client.Database("goDB").Collection(collectionName)
	return collection
}
