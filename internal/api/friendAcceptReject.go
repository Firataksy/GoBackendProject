package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func (rc *RedisClient) FriendAcceptReject(w http.ResponseWriter, r *http.Request) {
	var acceptReject AcceptReject
	headerUserID := r.Header.Get("userID")
	err := json.NewDecoder(r.Body).Decode(&acceptReject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if acceptReject.Status != "accept" && acceptReject.Status != "reject" {
		ResponseFail(w, "please write accept or reject")
		return
	}

	strID := strconv.Itoa(acceptReject.ID)

	if headerUserID == strID {
		ResponseFail(w, "you can not accept yourself request")
		return
	}

	friendRequestControl, err := rc.Client.ZScore(context.Background(), "friendrequest_"+headerUserID, strID).Result()
	if err != nil {
		ResponseFail(w, "you don't have a friend request")
		return
	}

	if acceptReject.Status == "accept" && friendRequestControl != 0.0 {

		date := time.Now()
		unixDate := int(date.Unix())

		rc.Client.ZAdd(context.Background(), "friend_"+headerUserID, redis.Z{
			Member: strID,
			Score:  float64(unixDate),
		}).Result()

		rc.Client.ZAdd(context.Background(), "friend_"+strID, redis.Z{
			Member: headerUserID,
			Score:  float64(unixDate),
		}).Result()

		rc.Client.ZRem(context.Background(), "friendrequest_"+headerUserID, strID)
		rc.Client.ZRem(context.Background(), "friendrequest_"+strID, headerUserID)

		ResponseSuccessMessage(w, "friend request accepted")
		return
	}

	if acceptReject.Status == "reject" && friendRequestControl != 0.0 {

		rc.Client.ZRem(context.Background(), "friendrequest_"+headerUserID, strID)

		ResponseFail(w, "friend request rejected")
		return
	}

	ResponseFail(w, "request not found")
}
