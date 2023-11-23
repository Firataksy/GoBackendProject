package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	fmt.Println(user)
	rediset()
	mux.HandleFunc("/list", userlist)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}

func rediset() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	users, _ := user[]
	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	err = client.Set("userid", jsonData, 0).Err()
	if err != nil {
		fmt.Println("Redis error:", err)
		return
	}

	fmt.Println("Map successful registered redis")
}