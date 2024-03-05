package authentication

import (
	"errors"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("go_is_awesome")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (c *Credentials) Valid() error {
	if c.Username == "" {
		return errors.New("missing username")
	}
	if c.Password == "" {
		return errors.New("missing password")
	}
	return nil
}

func GetTokenInfo(tokenString string) (jwtClaims Claims, err error) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return claims, err
		} else if !token.Valid {
			return claims, err
		} else if time.Until(claims.ExpiresAt.Time) < 0 {
			return claims, err
		}
	}
	return claims, nil
}

func GetUserFromToken(req *http.Request) (user models.User, err error) {
	token, err := api.AuthToken(req)
	if err != nil {
		return user, err
	}
	logger.Info("GetUserFromToken", "token: "+token)

	claims, err := GetTokenInfo(token)
	if err != nil {
		return user, err
	}
	logger.Info("GetUserFromToken", "username: "+claims.Username)

	user, err = repositories.FindUserByEmail(claims.Username)
	if err != nil {
		return user, err
	}
	logger.Info("GetUserFromToken", "User found: "+user.ID.Hex())

	return user, nil
}
