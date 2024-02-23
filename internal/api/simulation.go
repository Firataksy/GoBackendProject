package api

import (
	"context"
	"encoding/json"
	"log"

	"math/rand"
	"net/http"
)

func (rc *RedisClient) redisGetAllUser() []*Sign {
	var user *Sign

	users, err := rc.Client.ZRange(context.Background(), "leaderboard", 0, -1).Result()
	if err != nil {
		log.Fatal("Redis Could Not Get User ID", err)
	}

	allUser := make([]*Sign, len(users))
	for i, data := range users {
		userName, err := rc.Client.Get(context.Background(), "user:"+data).Result()
		if err != nil {
			log.Fatal("simulation get username err :", err)
			return nil
		}

		data, err := rc.Client.Get(context.Background(), userName).Result()
		if err != nil {
			log.Fatal("simulation get user data err :", err)
			return nil
		}

		err = json.Unmarshal([]byte(data), &user)
		if err != nil {
			log.Fatal("simulation unmarshal err: ", err)
		}

		allUser[i] = &Sign{
			Token:    user.Token,
			ID:       user.ID,
			UserName: user.UserName,
			Password: user.Password,
			Name:     user.Name,
			SurName:  user.SurName,
			Score:    user.Score,
		}
	}

	return allUser
}

func (rc *RedisClient) autoMatch(w http.ResponseWriter, users []*Sign) {
	var match Match

	for i := 0; i < len(users); i++ {
		user1 := users[i]
		for j := i; j < len(users); j++ {
			user2 := users[j]
			if i == j {
				continue
			} else {
				match.Score1 = rand.Intn(5)
				match.Score2 = rand.Intn(5)
				if match.Score1 > match.Score2 {
					rc.win(user1)
				}
				if match.Score1 < match.Score2 {
					rc.win(user2)
				}
				if match.Score1 == match.Score2 {
					rc.draw(user1, user2)
				}
			}
		}
	}
}

func (rc *RedisClient) Simulation(w http.ResponseWriter, r *http.Request) {
	var sim Simulation
	err := json.NewDecoder(r.Body).Decode(&sim)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users := make([]*Sign, sim.Count)
	for i := 0; i < len(users); i++ {
		ru := rc.RegisterUser()

		users[i] = ru

		rc.redisSetDataAndID(ru)
		defer rc.redisSetLeaderBoard(ru)
	}

	redisData := rc.redisGetAllUser()
	allUsers := append(users, redisData...)
	ResponseSuccessMessage(w, "")
	rc.autoMatch(w, allUsers)
}
