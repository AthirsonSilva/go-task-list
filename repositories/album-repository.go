package repositories

import (
	"context"
	"go-mongo/database"
	"go-mongo/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAll() ([]models.Album, error) {
	var albums []models.Album

	// Get all albums from the database
	cursor, err := database.Collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the returned cursor
	for cursor.Next(context.Background()) {
		var album models.Album

		cursor.Decode(&album)

		albums = append(albums, album)
	}

	// Return all albums
	return albums, nil
}

func FindById(id string) (models.Album, error) {
	var album models.Album

	// Get album from the database
	err := database.Collection.FindOne(context.Background(), bson.D{}).Decode(&album)

	if err != nil {
		log.Fatal(err)
	}

	// Return album
	return album, nil
}

func Create(album models.Album) (models.Album, error) {
	// Set ID
	album.ID = primitive.NewObjectID()
	album.CreatedAt = time.Now()
	album.UpdatedAt = time.Now()

	// Insert album in the database
	_, err := database.Collection.InsertOne(context.Background(), album)

	if err != nil {
		log.Fatal(err)
	}

	// Return album
	return album, nil
}

func Update(id string, album models.Album) (models.Album, error) {
	// Convert the id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	// Check if album exists in the database
	err = database.Collection.FindOne(context.Background(), bson.M{"_id": oid}).Err()

	if err != nil {
		return album, err
	}

	// Set updated_at
	album.UpdatedAt = time.Now()
	album.ID = oid

	// Update album in the database
	_, err = database.Collection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.D{{Key: "$set", Value: album}})

	if err != nil {
		log.Fatal(err)
	}

	// Return album
	return album, nil
}

func Delete(id string) error {
	// Convert the id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	// Check if album exists in the database
	err = database.Collection.FindOne(context.Background(), bson.M{"_id": oid}).Err()

	if err != nil {
		return err
	}

	// Delete album in the database
	_, err = database.Collection.DeleteOne(context.Background(), bson.M{"_id": oid})

	if err != nil {
		log.Fatal(err)
	}

	// Return album
	return nil
}
