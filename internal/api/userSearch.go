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
		responseFail(w, "can not be empty username in url")
		return
	}

	userID, _ := rc.Client.Get(context.Background(), "userID:"+searchedUserName).Result()

	if userID == "" {
		responseFail(w, "user not found")
		return
	}

	if headerUserID == userID {
		responseFail(w, "you can not search yourself")
		return
	}
	intID, _ := strconv.Atoi(userID)
	userSearch := UserSearchID{
		ID: intID,
	}

	responseSuccess(w, userSearch)
}
