package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models/dto"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/utils/api"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Creates an user
// @Tags users
// @Accept  application/json
// @Produce  application/json
// @Param user body models.UserRequest true "User request"
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Exception
// @Failure 400 {object} api.Exception
// @Failure 429 {object} api.Exception
// @Router /api/v1/users/signup [post]
func SignUp(res http.ResponseWriter, req *http.Request) {
	var request models.UserRequest
	var response api.Response

	if err := api.ReadBody(req, &request); err != nil {
		api.Error(res, req, "Malformed request", err, http.StatusBadRequest)
		return
	}

	if err := request.Validate(); err != nil {
		api.Error(res, req, "Validation error", err, http.StatusBadRequest)
		return
	}

	foundUser, err := repositories.FindUserByEmail(request.Email)
	if err == nil && foundUser.Email != "" {
		api.Error(res, req, "Given email is already in use", err, http.StatusBadRequest)
		return
	}

	user := request.ToModel()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		api.Error(res, req, "Error while creating user'", err, http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	user, err = repositories.CreateUser(user)
	if err != nil {
		api.Error(res, req, "Error while creating user", err, http.StatusInternalServerError)
		return
	}

	var prefix string
	if os.Getenv("ENV") == "production" {
		prefix = "https"
	} else {
		prefix = "http"
	}

	verificationLink := fmt.Sprintf("%s://%s/api/v1/users/verify?token=%s", prefix, req.Host, user.ID.Hex())
	log.Printf("Sending email to address: %s and verification link: %s", user.Email, verificationLink)
	body := fmt.Sprintf(`
		<h1>Email verification</h1>
		<p>Hello %s, please verify your email by clicking on the link below</p>
		<br>
		<a href="%s">Click here</a>
		<br>
		<p>Thanks!</p>
	`, user.Username, verificationLink)
	var emailData = dto.EmailData{
		To:      user.Email,
		Subject: "Email verification",
		Body:    body,
	}
	EmailChannel <- emailData

	response = api.Response{
		Message: "User registered successfully! Check your email to verify your account",
		Data:    user,
	}
	api.JSON(res, response, http.StatusCreated)
}
