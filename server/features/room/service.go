package room

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"snowball-community.com/chat/db"
	"snowball-community.com/chat/models"
)

func CreateRoom(newRoom models.Room) primitive.ObjectID {
	collection := db.GetCollection("rooms")
	result, err := collection.InsertOne(context.TODO(), newRoom)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return result.InsertedID.(primitive.ObjectID)
}

func GetRoom() models.Room {
	filter := bson.D{}
	collection := db.GetCollection("rooms")

	var result models.Room
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}
