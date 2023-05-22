package database

import (
	"context"
	"go-mongo/models"
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

	// Insert test data
	collection.InsertOne(context.TODO(), models.Album{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99})

	// Assign the collection to the database
	Collection = collection
}

func (db *DatabaseInstance) Disconnect() {
	log.Println("Disconnecting from MongoDB...")

	// Disconnect from MongoDB
	err := db.Client.Disconnect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Disconnected from MongoDB!")
}

func checkDatabaseConnection() {
	if Database == nil {
		log.Fatal("Database is not connected!")
	}

	if err := Database.Client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}
}

func insertAlbum(album models.Album) {
	checkDatabaseConnection()

	// Insert a single document
	insertResult, err := Collection.InsertOne(context.Background(), album)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)
}
