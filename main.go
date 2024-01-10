package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signUp)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/data", getUserData)
	mux.HandleFunc("/updateuser", updateUserData)
	mux.HandleFunc("/match", match)
	mux.HandleFunc("/leaderboard", listLeaderBoard)
	mux.HandleFunc("/simulation", simulation)
	fmt.Println("http listen started")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		fmt.Println(err)
	}
}

var rc *redis.Client

func redisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func init() {
	rc = redisConnect()
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func idCreate() int64 {
	userIncrID, _ := rc.Incr(context.Background(), "userIncrId").Result()
	return userIncrID
}

func jsonConvert(w http.ResponseWriter, input interface{}) []byte {
	Json, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Json Error", http.StatusInternalServerError)
		return nil
	}
	return Json
}

func responseSuccess(w http.ResponseWriter, input interface{}) {
	w.Header().Set("Content-Type", "application/json")

	rp := Response{
		Status: true,
		Data:   input,
	}

	response := jsonConvert(w, rp)
	w.Write(response)
}

func responseError(w http.ResponseWriter, input string) {
	w.Header().Set("Content-Type", "application/json")
	ms := &FailMessage{
		Status:  false,
		Message: input,
	}

	response := jsonConvert(w, ms)
	w.Write(response)
}

func redisSetData(w http.ResponseWriter, id int, data interface{}) {
	strID := strconv.Itoa(id)
	_, er := rc.Set(context.Background(), "user:player_"+strID, data, 0).Result()
	if er != nil {
		log.Fatal("Set User ID err: ", er)
	}
}

func redisSetID(w http.ResponseWriter, username string, id int) {

	_, er := rc.Set(context.Background(), "userID:"+username, id, 0).Result()
	if er != nil {
		log.Fatal("Set User ID err: ", er)
	}
}
