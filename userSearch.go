package main

import (
	"context"
	"net/http"
	"strconv"
)

func userSearch(w http.ResponseWriter, r *http.Request) {

	urlUserName := r.URL.Query().Get("username")

	headerUserID := r.Header.Get("userid")

	if urlUserName == "" {
		responseError(w, "can not be empty username in url")
		return
	}

	userID, _ := rc.Get(context.Background(), "userID:"+urlUserName).Result()

	if userID == "" {
		responseError(w, "user not found")
		return
	}

	if headerUserID == userID {
		responseError(w, "you can not search yourself")
		return
	}

	ID, _ := strconv.Atoi(userID)
	responseSuccess(w, ID)
}
