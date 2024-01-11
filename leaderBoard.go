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
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}
	if leaderBoard.Count < 0 || leaderBoard.Page < 0 {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}
	firstCount := (leaderBoard.Page - 1) * leaderBoard.Count
	lastCount := (firstCount + leaderBoard.Count - 1)

	leaderBoardList, err := rc.ZRevRangeWithScores(context.Background(), "leaderboard", int64(firstCount), int64(lastCount)).Result()
	if err != nil {
		log.Fatal("ERR list leaderboard", err)
		return
	}

	if firstCount < 0 {
		responseError(w, "Invalid page number")
		return
	}

	leaderBoardSlice := make([]UserLeaderBoard, len(leaderBoardList))
	for i, data := range leaderBoardList {
		s, _ := rc.Get(context.Background(), "player_"+data.Member.(string)).Result()
		err := json.Unmarshal([]byte(s), &userData)
		if err != nil {
			log.Fatal("Unmarshal err:", err)
		}
		leaderBoardSlice[i] = UserLeaderBoard{
			UserID:   userData.UserID,
			Score:    userData.Score,
			UserName: userData.UserName,
		}
	}
	responseSuccess(w, leaderBoardSlice)
}
