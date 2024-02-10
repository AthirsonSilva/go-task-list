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

func FindAll() ([]models.Album, error) {
	var albums []models.Album

	cursor, err := database.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.Background()) {
		var album models.Album

		cursor.Decode(&album)

		albums = append(albums, album)
	}

	return albums, nil
}

func FindById(id string) (models.Album, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	var album models.Album

	if err != nil {
		return album, err
	}

	err = database.Collection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&album)
	if err != nil {
		return album, err
	}

	err = database.Collection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&album)
	if err != nil {
		return album, err
	}

	return album, nil
}

func Create(album models.Album) (models.Album, error) {
	album.ID = primitive.NewObjectID()
	album.CreatedAt = time.Now()
	album.UpdatedAt = time.Now()

	_, err := database.Collection.InsertOne(context.Background(), album)
	if err != nil {
		log.Fatal(err)
	}

	return album, nil
}

func Update(id string, album models.Album) (models.Album, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return album, err
	}

	err = database.Collection.FindOne(context.Background(), bson.M{"_id": oid}).Err()
	if err != nil {
		return album, err
	}

	album.UpdatedAt = time.Now()
	album.ID = oid

	_, err = database.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": oid},
		bson.D{{Key: "$set", Value: album}},
	)
	if err != nil {
		log.Fatal(err)
	}

	return album, nil
}

func Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Collection.FindOne(context.Background(), bson.M{"_id": oid}).Err()
	if err != nil {
		return err
	}

	_, err = database.Collection.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
