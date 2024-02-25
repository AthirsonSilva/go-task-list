package handlers

import (
	"fmt"
	"log"
	"net/http"

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
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Router /api/v1/users/signup [post]
func SignUp(res http.ResponseWriter, req *http.Request) {
	var request models.UserRequest
	var response api.Response

	if err := api.ReadBody(req, &request); err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusBadRequest)
		return
	}

	if err := request.Validate(); err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusBadRequest)
		return
	}

	foundUser, err := repositories.FindUserByEmail(request.Email)
	if err == nil && foundUser.Email != "" {
		response = api.Response{
			Message: "Given email is already in use",
			Data:    request,
		}
		api.JSON(res, response, http.StatusBadRequest)
		return
	}

	user := request.ToModel()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	user, err = repositories.CreateUser(user)
	if err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusInternalServerError)
		return
	}

	log.Println("Sending email to address: ", user.Email)
	body := fmt.Sprintf(`
		<h1>Email verification</h1>
		<p>Hello %s, please verify your email by clicking on the link below</p>
		<br>
		<a href="http://localhost:8080/api/v1/users/verify?token=%s">Click here</a>
		<br>
		<p>Thanks!</p>
	`, user.Username, user.ID.Hex())
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
