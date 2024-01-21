package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Database

func GetDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	uri := os.Getenv("DB_URL")
	databaseName := os.Getenv("DB_NAME")
	opts := options.Client().ApplyURI(uri)
	connection, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	client = connection.Database(databaseName)
}

func GetCollection(name string) *mongo.Collection {
	return client.Collection(name)
}
