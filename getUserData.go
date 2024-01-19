package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func getUserData(w http.ResponseWriter, r *http.Request) {
	var user User
	token := r.Header.Get("token")
	idToken := tokenToID(w, token)
	if idToken == "" {
		return
	}

	idInt, _ := strconv.Atoi(idToken)

	val, _ := rc.Get(context.Background(), "player_"+idToken).Result()

	json.Unmarshal([]byte(val), &user)

	if user.ID == idInt && idInt != 0 {

		sd := SuccessData{
			ID:       user.ID,
			UserName: user.UserName,
			Name:     user.Name,
			SurName:  user.SurName,
		}

		responseSuccess(w, sd)
		return
	}
	responseError(w, "User not found")
}
