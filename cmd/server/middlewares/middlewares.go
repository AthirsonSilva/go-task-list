package middlewares

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/authentication"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/utils/api"
	"github.com/golang-jwt/jwt/v5"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Request method => %s", req.Method)
		log.Printf("Request protocol => %s", req.Proto)

		headers := []string{"Content-Type", "Accept", "User-Agent"}
		for _, h := range headers {
			headerValue := req.Header.Get(h)
			if headerValue == "" {
				log.Printf("Request header => %s: <empty>", h)
			} else {
				log.Printf("Request header => %s: %s", h, headerValue)
			}
		}

		next.ServeHTTP(res, req)
	})
}

func VerifyAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		rawToken, err := api.AuthToken(req)
		if err != nil {
			response := api.Response{
				Message: "Unauthorized",
				Data:    nil,
			}
			api.JSON(res, response, http.StatusUnauthorized)
			return
		}

		claims := authentication.Claims{}
		token, err := jwt.ParseWithClaims(rawToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return authentication.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				response := api.Response{
					Message: "Invalid JWT token signature",
					Data:    nil,
				}
				api.JSON(res, response, http.StatusBadRequest)
				return
			} else if !token.Valid {
				response := api.Response{
					Message: "Invalid JWT token",
					Data:    nil,
				}
				api.JSON(res, response, http.StatusUnauthorized)
				return
			}
		}

		log.Printf("User logged: %s", claims.Username)
		next.ServeHTTP(res, req)
	})
}
