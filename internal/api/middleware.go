package api

import (
	"context"
	"net/http"
)

func (rc *RedisClient) tokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			responseFail(w, "Token cannot be empty")
			return
		}

		idToken, err := rc.Client.Get(context.Background(), "token:"+token).Result()
		if err != nil {
			responseFail(w, "Invalid Token")
			return
		}

		if idToken == "" {
			return
		}

		r.Header.Set("userID", idToken)
		next.ServeHTTP(w, r)
	})
}
