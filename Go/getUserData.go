package main

import (
	"net/http"
	"strconv"
)

func getUserData(w http.ResponseWriter, r *http.Request) {

	idurl := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(idurl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := dataint[idInt]

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
