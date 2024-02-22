package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

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

	value, _ := rc.Client.ZScore(context.Background(), "friendrequest_"+headerUserID, strID).Result()

	if acceptReject.Status == "accept" && value != 0.0 {

		rc.Client.ZAdd(context.Background(), "friend_"+headerUserID, redis.Z{
			Member: strID,
			Score:  1,
		}).Result()

		rc.Client.ZAdd(context.Background(), "friend_"+strID, redis.Z{
			Member: headerUserID,
			Score:  1,
		}).Result()

		rc.Client.ZRem(context.Background(), "friendrequest_"+headerUserID, strID)
		rc.Client.ZRem(context.Background(), "friendrequest_"+strID, headerUserID)

		ResponseSuccessMessage(w, "friend request accepted")
		return
	}

	if acceptReject.Status == "reject" && value != 0.0 {

		rc.Client.ZRem(context.Background(), "friendrequest_"+headerUserID, strID)

		ResponseFail(w, "friend request rejected")
		return
	}

	ResponseFail(w, "request not found")
}
