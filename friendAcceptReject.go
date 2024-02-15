package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func friendAcceptReject(w http.ResponseWriter, r *http.Request) {
	var acceptReject AcceptReject
	headerID := r.Header.Get("userID")
	err := json.NewDecoder(r.Body).Decode(&acceptReject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value, _ := rc.ZRange(context.Background(), "friendrequest_"+headerID, 0, -1).Result()

	if acceptReject.Status == "accept" {
		for _, data := range value {
			fmt.Println(data)
			z := &redis.Z{
				Member: data,
			}

			_, err := rc.ZAdd(context.Background(), "friend_"+headerID, *z).Result()
			if err != nil {
				log.Fatal("Friend add err :", err)
				return
			}

			rc.ZRem(context.Background(), "friendrequest_"+headerID, data)
		}
		responseSuccess(w, "friend added")
	}

	if acceptReject.Status == "reject" {
		rc.ZRem(context.Background(), "friendrequest_"+headerID)
		responseError(w, "friend request rejected")
	}
}
