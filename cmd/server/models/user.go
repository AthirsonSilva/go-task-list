package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id"         bson:"_id,omitempty"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"id"         bson:"_id,omitempty"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (u *UserRequest) ToModel() User {
	var user User
	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password
	return user
}

func (u *User) ToResponse() UserResponse {
	var userResponse UserResponse
	userResponse.ID = u.ID
	userResponse.Username = u.Username
	userResponse.Email = u.Email
	userResponse.CreatedAt = u.CreatedAt
	userResponse.UpdatedAt = u.UpdatedAt
	return userResponse
}
