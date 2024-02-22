package api

import (
	"context"
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

	userID, _ := rc.Client.Get(context.Background(), "userID:"+searchedUserName).Result()

	if userID == "" {
		ResponseFail(w, "user not found")
		return
	}

	if headerUserID == userID {
		ResponseFail(w, "you can not search yourself")
		return
	}
	intID, _ := strconv.Atoi(userID)
	userSearch := UserSearchID{
		ID: intID,
	}

	ResponseSuccess(w, userSearch)
}
