package database

import (
    "context"
		"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"

		"github.com/joho/godotenv"
		"os"
		"fmt"
)

var client *mongo.Client
var database *mongo.Database

func InitMongoDBConnection() error {
	godotenv.Load()
	uri := os.Getenv("MONGODB_URI")

	fmt.Println(uri)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return nil
}

func GetDatabase() *mongo.Collection {
	return client.Database("developerDB").Collection("receipts")
}