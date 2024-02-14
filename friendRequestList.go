package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func friendRequestList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("a")
	var requestList LeaderBoard
	headerID := r.Header.Get("userID")

	er := json.NewDecoder(r.Body).Decode(&requestList)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	firstCount := (requestList.Page - 1) * requestList.Count
	lastCount := (firstCount + requestList.Count - 1)

	friendRequestList, err := rc.ZRange(context.Background(), "friendrequest_"+headerID, int64(firstCount), int64(lastCount)).Result()
	if err != nil {
		log.Fatal("ERR list leaderboard", err)
		return
	}
	fmt.Println("requestList :", friendRequestList)

	friendRequestSlice := make([]FriendRequestList, len(friendRequestList))

	for i, ID := range friendRequestList {
		data := rc.Get(context.Background(), "userID:"+ID)
		if err != nil {
			log.Fatal("Unmarshal err :", err)
		}
		intID, _ := strconv.Atoi(ID)

		friendRequestSlice[i] = FriendRequestList{
			ID:       intID,
			UserName: data.String(),
		}
	}
	responseSuccess(w, friendRequestSlice)
}
