package main

import (
	"net/http"
	"strconv"
)

func getUserData(w http.ResponseWriter, r *http.Request) {

	idurl := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(idurl)
	user, _ := dataint[idInt]

	if user.ID == idInt && idInt != 0 {
		responseSuccess(w, user.ID, user.UserName)
		return
	}
	responseError(w, "dataerror")

}
