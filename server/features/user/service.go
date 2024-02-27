package user

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"snowball-community.com/chat/db"
)

func CreateUser(newUser User) primitive.ObjectID {
	collection := db.GetCollection("users")
	result, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return result.InsertedID.(primitive.ObjectID)
}

func CreateRandomUserName() string {

	firstNames := []string{"잘생긴", "똑똑한", "용감한", "귀여운", "성실한", "음흉한", "멋진", "섹시한", "든든한", "무뚝뚝한", "빛나는", "이기적인"}
	lastNames := []string{"호랑이", "사자", "팬더", "사슴", "고양이", "여우", "곰", "원숭이", "개미핡기", "코끼리", "고래", "너구리", "라쿤"}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	return firstName + lastName
}
