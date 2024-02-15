package main

import (
	"context"
	"net/http"
	"strconv"
)

func userSearch(w http.ResponseWriter, r *http.Request) {

	searchedUserName := r.URL.Query().Get("username")

	ID := r.Header.Get("userid")

	if searchedUserName == "" {
		responseError(w, "can not be empty username in url")
		return
	}

	userID, _ := rc.Get(context.Background(), "userID:"+searchedUserName).Result()

	if userID == "" {
		responseError(w, "user not found")
		return
	}

	if ID == userID {
		responseError(w, "you can not search yourself")
		return
	}

	intID, _ := strconv.Atoi(userID)
	responseSuccess(w, intID)
}
