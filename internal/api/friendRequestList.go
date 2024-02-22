package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (rc *RedisClient) FriendRequestList(w http.ResponseWriter, r *http.Request) {
	var pageAndCount *LeaderBoard
	headerUserID := r.Header.Get("userID")

	err := json.NewDecoder(r.Body).Decode(&pageAndCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	firstCount := (pageAndCount.Page - 1) * pageAndCount.Count
	lastCount := (firstCount + pageAndCount.Count - 1)

	if firstCount < 0 {
		ResponseFail(w, "Invalid page number")
		return
	}

	friendRequestList, err := rc.Client.ZRangeWithScores(context.Background(), "friendrequest_"+headerUserID, int64(firstCount), int64(lastCount)).Result()
	if err != nil {
		log.Fatal("ERR list friend request list", err)
		return
	}

	if len(friendRequestList) == 0 {
		ResponseFail(w, "you don't have a request")
		return
	}

	friendRequestSlice := make([]FriendRequestList, len(friendRequestList))

	for i, requestList := range friendRequestList {
		data, _ := rc.Client.Get(context.Background(), "user:"+requestList.Member.(string)).Result()

		intID, _ := strconv.Atoi(requestList.Member.(string))

		friendRequestSlice[i] = FriendRequestList{
			ID:       intID,
			UserName: data,
			Date:     requestList.Score,
		}
	}
	ResponseSuccess(w, friendRequestSlice)
}
