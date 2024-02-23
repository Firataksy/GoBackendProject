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

	err := json.NewDecoder(r.Body).Decode(&leaderBoard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if leaderBoard.Count < 0 || leaderBoard.Page < 0 {
		ResponseFail(w, "Invalid count or page number")
		return
	}

	firstCount := (leaderBoard.Page - 1) * leaderBoard.Count
	lastCount := (firstCount + leaderBoard.Count - 1)

	if firstCount < 0 {
		ResponseFail(w, "Invalid page number")
		return
	}

	leaderBoardList, err := rc.Client.ZRevRangeWithScores(context.Background(), "leaderboard", int64(firstCount), int64(lastCount)).Result()

	if err != nil {
		log.Fatal("ERR list leaderboard", err)
		return
	}

	leaderBoardSlice := make([]UserLeaderBoard, len(leaderBoardList))
	for i, data := range leaderBoardList {

		s, err := rc.Client.Get(context.Background(), "user:"+data.Member.(string)).Result()
		if err != nil {
			log.Fatal("leaderboard get username err :", err)
			return
		}

		data, err := rc.Client.Get(context.Background(), s).Result()
		if err != nil {
			log.Fatal("leaderboard get user data err :", err)
			return
		}

		err = json.Unmarshal([]byte(data), &userData)
		if err != nil {
			log.Fatal("leaderboard unmarshal err:", err)
			return
		}

		leaderBoardSlice[i] = UserLeaderBoard{
			UserID:   userData.UserID,
			Score:    userData.Score,
			UserName: userData.UserName,
		}
	}

	ResponseSuccess(w, leaderBoardSlice)
}
