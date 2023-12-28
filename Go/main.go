package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/data", getUserData)
	mux.HandleFunc("/updateuser", updateUserData)
	mux.HandleFunc("/match", match)
	mux.HandleFunc("/leaderboard", listLeaderBoard)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
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
	rc.FlushAll(context.Background())
}

func jsonConvert(w http.ResponseWriter, input interface{}) []byte {
	Json, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Json Error", http.StatusInternalServerError)
		return nil
	}
	return Json
}

func jsonWrite(w http.ResponseWriter, input []byte) {
	w.Write(input)
}

func responseSuccess(w http.ResponseWriter, input interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rp := Response{
		Status: true,
		Data:   input,
	}

	response := jsonConvert(w, rp)
	jsonWrite(w, response)
}

func responseError(w http.ResponseWriter, input string) {
	w.Header().Set("Content-Type", "application/json")
	ms := &FailMessage{
		Status:  false,
		Message: input,
	}

	response := jsonConvert(w, ms)
	jsonWrite(w, response)
}
