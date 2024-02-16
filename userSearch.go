package main

import (
	"context"
	"net/http"
)

func userSearch(w http.ResponseWriter, r *http.Request) {

	searchedUserName := r.URL.Query().Get("username")

	ID := r.Header.Get("userid")

	if searchedUserName == "" {
		responseFail(w, "can not be empty username in url")
		return
	}

	userID, _ := rc.Get(context.Background(), "userID:"+searchedUserName).Result()

	if userID == "" {
		responseFail(w, "user not found")
		return
	}

	if ID == userID {
		responseFail(w, "you can not search yourself")
		return
	}

	responseSuccess(w, userID)
}
