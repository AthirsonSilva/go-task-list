package authentication

import (
	"errors"
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

func GetTokenInfo(tokenString string) (Claims, error) {
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
