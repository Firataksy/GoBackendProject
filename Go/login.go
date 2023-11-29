package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userlogin Login

	err := json.NewDecoder(r.Body).Decode(&userlogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userlogin.Password = md5Encode(userlogin.Password)
	userl[userlogin.UserName] = userlogin
	user := user[userlogin.UserName]

	if user.Pwd == userlogin.Password && user.UserName == userlogin.UserName {
		message := Statustrue()
		usersJSON, err := json.Marshal(Loginr(message.Status, user.ID, userlogin.UserName))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(usersJSON)
		fmt.Println(user)
		return
	} else {
		message := Statusfalse()
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	}
}
