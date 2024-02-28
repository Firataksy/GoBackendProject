package api

import (
	"context"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func (rc *RedisClient) FriendRequest(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userid")

	headerUserID := r.Header.Get("userID")

	_, err := rc.Client.Get(context.Background(), "user:"+userID).Result()
	if err != nil {
		ResponseFail(w, "user not found")
		return
	}

	friendControl, _ := rc.Client.ZScore(context.Background(), "friend_"+userID, headerUserID).Result()

	if friendControl != 0 {
		ResponseFail(w, "you are already friends")
		return
	}

	if userID == headerUserID {
		ResponseFail(w, "You cannot send yourself a friend request.")
		return
	}

	date := time.Now()
	unixDate := int(date.Unix())
	z := &redis.Z{
		Score:  float64(unixDate),
		Member: headerUserID,
	}

	rc.Client.ZAdd(context.Background(), "friendrequest_"+userID, *z)
	ResponseSuccessMessage(w, "request sent successfully")
}
