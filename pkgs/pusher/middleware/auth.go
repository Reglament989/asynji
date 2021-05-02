package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	TestApiKey := "test_api_key"
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var apiKey string
		if apiKey = req.Header.Get("X-Api-Key"); apiKey != TestApiKey {
			log.Printf("bad auth api key: %s", apiKey)
			rw.WriteHeader(http.StatusForbidden)
			return
		}
		// ctx, cancel := context.WithTimeout(r.Context(), time.Duration(60*time.Second))
		// defer cancel()
		// req = req.WithContext(ctx)
		// ctx := context.WithValue(r.Context(), "fcm-", "theuser")
		next.ServeHTTP(rw, req)
	})
}
