package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func getUserData(w http.ResponseWriter, r *http.Request) {
	var user Sign

	idurl := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(idurl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	val, _ := rc.Get(context.Background(), "user:"+idurl).Result()

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
