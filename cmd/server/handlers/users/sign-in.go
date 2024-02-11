package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/authentication"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/utils/api"
	"github.com/golang-jwt/jwt/v5"
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

	// Get the JSON body and decode into credentials
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		response = api.Response{
			Message: "Invalid request",
			Data:    nil,
		}
		api.JSON(res, response, http.StatusUnauthorized)
		return
	}

	// Get the expected password from our in memory map
	foundUser, err := repositories.FindUserByUsername(creds.Username)
	if err != nil {
		response = api.Response{
			Message: "Invalid request",
			Data:    nil,
		}
		api.JSON(res, response, http.StatusUnauthorized)
		return
	}

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if foundUser.Password != creds.Password {
		response = api.Response{
			Message: "Invalid request",
			Data:    nil,
		}
		api.JSON(res, response, http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &authentication.Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(authentication.JwtKey)
	if err != nil {
		response = api.Response{
			Message: "Invalid request",
			Data:    nil,
		}
		api.JSON(res, response, http.StatusUnauthorized)
		return
	}

	// Finally, we return a string for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	response = api.Response{
		Message: "Login successful",
		Data:    map[string]any{"token": tokenString, "expires": expirationTime},
	}
	api.JSON(res, response, http.StatusOK)
}
