package handlers

import (
	"errors"
	"fmt"
	awsservice "github.com/AthirsonSilva/music-streaming-api/cmd/server/aws"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

// SignUp @Summary Creates a user
//
//	@Tags		users
//	@Accept		application/json
//	@Produce	application/json
//	@Param		user	formData	models.UserRequest	true	"User request"
//	@Param		file	formData	file				false	"File"
//	@Success	200		{object}	api.Response
//	@Failure	500		{object}	api.Exception
//	@Failure	400		{object}	api.Exception
//	@Failure	429		{object}	api.Exception
//	@Router		/api/v1/users/signup [post]
func SignUp(res http.ResponseWriter, req *http.Request) {
	var request models.UserRequest
	var response api.Response

	username := req.FormValue("username")
	email := req.FormValue("email")
	password := req.FormValue("password")
	if username == "" || email == "" || password == "" {
		api.Error(res, req, "Username, email and password are required", errors.New("username, email and password are required"), http.StatusBadRequest)
		return
	}

	request.Username = username
	request.Email = email
	request.Password = password
	if err := request.Validate(); err != nil {
		api.Error(res, req, "Validation error", err, http.StatusBadRequest)
		return
	}

	foundUser, err := repositories.FindUserByEmail(request.Email)
	if err == nil && foundUser.Email != "" {
		err = errors.New("email already in use")
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

	photoUrl, fileName, err := api.FileUpload(req)
	if err != nil {
		api.Error(res, req, "Error while uploading photo", err, http.StatusInternalServerError)
		return
	}

	user, err = repositories.CreateUser(user)
	if err != nil {
		api.Error(res, req, "Error while creating user", err, http.StatusInternalServerError)
		return
	}

	if photoUrl != "" {
		go func() {
			filePath := "/tmp/" + fileName
			key := "users/" + user.ID.Hex() + "-" + user.Username

			s3FilePath, err := awsservice.PutBucketObject(key, filePath)
			if err != nil {
				log.Printf("Error while uploading photo: %s", err)
			}

			user.PhotoUrl = s3FilePath
			_, err = repositories.UpdateUserByID(user)
			if err != nil {
				log.Printf("Error while uploading photo: %s", err)
			}
		}()
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
	var emailData = models.EmailDto{
		To:      user.Email,
		Subject: "Email verification",
		Body:    body,
	}
	EmailChannel <- emailData

	userResponse := user.ToResponse()
	response = api.Response{
		Message: "User registered successfully! Check your email to verify your account",
		Data:    userResponse,
	}
	api.JSON(res, response, http.StatusCreated)
}
