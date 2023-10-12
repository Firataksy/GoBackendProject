package main

import (
	"fmt"
	"net/http"
)

var userregister = []Users{
	{UName: "fırat123", Pwd: "123", Name: "Fırat", SName: "Aksoy"},
	{UName: "", Pwd: "", Name: "", SName: ""},
}

type Users struct {
	UName string `json:"uname"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	var userby Users
	userregister = append(userregister, userby)

	if userby.UName != "" && userby.Pwd != "" && userby.Name != "" && userby.SName != "" {
		fmt.Fprint(w, "true", "Succesful signup", userby.Name)
	} else {
		fmt.Fprint(w, "false ", "Information cannot be empty")
		return
	}
}

func listusers(w http.ResponseWriter, r *http.Request) {
	var list Users
	fmt.Fprint(w, list)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/getusers", listusers)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
