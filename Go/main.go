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
	usersid   = make(map[int]Userid)
	user      = make(map[string]Usersign)
	currentID = 1
)

type Userid struct {
	ID int `json:"id"`
}

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
	var userby Usersign
	var userid Userid

	err := json.NewDecoder(r.Body).Decode(&userby)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersign = append(usersign, userby)

	if userby.UName != "" && userby.Pwd != "" && userby.Name != "" && userby.SName != "" {
		userid.ID = currentID
		currentID++
		usersid[userid.ID] = userid
		fmt.Fprint(w, "Status: True", "\nMessage: Successful signup", "\nUserıd: ", userid.ID, "\nUsername: ", userby.UName)
		return
	} else {
		fmt.Fprint(w, "Status: False", "\nMessage: Information cannot be empty")
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userbyl Userlogin
	var userid Userid

	err := json.NewDecoder(r.Body).Decode(&userbyl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for id, user := range usersign {
		if user.UName == userbyl.UName && user.Pwd == userbyl.Pwd {
			userid.ID = id
			fmt.Fprint(w, "Status: True", "\nMessage: Successful login", "\nUserıd: ", userid.ID, "\nUsername: ", userbyl.UName)
			return
		}
	}
	fmt.Fprint(w, "Status: False", "\nMessage: Wrong username or password")
}

func getusers(w http.ResponseWriter, r *http.Request) {
	var userid Userid
	fmt.Fprint(w, "----Informations----\n")

	for id, user := range usersign {
		if id != 0 {
			userid.ID = id
			fmt.Fprint(w, "Status: True", "\nUserId: ", userid.ID, "\nUsername: ", user.UName, "\nUserFirstName: ", user.Name, "\nUserLastName: ", user.SName, "\n----------------------", "\n")
		}

	}
	return
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
