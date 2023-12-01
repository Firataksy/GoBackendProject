package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func userdata(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var Data MainData
	var error Error

	idurl := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(idurl)

	for _, user := range user {
		if user.ID == idInt {
			Data.Status, Data.Data.ID, Data.Data.UserName, Data.Data.Name, Data.Data.SurName = Status.StatTrue(Stat{}), user.ID, user.UserName, user.Name, user.SurName
			userdata, _ := json.Marshal(Data)
			w.Write(userdata)
			return
		}
	}

	stat := Status.StatFalse(Stat{})
	error.Status, error.Message = stat, "User not found"
	errorJSON, _ := json.Marshal(error)
	w.Write(errorJSON)
}
