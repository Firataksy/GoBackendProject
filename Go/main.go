package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

type Message struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func signup(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var usersignup Sign
	var message Message

	err := json.NewDecoder(r.Body).Decode(&usersignup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersign = append(usersign, usersignup)
	_, control := user[usersignup.UName]

	if control != false {
		message.Status = false
		message.Message = "Username is used"
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	} else if control != true && usersignup.UName != "" && usersignup.Pwd != "" && usersignup.Name != "" && usersignup.SName != "" {
		message.Status = true
		message.Message = "Successful signup"
		currentID++
		usersignup.ID = currentID
		user[usersignup.UName] = usersignup
		GetMD5()
		messageJSON, _ := json.Marshal(message)
		usersJSON, err := json.Marshal(usersignup)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(messageJSON)
		w.Write(usersJSON)
		return
	} else {
		message.Status = false
		message.Message = "Information cannot be empty"
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	}
}

func GetMD5() {
	var sign Sign
	user := sign
	fmt.Println(user.Pwd, "a")
	hash := []byte(user.Pwd)
	user[sign.Pwd] = md5.Sum(hash)
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userlogin Login
	var message Message

	err := json.NewDecoder(r.Body).Decode(&userlogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userlogin.ID = currentID
	userl[userlogin.UName] = userlogin
	user, control := user[userlogin.UName]
	userlogin.ID = user.ID
	if control == true && user.Pwd == userlogin.Pwd {
		message.Status = true
		message.Message = "Succesful login"

		usersJSON, err := json.Marshal(userlogin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		w.Write(usersJSON)
		return
	} else {
		message.Status = false
		message.Message = "Wrong username or password"
		messageJSON, _ := json.Marshal(message)
		w.Write(messageJSON)
		return
	}
}

func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message Message

	idurl := r.URL.Query().Get("id")

	idInt, _ := strconv.Atoi(idurl)

	for _, user := range user {
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
		fmt.Print(user.ID)
	}

	message.Status = false
	message.Message = "Wrong id try again"
	mes, _ := json.Marshal(message)
	w.Write(mes)
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
