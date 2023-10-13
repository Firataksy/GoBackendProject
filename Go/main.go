package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var usersign = []Usersign{
	{ID: 1, UName: "", Pwd: "", Name: "", SName: ""},
}

var userlogin = []Userlogin{
	{UName: "", Pwd: ""},
}

type Usersign struct {
	ID    int
	UName string `json:"username"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}
type Userlogin struct {
	UName string `json:"username"`
	Pwd   string `json:"password"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userby Usersign
	err := json.NewDecoder(r.Body).Decode(&userby)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersign = append(usersign, userby)

	if userby.UName != "" && userby.Pwd != "" && userby.Name != "" && userby.SName != "" {
		fmt.Fprint(w, `{"success": "True", "message": "Successful signup"}`, userby.UName, " ", userby.ID)
	} else {
		fmt.Fprint(w, `{"success": "False", "message": " Information cannot be empty"}`)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var userbyl Userlogin
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&userbyl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var users bool
	for _, user := range usersign {
		if user.UName == userbyl.UName && user.Pwd == userbyl.Pwd {
			users = true
			break
		}
	}

	if users {
		fmt.Fprint(w, `{"success": "True", "message": "Successful login"}`)
	} else {
		fmt.Fprint(w, `{"success": "False", "message": "Wrong username or password"}`)
	}
}
func idd() {
	var ids Usersign

	if ids.ID != -1 {
		ids.ID++
		return
	}
	return
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
