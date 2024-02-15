package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func friendRequestList(w http.ResponseWriter, r *http.Request) {
	var pageAndCount *LeaderBoard
	ID := r.Header.Get("userID")

	err := json.NewDecoder(r.Body).Decode(&pageAndCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	firstCount := (pageAndCount.Page - 1) * pageAndCount.Count
	lastCount := (firstCount + pageAndCount.Count - 1)

	if firstCount < 0 {
		responseFail(w, "Invalid page number")
		return
	}

	friendRequestList, err := rc.ZRangeWithScores(context.Background(), "friendrequest_"+ID, int64(firstCount), int64(lastCount)).Result()
	if err != nil {
		log.Fatal("ERR list friend request list", err)
		return
	}

	if len(friendRequestList) == 0 {
		responseFail(w, "you don't have a request")
		return
	}

	friendRequestSlice := make([]FriendRequestList, len(friendRequestList))

	for i, requestList := range friendRequestList {
		data, _ := rc.Get(context.Background(), "user:"+requestList.Member.(string)).Result()

		intID, _ := strconv.Atoi(requestList.Member.(string))

		friendRequestSlice[i] = FriendRequestList{
			ID:       intID,
			UserName: data,
			Date:     requestList.Score,
		}
	}
	responseSuccess(w, friendRequestSlice)
}
