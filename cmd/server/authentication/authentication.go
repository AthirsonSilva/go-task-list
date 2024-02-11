package authentication

import (
	"errors"

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
