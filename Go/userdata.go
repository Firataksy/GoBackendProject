package main

import (
	"encoding/json"
	"net/http"
)

func userdata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
	found := false
	idurl := r.URL.Query().Get("id")

	user, found := user["ID"]

	if found != true {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if found != false {
		userJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
		w.Write(userJSON)
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Write(userJSON)
	message = Statusfalse()
	mes, _ := json.Marshal(message)
	w.Write(mes)
}
