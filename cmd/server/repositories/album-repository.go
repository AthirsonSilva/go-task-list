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

func FindAllAlbums() ([]models.Album, error) {
	var albums []models.Album

	cursor, err := database.AlbumCollection.Find(context.Background(), bson.M{})
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

func FindAlbumById(id string) (models.Album, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	var album models.Album

	if err != nil {
		return album, err
	}

	err = database.AlbumCollection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&album)
	if err != nil {
		return album, err
	}

	err = database.AlbumCollection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&album)
	if err != nil {
		return album, err
	}

	return album, nil
}

func CreateAlbum(album models.Album) (models.Album, error) {
	album.ID = primitive.NewObjectID()
	album.CreatedAt = time.Now()
	album.UpdatedAt = time.Now()

	_, err := database.AlbumCollection.InsertOne(context.Background(), album)
	if err != nil {
		log.Fatal(err)
	}

	return album, nil
}

func UpdateAlbumById(id string, album models.Album) (models.Album, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return album, err
	}

	err = database.AlbumCollection.FindOne(context.Background(), bson.M{"_id": oid}).Err()
	if err != nil {
		return album, err
	}

	album.UpdatedAt = time.Now()
	album.ID = oid

	_, err = database.AlbumCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": oid},
		bson.D{{Key: "$set", Value: album}},
	)
	if err != nil {
		log.Fatal(err)
	}

	return album, nil
}

func DeleteAlbumById(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	err = database.AlbumCollection.FindOne(context.Background(), bson.M{"_id": oid}).Err()
	if err != nil {
		return err
	}

	_, err = database.AlbumCollection.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
