package main

import (
	"encoding/json"
	"net/http"
)

var (
	data      = make(map[string]Sign)
	dataint   = make(map[int]Sign)
	currentID = 0
)

type Sign struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Pwd      string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
}

type Login struct {
	Status   bool   `json:"status"`
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type userData struct {
	Status bool `json:"status"`
	Data   struct {
		ID       int    `json:"id"`
		UserName string `json:"username"`
		Name     string `json:"name"`
		SurName  string `json:"surname"`
	}
}

type SuccessMessage struct {
	Status bool `json:"status"`
	Data   struct {
		ID       int    `json:"id"`
		UserName string `json:"username"`
	}
}

type FailMessage struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func jsonWrite(w http.ResponseWriter, input interface{}) {
	w.Header().Set("Content-Type", "application/json")
	Json, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Json Error", http.StatusInternalServerError)
		return
	}
	w.Write(Json)
}

func responseSuccess(w http.ResponseWriter, id int, username string) {
	w.Header().Set("Content-Type", "application/json")

	sm := &SuccessMessage{
		Status: true,
		Data: struct {
			ID       int    `json:"id"`
			UserName string `json:"username"`
		}{
			ID:       id,
			UserName: username},
	}

	w.WriteHeader(http.StatusOK)
	jsonWrite(w, sm)
}

func responseError(w http.ResponseWriter, input string) {
	var Message FailMessage

	Message.Status = false

	if input == "usernameerror" {
		Message.Message = "Username is used"

	} else if input == "signerror" {
		Message.Message = "Information cannot be empty"

	} else if input == "loginerror" {
		Message.Message = "Wrong username or password"

	} else if input == "dataerror" {
		Message.Message = "User not found"

	}
	jsonWrite(w, Message)
}
