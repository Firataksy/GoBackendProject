package main

import (
	"encoding/json"
	"net/http"
)

var (
	user      = make(map[interface{}]Sign)
	currentID = 0
)

type Sign struct {
	Status   bool   `json:"status"`
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

type MainData struct {
	Status bool `json:"status"`
	Data   struct {
		ID       int    `json:"id"`
		UserName string `json:"username"`
		Name     string `json:"name"`
		SurName  string `json:"surname"`
	}
}

type Signlogin struct {
	Status bool `json:"status"`
	Data   struct {
		ID       int    `json:"id"`
		UserName string `json:"username"`
	}
}

type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Status interface {
	StatTrue() bool
	StatFalse() bool
}

func (Stat) StatTrue() bool {
	return true
}
func (Stat) StatFalse() bool {
	return false
}

type Stat struct {
}

func Jsonwrite(w http.ResponseWriter, input interface{}) {
	w.Header().Set("Content-Type", "application/json")
	Json, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Json Error", http.StatusInternalServerError)
		return
	}
	w.Write(Json)
}
