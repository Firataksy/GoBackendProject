package main

import (
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func friendRequest(w http.ResponseWriter, r *http.Request) {
	IDUrl := r.URL.Query().Get("userid")

	headerID := r.Header.Get("userID")

	userControl, _ := rc.Get(context.Background(), "user:"+IDUrl).Result()

	if userControl == "" {
		responseError(w, "User not found")
		return
	}

	if IDUrl == headerID {
		responseError(w, "You cannot send yourself a friend request.")
		return
	}

	z := &redis.Z{
		Member: headerID,
	}

	rc.ZAdd(context.Background(), "friendrequest_"+IDUrl, *z)
	responseSuccess(w, "request sent successfully")
}
