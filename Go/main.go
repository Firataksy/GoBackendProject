package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var usersign = []Usersign{
	{UName: "", Pwd: "", Name: "", SName: ""},
}

var userlogin = []Userlogin{
	{UName: "", Pwd: ""},
}

type Usersign struct {
	ID    int    `json:"ıd"`
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

		fmt.Fprint(w, `{"success": true, "message": "Successful signup"}`, userby.Name)
		return
	} else {
		fmt.Fprint(w, `{"success": true, "message": " Information cannot be empty"}`)
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
		fmt.Fprint(w, `{"success": true, "message": "Successful login"}`)
	} else {
		fmt.Fprint(w, `{"success": false, "message": "Wrong username or password"}`)
	}
}

/*func login(w http.ResponseWriter, r *http.Request) {
	var userbyl Userlogin
	var userbys Usersign
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&userbyl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userbys.UName == userbyl.UName && userbys.Pwd == userbyl.Pwd {
		fmt.Fprint(w, `{"success": true, "message": "Successful login"}`")
		fmt.Fprint(w, userbyl, &userbys)
	} else {
		fmt.Fprint(w, `{"success": false, "message": "Wrong username or password"})
		fmt.Fprint(w, userbyl, &userbys)
	}
}*/

func listusers(w http.ResponseWriter, r *http.Request) {
	var list Usersign
	fmt.Fprint(w, list)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/listusers", listusers)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
