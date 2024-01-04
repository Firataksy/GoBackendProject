package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func listLeaderBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var leaderBoard LeaderBoard
	var userData UserLeaderBoard

	er := json.NewDecoder(r.Body).Decode(&leaderBoard)
	if er != nil {
		log.Fatal("ERR", er)
		return
	}
	if leaderBoard.Count < 0 || leaderBoard.Page < 0 {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	leaderBoardList, err := rc.ZRevRangeWithScores(context.Background(), "leaderboard", 0, -1).Result()
	if err != nil {
		log.Fatal("ERR list leaderboard", err)
		return
	}

	leaderBoardSlice := make([]UserLeaderBoard, len(leaderBoardList))
	for rank, data := range leaderBoardList {
		data, _ := rc.Get(context.Background(), "user:"+data.Member.(string)).Result()
		json.Unmarshal([]byte(data), &userData)

		leaderBoardSlice[rank] = UserLeaderBoard{
			Rank:     rank + 1,
			UserID:   userData.UserID,
			Score:    userData.Score,
			UserName: userData.UserName,
		}
	}

	firstCount := (leaderBoard.Page - 1) * leaderBoard.Count
	lastCount := firstCount + leaderBoard.Count

	if firstCount < 0 || firstCount >= len(leaderBoardList) {
		responseError(w, "Invalid page number")
		return
	}

	if lastCount > len(leaderBoardList) {
		lastCount = len(leaderBoardList)
	}

	list := leaderBoardSlice[firstCount:lastCount]
	responseSuccess(w, list)
}
