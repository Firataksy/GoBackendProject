package main

import (
	"encoding/json"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userlogin Login
<<<<<<< HEAD
	var loginw Loginw
=======
>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104
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
<<<<<<< HEAD

	if control == true && user.Pwd == userlogin.Pwd {
		loginw.Status, loginw.Information.ID, loginw.Information.Uname = true, userlogin.ID, userlogin.UName
		userlogin.Pwd = ""
		usersJSON, err := json.Marshal(loginw)
=======
	if control == true && user.Pwd == userlogin.Pwd {
		message.Status = true
		message.Message = "Succesful login"
		userlogin.Pwd = ""
		usersJSON, err := json.Marshal(userlogin)
>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
<<<<<<< HEAD
		w.Write(usersJSON)
		fmt.Println(user)
=======
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		w.Write(usersJSON)
>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104
		return
	} else {
		message.Status = false
		message.Message = "Wrong username or password"
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	}
}
