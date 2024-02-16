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

	value, _ := rc.ZRange(context.Background(), "friendrequest_"+headerUserID, 0, -1).Result()
	if err != nil {
		log.Fatal(w, "friend request not found err:", err)
		return
	}

	if value == nil {
		responseFail(w, "you don't have a friend request")
		return
	}

	if acceptReject.Status == "accept" {
		for _, data := range value {

			if data == strID {

				z := &redis.Z{
					Member: data,
				}
				r := &redis.Z{
					Member: headerUserID,
				}

				//player_1 = ecea5a261d699bd5 // player_2 = a2cf6530a027426c // player_3 = f4a028f08cac98d0  // player_4 = c051f0f740018cbd
				rc.ZAdd(context.Background(), "friend_"+headerUserID, *z).Result()
				rc.ZAdd(context.Background(), "friend_"+strID, *r).Result()

				rc.ZRem(context.Background(), "friendrequest_"+headerUserID, strID)
				rc.ZRem(context.Background(), "friendrequest_"+strID, headerUserID)
			}
		}
		responseSuccessMessage(w, "friend request accepted")
	}

	if acceptReject.Status == "reject" {

		rc.ZRem(context.Background(), "friendrequest_"+headerUserID, strID)

		responseFail(w, "friend request rejected")
	}
}
