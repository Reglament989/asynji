package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/Reglament989/asynji/pkgs/asynji/encryption"
)

type ContextKey string

const ContextUserKey ContextKey = "userId"

type WsError struct {
	Err string
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		apiKey := req.Header.Get("X-Api-Key")
		userId, err := encryption.VerifyToken(apiKey)
		log.Printf("Userid - %s", userId)
		if err != nil {
			log.Println(err)
			// response, _ := json.Marshal(&WsError{Err: err.Error()})
			// rw.Write(response)
			rw.WriteHeader(http.StatusForbidden)
			return
		}
		ctx := context.WithValue(req.Context(), ContextUserKey, userId)
		req = req.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
