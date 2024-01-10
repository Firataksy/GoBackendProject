package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func registerUser(w http.ResponseWriter) *Sign {

	id := idCreate()
	strID := strconv.Itoa(int(id))
	sn := &Sign{
		ID:       int(id),
		UserName: "player_" + strID,
		Password: "12345",
		Name:     RandStringRunes(5),
		SurName:  RandStringRunes(5),
	}
	sm := SuccessMessage{
		ID:       sn.ID,
		UserName: sn.UserName,
	}
	responseSuccess(w, sm)
	return sn
}

func win(w http.ResponseWriter, user *Sign) {
	strID := strconv.Itoa(user.ID)
	user.Score += 3
	users := jsonConvert(w, user)
	_, user1WinError := rc.Set(context.Background(), "user:player_"+strID, users, 0).Result()
	if user1WinError != nil {
		log.Fatal("User1 win set error", user1WinError)
		return
	}

	z := &redis.Z{
		Score:  float64(user.Score),
		Member: user.ID,
	}

	rc.ZAdd(context.Background(), "leaderboard", *z).Result()
}

func draw(w http.ResponseWriter, user1 *Sign, user2 *Sign) {
	user1.Score += 1
	user2.Score += 1
	str1ID := strconv.Itoa(user1.ID)
	users1 := jsonConvert(w, user1)
	_, user1WinError := rc.Set(context.Background(), "user:player_"+str1ID, users1, 0).Result()
	if user1WinError != nil {
		log.Fatal("User1 win set error", user1WinError)
		return
	}

	z := &redis.Z{
		Score:  float64(user1.Score),
		Member: user1.ID,
	}

	rc.ZAdd(context.Background(), "leaderboard", *z).Result()

	str2ID := strconv.Itoa(user2.ID)
	users2 := jsonConvert(w, user2)
	_, user2WinError := rc.Set(context.Background(), "user:player_"+str2ID, users2, 0).Result()
	if user2WinError != nil {
		log.Fatal("User1 win set error", user2WinError)
		return
	}
	rz := &redis.Z{
		Score:  float64(user2.Score),
		Member: user2.ID,
	}

	rc.ZAdd(context.Background(), "leaderboard", *rz).Result()
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

		_, er := rc.Set(context.Background(), "userID:"+ru.UserName, ru.ID, 0).Result()
		if er != nil {
			log.Fatal("Set User ID err: ", er)
		}

		id := strconv.Itoa(ru.ID)
		hashPwd := md5Encode(ru.Password)
		ru.Password = hashPwd
		ruJson := jsonConvert(w, ru)
		_, error := rc.Set(context.Background(), "user:player_"+id, ruJson, 0).Result()
		if error != nil {
			log.Fatal("Set User Data err: ", error)
		}
	}
	autoMatch(w, users)
}
