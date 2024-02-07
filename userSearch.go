package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func userSearch(w http.ResponseWriter, r *http.Request) {

	urlUserName := r.URL.Query().Get("username")
	fmt.Println(urlUserName)
	headerUserID := r.Header.Get("userid")
	fmt.Println(headerUserID)
	userID, _ := rc.Get(context.Background(), "userID:"+urlUserName).Result()
	if headerUserID != userID {
		ID, _ := strconv.Atoi(userID)
		responseSuccess(w, ID)
		return
	}
	if headerUserID == userID {
		responseError(w, "you can not search yourself")
		return
	}
	responseError(w, "user not found")
}
