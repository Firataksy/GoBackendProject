package main

import (
	"encoding/json"
	"net/http"
)

var (
	data    = make(map[string]Sign)
	dataint = make(map[int]Sign)
)

type Sign struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Pwd      string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
}

type Login struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type SuccessData struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
}

type SuccessMessage struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

type FailMessage struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Status bool `json:"status"`
	Data   interface{}
}

func jsonConvert(w http.ResponseWriter, input interface{}) []byte {
	Json, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Json Error", http.StatusInternalServerError)
		return nil
	}
	return Json
}

func jsonWrite(w http.ResponseWriter, input []byte) {

	w.Write(input)
}

func responseSuccess(w http.ResponseWriter, input interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rp := Response{
		Status: true,
		Data:   input,
	}

	response := jsonConvert(w, rp)
	jsonWrite(w, response)
}

func responseError(w http.ResponseWriter, input string) {
	w.Header().Set("Content-Type", "application/json")
	ms := &FailMessage{
		Status:  false,
		Message: input,
	}

	response := jsonConvert(w, ms)
	jsonWrite(w, response)
}
