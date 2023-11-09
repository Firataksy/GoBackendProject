package main

import (
	"encoding/json"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userlogin Login
	var message Message

	err := json.NewDecoder(r.Body).Decode(&userlogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userlogin.Pwd = md5Encode(userlogin.Pwd)
	userlogin.ID = currentID
	userl[userlogin.UName] = userlogin
	user, control := user[userlogin.UName]
	userlogin.ID = user.ID
	if control == true && user.Pwd == userlogin.Pwd {
		message.Status = true
		message.Message = "Succesful login"

		usersJSON, err := json.Marshal(userlogin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		w.Write(usersJSON)
		return
	} else {
		message.Status = false
		message.Message = "Wrong username or password"
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	}
}
