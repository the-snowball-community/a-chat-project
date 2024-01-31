package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"snowball-community.com/chat/db"
	"snowball-community.com/chat/models"
)

const CHAT_COLLECTION_NAME = "chats"

func CreateChat(newChat models.Chat) primitive.ObjectID {
	collection := db.GetCollection(CHAT_COLLECTION_NAME)
	result, err := collection.InsertOne(context.TODO(), newChat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return result.InsertedID.(primitive.ObjectID)
}

func FindMany(limit, skip int64) []models.Chat {
	collection := db.GetCollection(CHAT_COLLECTION_NAME)
	// cursor, err := collection.Find(context.TODO(), bson.D{{Key: "userName", Value: "test"}})
	cursor, err := collection.Find(context.TODO(), options.FindOptions{Limit: &limit, Skip: &skip})
	if err != nil {
		panic(err)
	}
	var chats []models.Chat
	for cursor.Next(context.TODO()) {
		var chat models.Chat
		if err := cursor.Decode(&chat); err != nil {
			panic(err)
		}
		chats = append(chats, chat)
	}
	return chats
}
