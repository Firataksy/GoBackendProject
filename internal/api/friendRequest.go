package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func (rc *RedisClient) FriendRequest(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userid")

	headerUserID := r.Header.Get("userID")

	userControl, err := rc.Client.Get(context.Background(), "user:"+userID).Result()
	if err != nil {
		log.Fatal("friendRequest user control err :", err)
		return
	}

	if userControl == "" {
		ResponseFail(w, "User not found")
		return
	}

	friendControl, err := rc.Client.ZScore(context.Background(), "friend_"+userID, headerUserID).Result()
	if err != nil {
		log.Fatal("friendRequest friend control err :", err)
		return
	}

	if friendControl == 1 {
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
