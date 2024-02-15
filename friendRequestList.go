package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func friendRequestList(w http.ResponseWriter, r *http.Request) {
	var pageAndCount *LeaderBoard
	headerID := r.Header.Get("userID")

	err := json.NewDecoder(r.Body).Decode(&pageAndCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	firstCount := (pageAndCount.Page - 1) * pageAndCount.Count
	lastCount := (firstCount + pageAndCount.Count - 1)

	if firstCount < 0 {
		responseError(w, "Invalid page number")
		return
	}

	friendRequestList, err := rc.ZRange(context.Background(), "friendrequest_"+headerID, int64(firstCount), int64(lastCount)).Result()
	if err != nil {
		log.Fatal("ERR list friend request list", err)
		return
	}

	friendRequestSlice := make([]FriendRequestList, len(friendRequestList))

	for i, ID := range friendRequestList {
		data, _ := rc.Get(context.Background(), "user:"+ID).Result()

		intID, _ := strconv.Atoi(ID)
		dateTime := time.Now()

		friendRequestSlice[i] = FriendRequestList{
			ID:       intID,
			UserName: data,
			Date:     dateTime,
		}
	}
	responseSuccess(w, friendRequestSlice)
}
