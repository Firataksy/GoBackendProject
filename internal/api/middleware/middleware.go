package api

import (
	"context"
	"net/http"

	"github.com/my/repo/internal/api"
)

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rc, err := api.ConnectRedis()
		if err != nil {
			return
		}
		token := r.Header.Get("token")
		if token == "" {
			api.ResponseFail(w, "Token cannot be empty")
			return
		}

		idToken, err := rc.Client.Get(context.Background(), "token:"+token).Result()
		if err != nil {
			api.ResponseFail(w, "Invalid Token")
			return
		}

		r.Header.Set("userID", idToken)
		next.ServeHTTP(w, r)
	})
}
