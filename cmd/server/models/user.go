package models

import (
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id"         bson:"_id,omitempty"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Enabled   bool               `json:"enabled"`
	PhotoUrl  string             `json:"photo_url"                       binding:"required"`
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

func (u *UserRequest) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if len(u.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}

	if _, err := regexp.Match(`/^[a-z0-9.]+@[a-z0-9]+\.[a-z]+\.([a-z]+)?$/i`, []byte(u.Email)); err != nil {
		return errors.New("invalid email address")
	}

	return nil
}
