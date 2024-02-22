package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (rc *RedisClient) ListLeaderBoard(w http.ResponseWriter, r *http.Request) {
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

	if firstCount < 0 {
		responseFail(w, "Invalid page number")
		return
	}

	leaderBoardList, err := rc.Client.ZRevRangeWithScores(context.Background(), "leaderboard", int64(firstCount), int64(lastCount)).Result()
	if err != nil {
		log.Fatal("ERR list leaderboard", err)
		return
	}

	leaderBoardSlice := make([]UserLeaderBoard, len(leaderBoardList))
	for i, data := range leaderBoardList {
		s, _ := rc.Client.Get(context.Background(), "user:"+data.Member.(string)).Result()
		data, _ := rc.Client.Get(context.Background(), s).Result()

		err := json.Unmarshal([]byte(data), &userData)
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
