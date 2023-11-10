package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
	var list Listw
	idurl := r.URL.Query().Get("id")

	idInt, _ := strconv.Atoi(idurl)
	for _, user := range usersign {
		if idInt != 0 {
			if user.ID == idInt {
				list.Status, list.Information.ID, list.Information.Uname, list.Information.Name, list.Information.SName = true, user.ID, user.UName, user.Name, user.SName
				message.Message = "Successfully listed"
				userlist, _ := json.Marshal(list)
				w.Write(userlist)
				return
			}
		}
	}
	message.Status = false
	message.Message = "User not found"
	mes, _ := json.Marshal(message)
	w.Write(mes)
}
