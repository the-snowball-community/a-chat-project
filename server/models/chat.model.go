package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId    string             `bson:"userId" json:"userId"`
	RoomId    string             `bson:"roomId" json:"roomId"`
	UserName  string             `bson:"userName" json:"userName"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt int64              `bson:"createdAt" json:"createdAt"`
}
