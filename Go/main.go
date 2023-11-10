package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/list", getusers)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
