package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	w.WriteHeader(http.StatusOK)

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
