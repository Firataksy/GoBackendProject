package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/list", userlist)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:9000",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}
