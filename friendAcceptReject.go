package main

import (
	"context"
	"encoding/json"
	"fmt"
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
		log.Fatal(w, "friend request found err:", err)
		return
	}

	fmt.Println(value)
	if value == nil {
		responseFail(w, "you don't have a friend request")
		return
	}

	if acceptReject.Status == "accept" {
		for _, data := range value {
			fmt.Println("a: ", data, strID)
			if data != strID {
				responseFail(w, "friend request not found")
				return
			}

			if data == strID {

				//player_1 = 09999bfacabc6ed2 // player_2 = 21b102e34eeb4b82 // player_3 = b843843f0c6cf06e  // player_4 = 2d6213e3d94a0150
				rc.ZAdd(context.Background(), "friend_"+headerUserID, redis.Z{
					Member: data,
				}).Result()

				rc.ZAdd(context.Background(), "friend_"+strID, redis.Z{
					Member: headerUserID,
				}).Result()

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
