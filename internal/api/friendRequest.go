package api

import (
	"context"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func (rc *RedisClient) FriendRequest(w http.ResponseWriter, r *http.Request) {
	urlUserID := r.URL.Query().Get("userid")

	headerUserID := r.Header.Get("userID")

	userControl, _ := rc.Client.Get(context.Background(), "user:"+urlUserID).Result()

	if userControl == "" {
		responseFail(w, "User not found")
		return
	}

	friendControl, _ := rc.Client.ZScore(context.Background(), "friend_"+urlUserID, headerUserID).Result()

	if friendControl == 1 {
		responseFail(w, "you are already friends")
		return
	}

	if urlUserID == headerUserID {
		responseFail(w, "You cannot send yourself a friend request.")
		return
	}

	date := time.Now()
	unixDate := int(date.Unix())
	z := &redis.Z{
		Score:  float64(unixDate),
		Member: headerUserID,
	}

	rc.Client.ZAdd(context.Background(), "friendrequest_"+urlUserID, *z)
	responseSuccessMessage(w, "request sent successfully")
}
