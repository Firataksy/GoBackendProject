package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message Message

	idurl := r.URL.Query().Get("id")

	idInt, _ := strconv.Atoi(idurl)
	for _, user := range usersign {
		if user.ID == idInt {
			message.Status = true
			message.Message = "Successfully listed"
			mes, _ := json.Marshal(message)
			userlist, _ := json.Marshal(user)
			w.Write(mes)
			fmt.Fprint(w, "\n")
			w.Write(userlist)
			return
		}
	}
	message.Status = false
	message.Message = "Wrong id try again "
	mes, _ := json.Marshal(message)
	w.Write(mes)
}
