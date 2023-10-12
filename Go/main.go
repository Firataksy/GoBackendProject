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
	json.NewEncoder(w).Encode(&userby)

	if userby.UName != "" && userby.Pwd != "" && userby.Name != "" && userby.SName != "" {
		usersign = append(usersign, userby)
		fmt.Fprint(w, "true ", "Succesful signup", userby.Name)
		return
	} else {
		fmt.Fprint(w, "false ", "Information cannot be empty")
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var userbyl Userlogin
	var userbys Usersign

	w.Header().Set("Content-Type", "application/json")

	if userbyl.UName != "" && userbyl.Pwd != "" {
		userlogin = append(userlogin, userbyl)
		fmt.Fprint(w, "true ", "Succesful login")
		return
	} else if userbyl.UName == "" && userbyl.Pwd == "" {
		fmt.Fprint(w, "false ", "Information cannot be empty")
		return
	} else if userbyl.UName != userbys.UName && userbyl.Pwd != userbys.Pwd {
		fmt.Fprint(w, "false ", "wrong username or password")
		return
	}
	fmt.Fprint(w, userbyl)
}

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
