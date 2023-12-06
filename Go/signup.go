package main

import (
	"encoding/json"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usersignup Sign
	var signlogin Signlogin
	var error Error

	err := json.NewDecoder(r.Body).Decode(&usersignup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, control := user[usersignup.UserName]

	if control != false {
		stat := Status.StatFalse(Stat{})
		error.Status, error.Message = stat, "Username is Used"
		Jsonwrite(w, error)
		return
	} else if control != true && usersignup.UserName != "" && usersignup.Pwd != "" && usersignup.Name != "" && usersignup.SurName != "" {
		currentID++
		usersignup.ID = currentID
		usersignup.Pwd = md5Encode(usersignup.Pwd)
		user[usersignup.UserName] = usersignup
		user[usersignup.ID] = usersignup
		message := Status.StatTrue(Stat{})
		signlogin.Status, signlogin.Data.ID, signlogin.Data.UserName = message, usersignup.ID, usersignup.UserName
		Jsonwrite(w, signlogin)
		return
	} else {
		stat := Status.StatFalse(Stat{})
		error.Status, error.Message = stat, "Information Cannot be Empty"
		Jsonwrite(w, error)
		return
	}
}
