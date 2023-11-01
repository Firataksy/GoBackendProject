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

var getid = []Getid{
	{ID: 0},
}

var (
	user      = make(map[string]Sign)
	userl     = make(map[string]Login)
	currentID = 0
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
}

type Getid struct {
	ID int `json:"id"`
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
		currentID++
		usersignup.ID = currentID
		user[usersignup.UName] = usersignup
		usersJSON, err := json.Marshal(usersignup)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "Status: True", "\nMessage: Successful signup\n")
		w.Write(usersJSON)
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

	userlogin.ID = currentID
	if userlogin.UName != "" && userlogin.Pwd != "" {
		userl[userlogin.UName] = userlogin
	}

	userJSON, err := json.Marshal(userlogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, control := user[userlogin.UName]

	if control == true && user.Pwd == userlogin.Pwd {
		fmt.Fprint(w, "Status: True", "\nMessage: Successful login\n")
		w.Write(userJSON)
		return
	} else {
		fmt.Fprint(w, "Status: False", "\nMessage: Wrong username or password")
		return
	}
}

func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id Getid
	var user Login
	fmt.Fprint(w, "----Informations----\n")

	err := json.NewDecoder(r.Body).Decode(&getid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userJSON, err := json.Marshal(userl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.ID == id.ID {
		w.Write(userJSON)
		return
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
