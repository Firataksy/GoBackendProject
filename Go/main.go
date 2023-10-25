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

var (
	user      = make(map[string]Usersign)
	userl     = make(map[string]Userlogin)
	currentID = 1
)

type Usersign struct {
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}

type Userlogin struct {
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usersignup Usersign
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
	var userbylogin Userlogin

	err := json.NewDecoder(r.Body).Decode(&userbylogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userl[userbylogin.UName] = userbylogin

	for _, user := range user {
		fmt.Println(user.UName, " ", userbylogin.UName, " ", userbylogin.Pwd, " ", user.Pwd)
		if user.UName == userbylogin.UName && user.Pwd == userbylogin.Pwd {
			fmt.Fprint(w, "Status: True", "\nMessage: Successful login", "\nUserıd: ", user.ID, "\nUsername: ", userbylogin.UName)
			return
		} else {
			fmt.Fprint(w, "Status: False", "\nMessage: Wrong username or password")
			return
		}
	}
}

func getusers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "----Informations----\n")
	for _, user := range user {
		if user.ID != 0 {
			fmt.Fprint(w, "Status: True", "\nUserId: ", user.ID, "\nUsername: ", user.UName, "\nUserFirstName: ", user.Name, "\nUserLastName: ", user.SName, "\n----------------------", "\n")
		}
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
