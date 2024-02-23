package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
)

func (rc *RedisClient) UserSearch(w http.ResponseWriter, r *http.Request) {

	searchedUserName := r.URL.Query().Get("username")

	headerUserID := r.Header.Get("userid")

	if searchedUserName == "" {
		ResponseFail(w, "can not be empty username in url")
		return
	}

	userID, err := rc.Client.Get(context.Background(), "userID:"+searchedUserName).Result()
	if err != nil {
		log.Fatal("userSearch get userID err :", err)
		return
	}

	if userID == "" {
		ResponseFail(w, "user not found")
		return
	}

	if headerUserID == userID {
		ResponseFail(w, "you can not search yourself")
		return
	}

	intID, err := strconv.Atoi(userID)
	if err != nil {
		log.Fatal("userSearch convert err :", err)
		return
	}

	userSearch := UserSearchID{
		ID: intID,
	}

	ResponseSuccess(w, userSearch)
}
