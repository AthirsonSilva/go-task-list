package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseInstance struct {
	Client *mongo.Client
}

var (
	Database        *DatabaseInstance
	AlbumCollection *mongo.Collection
	UserCollection  *mongo.Collection
)

func (db *DatabaseInstance) Connect() {
	log.Println("Connecting to MongoDB...")

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	e := client.Ping(context.Background(), nil)

	if e != nil {
		log.Fatal(e)
	}

	log.Println("Connected to MongoDB!")

	AlbumCollection = client.Database("music-api").Collection("albums")
	UserCollection = client.Database("music-api").Collection("users")

	Database = &DatabaseInstance{Client: client}
}
