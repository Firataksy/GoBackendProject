package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (rc *RedisClient) FriendList(w http.ResponseWriter, r *http.Request) {
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

	friendList, err := rc.Client.ZRangeWithScores(context.Background(), "friend_"+headerUserID, int64(firstCount), int64(lastCount)).Result()
	if err != nil {
		log.Fatal("ERR list friend request list", err)
		return
	}

	if len(friendList) == 0 {
		ResponseFail(w, "you don't have any friends, how about adding a friend?")
		return
	}

	friendSlice := make([]FriendRequestList, len(friendList))

	for i, data := range friendList {
		userName, err := rc.Client.Get(context.Background(), "user:"+data.Member.(string)).Result()
		if err != nil {
			log.Fatal("friendlist get username err :", err)
		}

		intID, _ := strconv.Atoi(data.Member.(string))

		friendSlice[i] = FriendRequestList{
			ID:       intID,
			UserName: userName,
			Date:     data.Score,
		}
	}
	ResponseSuccess(w, friendSlice)
}
