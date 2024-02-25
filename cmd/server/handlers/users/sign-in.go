package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/authentication"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/utils/api"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// @Summary SignIn the user and returns a JWT token
// @Tags users
// @Accept  application/json
// @Produce  application/json
// @Param user body authentication.Credentials true "SignUp request"
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Router /api/v1/users/signin [post]
func SignIn(res http.ResponseWriter, req *http.Request) {
	var creds authentication.Credentials
	var response api.Response

	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		api.Error(res, req, "Malformed request", err, http.StatusBadRequest)
		return
	}

	foundUser, err := repositories.FindUserByEmail(creds.Username)
	if err != nil {
		api.Error(res, req, "Invalid email or password provided", err, http.StatusBadRequest)
		return
	} else if !foundUser.Enabled {
		api.Error(res, req, "You must verify your account first", err, http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(creds.Password))
	if err != nil {
		api.Error(res, req, "Wrong password provided", err, http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &authentication.Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(authentication.JwtKey)
	if err != nil {
		api.Error(res, req, "Invalid request", err, http.StatusBadRequest)
		return
	}

	response = api.Response{
		Message: "Login successful",
		Data:    map[string]any{"token": tokenString, "expires": expirationTime},
	}
	api.JSON(res, response, http.StatusOK)
}
