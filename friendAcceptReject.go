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
	ID := r.Header.Get("userID")
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

	value, err := rc.ZRange(context.Background(), "friendrequest_"+ID, 0, -1).Result()
	if err != nil {
		log.Fatal(w, "friend request not found err:", err)
		return
	}

	if len(value) == 0 {
		responseFail(w, "you don't have a friend request")
		return
	}

	if acceptReject.Status == "accept" {
		for _, data := range value {
			z := &redis.Z{
				Member: data,
			}
			r := &redis.Z{
				Member: ID,
			}

			if strID == data {
				rc.ZAdd(context.Background(), "friend_"+ID, *z).Result()
				rc.ZAdd(context.Background(), "friend_"+data, *r).Result()

				rc.ZRem(context.Background(), "friendrequest_"+ID, strID)
			}

		}
		responseSuccess(w, "friend request accepted")
	}

	if acceptReject.Status == "reject" {
		for _, data := range value {
			if data == strID {
				rc.ZRem(context.Background(), "friendrequest_"+ID, strID)
			}
		}
		responseFail(w, "friend request rejected")
	}
}
