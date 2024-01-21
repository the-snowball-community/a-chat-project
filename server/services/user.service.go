package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"snowball-community.com/chat/db"
	"snowball-community.com/chat/models"
)

func CreateUser(newUser models.User) primitive.ObjectID {
	collection := db.GetCollection("users")
	result, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return result.InsertedID.(primitive.ObjectID)
}
