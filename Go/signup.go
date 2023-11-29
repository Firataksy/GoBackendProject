package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usersignup Sign
	var message Message

	err := json.NewDecoder(r.Body).Decode(&usersignup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, control := user[usersignup.UserName]

	if control != false {
		message := Statusfalse()
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	} else if control != true && usersignup.UserName != "" && usersignup.Pwd != "" && usersignup.Name != "" && usersignup.SurName != "" {
		currentID++
		usersignup.ID = currentID
		usersignup.Pwd = md5Encode(usersignup.Pwd)
		message = Statustrue()
		user[usersignup.UserName] = usersignup
		userJSON, _ := json.Marshal(Signr(message.Status, usersignup.ID, usersignup.UserName))
		w.Write(userJSON)
		fmt.Println(user)
		return
	} else {
		message = Statusfalse()
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	}
}
