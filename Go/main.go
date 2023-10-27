package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var usersign = []Sign{
	{UName: "", Pwd: "", Name: "", SName: ""},
}

var userlogin = []Login{
	{UName: "", Pwd: ""},
}

var (
	user      = make(map[string]Sign)
	userl     = make(map[string]Login)
	currentID = 1
)

type Sign struct {
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}

type Login struct {
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usersignup Sign

	err := json.NewDecoder(r.Body).Decode(&usersignup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersign = append(usersign, usersignup)
	_, control := user[usersignup.UName]

	if control != false {
		fmt.Fprint(w, "Username is used")
		return
	} else if control != true && usersignup.UName != "" && usersignup.Pwd != "" && usersignup.Name != "" && usersignup.SName != "" {
		usersignup.ID = currentID
		currentID++
		user[usersignup.UName] = usersignup
		fmt.Fprint(w, "Status: True", "\nMessage: Successful signup", "\nUserıd: ", usersignup.ID, "\nUsername: ", usersignup.UName, "\n")
		return
	} else if usersignup.UName == "" && usersignup.Pwd == "" && usersignup.Name == "" && usersignup.SName == "" {
		fmt.Fprint(w, "Status: False", "\nMessage: Information cannot be empty")
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userlogin Login

	err := json.NewDecoder(r.Body).Decode(&userlogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userl[userlogin.UName] = userlogin

	user, control := user[userlogin.UName]

	if control == true && user.Pwd == userlogin.Pwd {
		fmt.Fprint(w, "Status: True", "\nMessage: Successful login", "\nUserId: ", user.ID, "\nUsername: ", user.UName)
	} else {
		fmt.Fprint(w, "Status: False", "\nMessage: Wrong username or password")
	}
}

func pwdhash() {

}

func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, "----Informations----\n")
	for _, user := range user {
		fmt.Fprint(w, "Status: True", "\nUserId: ", user.ID, "\nUsername: ", user.UName, "\nUserFirstName: ", user.Name, "\nUserLastName: ", user.SName, "\n----------------------", "\n")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/list", getusers)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
