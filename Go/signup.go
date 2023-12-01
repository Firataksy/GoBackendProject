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
		errorJSON, _ := json.Marshal(error)
		w.Write(errorJSON)
		return
	} else if control != true && usersignup.UserName != "" && usersignup.Pwd != "" && usersignup.Name != "" && usersignup.SurName != "" {

		currentID++
		usersignup.ID = currentID
		usersignup.Pwd = md5Encode(usersignup.Pwd)
		message := Status.StatTrue(Stat{})
		user[usersignup.UserName] = usersignup

		signlogin.Status, signlogin.Data.ID, signlogin.Data.UserName = message, usersignup.ID, usersignup.UserName

		userJSON, _ := json.Marshal(signlogin)
		w.Write(userJSON)

		return
	} else {
		stat := Status.StatFalse(Stat{})
		error.Status, error.Message = stat, "Information Cannot be Empty"
		errorJSON, _ := json.Marshal(error)
		w.Write(errorJSON)
		return
	}
}
