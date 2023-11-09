package main

import (
	"encoding/json"
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

	_, control := user[usersignup.UName]

	if control != false {
		message.Status = false
		message.Message = "Username is used"
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	} else if control != true && usersignup.UName != "" && usersignup.Pwd != "" && usersignup.Name != "" && usersignup.SName != "" {
		message.Status = true
		message.Message = "Successful signup"
		currentID++
		usersignup.ID = currentID
		usersignup.Pwd = md5Encode(usersignup.Pwd)
		user[usersignup.UName] = usersignup
		usersignup.Pwd = ""
		usersign = append(usersign, usersignup)
		userJSON, _ := json.Marshal(usersignup)
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		w.Write(userJSON)
		return
	} else {
		message.Status = false
		message.Message = "Information cannot be empty"
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	}
}
