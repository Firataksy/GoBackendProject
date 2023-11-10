package main

import (
	"encoding/json"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var usersignup Sign
	var message Message
<<<<<<< HEAD
	var sign Signw
=======
>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104

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
<<<<<<< HEAD
=======
		message.Status = true
		message.Message = "Successful signup"
>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104
		currentID++
		usersignup.ID = currentID
		usersignup.Pwd = md5Encode(usersignup.Pwd)
		user[usersignup.UName] = usersignup
		usersignup.Pwd = ""
		usersign = append(usersign, usersignup)
<<<<<<< HEAD
		sign.Status, sign.Information.ID, sign.Information.Uname = true, usersignup.ID, usersignup.UName
		userJSON, _ := json.Marshal(sign)
		w.Write(userJSON)
		return
	} else {
=======
		userJSON, _ := json.Marshal(usersignup)
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		w.Write(userJSON)
		return
	} else {

>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104
		message.Status = false
		message.Message = "Information cannot be empty"
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	}
}
