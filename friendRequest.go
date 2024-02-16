package main

import (
	"context"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func friendRequest(w http.ResponseWriter, r *http.Request) {
	IDUrl := r.URL.Query().Get("userid")

	ID := r.Header.Get("userID")

	userControl, _ := rc.Get(context.Background(), "user:"+IDUrl).Result()

	if userControl == "" {
		responseFail(w, "User not found")
		return
	}

	friendControl, _ := rc.ZRange(context.Background(), "friend_"+IDUrl, 0, -1).Result()

	for _, data := range friendControl {
		if data == ID {
			responseFail(w, "you are already friend")
			return
		}
	}

	if IDUrl == ID {
		responseFail(w, "You cannot send yourself a friend request.")
		return
	}

	date := time.Now()
	unixDate := int(date.Unix())
	z := &redis.Z{
		Score:  float64(unixDate),
		Member: ID,
	}

	rc.ZAdd(context.Background(), "friendrequest_"+IDUrl, *z)
	responseSuccess(w, "request sent successfully")
}
