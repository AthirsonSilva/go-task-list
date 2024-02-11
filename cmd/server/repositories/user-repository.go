package repositories

import (
	"context"
	"log"
	"time"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/database"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUserById(id string) (models.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	var user models.User

	if err != nil {
		return user, err
	}

	err = database.UserCollection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return user, err
	}

	err = database.UserCollection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := database.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}
