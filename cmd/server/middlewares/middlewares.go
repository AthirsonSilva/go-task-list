package middlewares

import (
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/authentication"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/database"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	"golang.org/x/time/rate"
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
		logger.Info("WriteToConsole", "Request path => "+req.URL.Path)
		logger.Info("WriteToConsole", "Request method => "+req.Method)
		logger.Info("WriteToConsole", "Request protocol => "+req.Proto)

		headers := []string{"Content-Type", "Accept", "User-Agent"}
		for _, h := range headers {
			headerValue := req.Header.Get(h)
			if headerValue == "" {
				logger.Info("WriteToConsole", "header not found: "+h)
			} else {
				logger.Info("WriteToConsole", h+": "+headerValue)
			}
		}

		next.ServeHTTP(res, req)
	})
}

func VerifyAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		logger.Info("VerifyAuthentication", "Verifying authentication...")

		rawToken, err := api.AuthToken(req)
		if err != nil {
			response := api.Response{
				Message: "Unauthorized",
				Data:    nil,
			}
			api.Error(res, req, response.Message, err, http.StatusUnauthorized)
			return
		}

		claims, err := authentication.GetTokenInfo(rawToken)
		if err != nil {
			logger.Error("VerifyAuthentication", err.Error())
			api.Error(res, req, "Invalid or expired token", err, http.StatusUnauthorized)
			return
		}

		redisToken, err := database.GetRedisObj(claims.Username)
		if err != nil {
			logger.Error("VerifyAuthentication", err.Error())
			api.Error(res, req, "Invalid or expired token", err, http.StatusUnauthorized)
			return
		}

		logger.Info("VerifyAuthentication", "redis token: "+redisToken+" | raw token: "+rawToken)
		if redisToken != rawToken {
			logger.Error("VerifyAuthentication", "Invalid or expired token")
			api.Error(res, req, "Invalid or expired token", err, http.StatusUnauthorized)
			return
		}

		logger.Info("VerifyAuthentication", "username: "+claims.Username)
		next.ServeHTTP(res, req)
	})
}
