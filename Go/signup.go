package main

import (
	"encoding/json"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var usersignup Sign

	err := json.NewDecoder(r.Body).Decode(&usersignup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, control := data[usersignup.UserName]

	if control != false {
		responseError(w, "usernameerror")
		return
	}

	if control != true && usersignup.UserName != "" && usersignup.Pwd != "" && usersignup.Name != "" && usersignup.SurName != "" {
		currentID++
		usersignup.ID = currentID
		usersignup.Pwd = md5Encode(usersignup.Pwd)
		data[usersignup.UserName] = usersignup
		dataint[usersignup.ID] = usersignup
		responseSuccess(w, usersignup.ID, usersignup.UserName)
		return
	}

	responseError(w, "signerror")
	return

}
