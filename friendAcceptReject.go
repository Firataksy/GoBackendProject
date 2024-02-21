package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func friendAcceptReject(w http.ResponseWriter, r *http.Request) {
	var acceptReject AcceptReject
	headerUserID := r.Header.Get("userID")
	err := json.NewDecoder(r.Body).Decode(&acceptReject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if acceptReject.Status != "accept" && acceptReject.Status != "reject" {
		responseFail(w, "please write accept or reject")
		return
	}

	strID := strconv.Itoa(acceptReject.ID)

	if headerUserID == strID {
		responseFail(w, "you can not accept yourself request")
		return
	}

	value, _ := rc.ZScore(context.Background(), "friendrequest_"+headerUserID, strID).Result()
	if err != nil {
		log.Fatal(w, "friend request found err:", err)
		return
	}

	if acceptReject.Status == "accept" && value != 0.0 {

		rc.ZAdd(context.Background(), "friend_"+headerUserID, redis.Z{
			Member: strID,
		}).Result()

		rc.ZAdd(context.Background(), "friend_"+strID, redis.Z{
			Member: headerUserID,
		}).Result()

		rc.ZRem(context.Background(), "friendrequest_"+headerUserID, strID)
		rc.ZRem(context.Background(), "friendrequest_"+strID, headerUserID)

		responseSuccessMessage(w, "friend request accepted")
		return
	}

	if acceptReject.Status == "reject" && value != 0.0 {

		rc.ZRem(context.Background(), "friendrequest_"+headerUserID, strID)

		responseFail(w, "friend request rejected")
		return
	}

	responseFail(w, "request not found")
}
