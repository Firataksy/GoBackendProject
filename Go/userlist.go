package main

import (
	"encoding/json"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104
	"net/http"
	"strconv"
)

func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
<<<<<<< HEAD
	var list Listw
=======
>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104

	idurl := r.URL.Query().Get("id")

	idInt, _ := strconv.Atoi(idurl)
	for _, user := range usersign {
<<<<<<< HEAD
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
=======
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
>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104
	mes, _ := json.Marshal(message)
	w.Write(mes)
}
