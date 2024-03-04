package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/brianvoe/gofakeit/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Instance struct {
	Client *mongo.Client
}

var (
	Database       *Instance
	TaskCollection *mongo.Collection
	UserCollection *mongo.Collection
)

func (db *Instance) Connect() {
	var clientOptions *options.ClientOptions
	databaseUri := os.Getenv("MONGO_URL")
	if databaseUri == "" {
		log.Printf("Attempting to connect to local MongoDB instance")
		clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/todo-list-api")
	} else {
		log.Printf("Connecting to MongoDB with URI => %s", databaseUri)
		clientOptions = options.Client().ApplyURI(databaseUri)
	}

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	e := client.Ping(context.Background(), nil)

	if e != nil {
		log.Fatal(e)
	}

	log.Println("Connected to MongoDB!")

	TaskCollection = client.Database("music-api").Collection("tasks")
	UserCollection = client.Database("music-api").Collection("users")

	Database = &Instance{Client: client}

	go migrateData(TaskCollection, "task")
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
	err := gofakeit.Seed(time.Now().UnixNano())
	if err != nil {
		return nil
	}

	switch entity {
	case "task":
		return models.Task{
			Title:       gofakeit.Sentence(2),
			Description: gofakeit.Sentence(4),
			Finished:    false,
			EndDate:     time.Now(),
			User:        models.User{},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
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
			Enabled:   true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	default:
		log.Println("Invalid entity")
		return nil
	}
}
