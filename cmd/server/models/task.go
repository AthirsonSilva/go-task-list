package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"id"         bson:"_id,omitempty"`
	Title       string             `json:"title"                           binding:"required"`
	Description string             `json:"description"`
	User        User               `json:"user"`
	EndDate     time.Time          `json:"end_date"`
	Finished    bool               `json:"finished"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type TaskRequest struct {
	Title       string `json:"title"  binding:"required"`
	Description string `json:"description"`
	EndDate     string `json:"end_date"`
}

type TaskResponse struct {
	ID        primitive.ObjectID `json:"id"         bson:"_id,omitempty"`
	Title     string             `json:"title"                           binding:"required"`
	EndDate   time.Time          `json:"end_date"`
	Finished  bool               `json:"finished"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (a *TaskRequest) ToModel() Task {
	var task Task
	task.Title = a.Title
	task.Description = a.Description
	endDate, err := time.Parse("2006-01-02", a.EndDate)
	if err != nil {
		task.EndDate = time.Now()
	} else {
		task.EndDate = endDate
	}
	return task
}

func (a *Task) ToResponse() TaskResponse {
	var task TaskResponse
	task.ID = a.ID
	task.Title = a.Title
	task.EndDate = a.EndDate
	task.Finished = a.Finished
	task.CreatedAt = a.CreatedAt
	task.UpdatedAt = a.UpdatedAt
	return task
}
