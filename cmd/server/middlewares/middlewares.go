package middlewares

import (
	"log"
	"net/http"
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
