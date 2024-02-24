package database

import (
	"context"
	"log"
	"time"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/brianvoe/gofakeit/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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

	go migrateData(AlbumCollection, "album")
	go migrateData(UserCollection, "user")
}

func migrateData(mongoCollection *mongo.Collection, collectionName string) {
	rows, err := mongoCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}

	if rows > 0 {
		return
	}

	for i := range 10 {
		log.Printf("Inserting %s %d", collectionName, i)
		model := generateModel(collectionName)
		_, err = mongoCollection.InsertOne(context.Background(), model)
		if err != nil {
			log.Println(err)
		}
	}
}

func generateModel(entity string) any {
	gofakeit.Seed(time.Now().UnixNano())

	switch entity {
	case "album":
		return models.Album{
			Artist:    gofakeit.Name(),
			Title:     gofakeit.Sentence(3),
			Price:     gofakeit.Price(0, 100),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	case "user":
		password := gofakeit.Password(true, true, true, true, true, 8)

		pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
		}

		return models.User{
			Username:  gofakeit.Username(),
			Email:     gofakeit.Email(),
			Password:  string(pass),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	default:
		log.Println("Invalid entity")
		return nil
	}
}
