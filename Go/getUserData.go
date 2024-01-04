package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func getUserData(w http.ResponseWriter, r *http.Request) {
	var user User

	idURL := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(idURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	val, _ := rc.Get(context.Background(), "user:"+idURL).Result()

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
