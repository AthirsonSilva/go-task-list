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

var Database *DatabaseInstance
var Collection *mongo.Collection

func (db *DatabaseInstance) Connect() {
	log.Println("Connecting to MongoDB...")

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	e := client.Ping(context.Background(), nil)

	if e != nil {
		log.Fatal(e)
	}

	log.Println("Connected to MongoDB!")

	// Get the collection as ref
	collection := client.Database("go-rest").Collection("albums")

	// Assign the collection to the database
	Database = &DatabaseInstance{Client: client}

	// Assign the collection to the database
	Collection = collection
}
