package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

func registerUser(w http.ResponseWriter) *Sign {

	id := idCreate()
	strID := strconv.Itoa(int(id))
	sn := &Sign{
		ID:       int(id),
		UserName: "player_" + strID,
		Password: "12345",
		Name:     RandStringRunes(5),
		SurName:  RandStringRunes(6),
	}
	sm := SuccessMessage{
		ID:       sn.ID,
		UserName: sn.UserName,
	}
	responseSuccess(w, sm)
	return sn
}

func win(w http.ResponseWriter, user *Sign) {
	user.Score += 3

	redisSetJustData(w, user.ID, user)
	redisSetLeaderBoard(user)

}

func draw(w http.ResponseWriter, user1 *Sign, user2 *Sign) {
	user1.Score += 1
	user2.Score += 1

	redisSetJustData(w, user1.ID, user1)
	redisSetLeaderBoard(user1)

	redisSetJustData(w, user2.ID, user2)
	redisSetLeaderBoard(user2)
}

func autoMatch(w http.ResponseWriter, users []*Sign) {
	var match Match
	for i := 0; i < len(users); i++ {
		user1 := users[i]
		for j := i; j < len(users); j++ {
			user2 := users[j]
			if i == j {
				continue
			}
			match.Score1 = rand.Intn(5)
			match.Score2 = rand.Intn(5)
			if match.Score1 > match.Score2 {
				win(w, user1)
			}
			if match.Score1 < match.Score2 {
				win(w, user2)
			}
			if match.Score1 == match.Score2 {
				draw(w, user1, user2)
			}
		}
	}
}

func simulation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sim Simulation
	err := json.NewDecoder(r.Body).Decode(&sim)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users := make([]*Sign, sim.Count)
	for i := 0; i < sim.Count; i++ {
		ru := registerUser(w)
		users[i] = ru
		hashPwd := md5Encode(ru.Password)
		ru.Password = hashPwd

		redisSetDataAndID(w, ru.UserName, ru.ID, ru)
	}
	autoMatch(w, users)
}
