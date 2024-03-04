package middlewares

import (
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/authentication"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"golang.org/x/time/rate"
	"log"
	"net/http"
)

func RateLimiter(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(3, 5) // max 5 requests per second and burst of 3
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if !limiter.Allow() {
			res.WriteHeader(http.StatusTooManyRequests)
			_, err := res.Write([]byte("Too many requests"))
			if err != nil {
				return
			}
			return
		} else {
			next.ServeHTTP(res, req)
		}
	})
}

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Request method => %s", req.Method)
		log.Printf("Request protocol => %s", req.Proto)
		log.Printf("Request URL => %s", req.URL)
		log.Printf("Request HOST => %s", req.Host)

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

		claims, err := authentication.GetTokenInfo(rawToken)
		if err != nil {
			api.Error(res, req, "Invalid or expired token", err, http.StatusUnauthorized)
			return
		}

		log.Printf("User logged in: %s", claims.Username)
		next.ServeHTTP(res, req)
	})
}
