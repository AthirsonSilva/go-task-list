package repositories

import (
	"context"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"log"
	"time"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/database"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setFindOptions(options *options.FindOptions, pagination api.Pagination) {
	options.SetLimit(int64(pagination.PageSize))
	if pagination.PageNumber == 0 {
		options.SetSkip(0)
	} else {
		options.SetSkip(int64((pagination.PageNumber - 1) * pagination.PageSize))
	}
	options.SetSort(bson.D{{Key: pagination.SortField, Value: pagination.SortDirection}})
}

func FindAllTasks(pagination api.Pagination, userID primitive.ObjectID) ([]models.Task, error) {
	var tasks []models.Task
	var findOptions = options.Find()
	setFindOptions(findOptions, pagination)

	searchParams := bson.M{}
	if pagination.SearchName != "" {
		searchParams["title"] = bson.M{"$regex": pagination.SearchName, "$findOptions": "i"}
	}

	if userID != primitive.NilObjectID {
		searchParams["user._id"] = userID
	}

	cursor, err := database.TaskCollection.Find(context.Background(), searchParams, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.Background()) {
		var task models.Task

		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func FindTaskById(id string) (models.Task, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	var task models.Task

	if err != nil {
		return task, err
	}

	err = database.TaskCollection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&task)
	if err != nil {
		return task, err
	}

	err = database.TaskCollection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&task)
	if err != nil {
		return task, err
	}

	return task, nil
}

func CreateTask(task models.Task) (models.Task, error) {
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := database.TaskCollection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}

	return task, nil
}

func UpdateTaskById(id string, task models.Task) (models.Task, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, err
	}

	err = database.TaskCollection.FindOne(context.Background(), bson.M{"_id": oid}).Err()
	if err != nil {
		return task, err
	}

	task.UpdatedAt = time.Now()
	task.ID = oid

	_, err = database.TaskCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": oid},
		bson.D{{Key: "$set", Value: task}},
	)
	if err != nil {
		return task, err
	}

	return task, nil
}

func DeleteTaskById(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = database.TaskCollection.FindOne(context.Background(), bson.M{"_id": oid}).Err()
	if err != nil {
		return err
	}

	_, err = database.TaskCollection.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return err
	}

	return nil
}
