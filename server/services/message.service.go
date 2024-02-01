package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"snowball-community.com/chat/db"
	"snowball-community.com/chat/models"
)

const CHAT_COLLECTION_NAME = "messages"

func CreateMessage(newChat models.Message) primitive.ObjectID {
	collection := db.GetCollection(CHAT_COLLECTION_NAME)
	result, err := collection.InsertOne(context.TODO(), newChat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return result.InsertedID.(primitive.ObjectID)
}

func FindMany(limit, skip int64) []models.Message {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	opts.SetLimit(limit)
	opts.SetSkip(skip)
	collection := db.GetCollection(CHAT_COLLECTION_NAME)
	// cursor, err := collection.Find(context.TODO(), bson.D{{Key: "userName", Value: "test"}})
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	var messages []models.Message
	for cursor.Next(context.TODO()) {
		var message models.Message
		if err := cursor.Decode(&message); err != nil {
			panic(err)
		}
		messages = append(messages, message)
	}
	return messages
}
