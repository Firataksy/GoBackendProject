package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func listLeaderBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var leaderboard LeaderBoard

	er := json.NewDecoder(r.Body).Decode(&leaderboard)
	if er != nil {
		log.Fatal("ERR", er)
		return
	}

	leaderboardlist, err := rc.ZRevRangeWithScores(context.Background(), "leaderboard", 0, -1).Result()
	if err != nil {
		log.Fatal("ERR list leaderboard", err)
		return
	}

	startIndex := (leaderboard.Page - 1) * leaderboard.Count
	endIndex := startIndex + leaderboard.Count

	if startIndex < 0 || startIndex >= len(leaderboardlist) {
		responseError(w, "Invalid page number")
		return
	}

	if endIndex > len(leaderboardlist) {
		endIndex = len(leaderboardlist)
	}

	list := leaderboardlist[startIndex:endIndex]
	responseSuccess(w, list)
}
