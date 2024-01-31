package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"snowball-community.com/chat/db"
	"snowball-community.com/chat/models"
)

func CreateChat(newChat models.Chat) primitive.ObjectID {
	collection := db.GetCollection("chats")
	result, err := collection.InsertOne(context.TODO(), newChat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return result.InsertedID.(primitive.ObjectID)
}
