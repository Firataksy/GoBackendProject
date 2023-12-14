package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var usersignup Sign

	err := json.NewDecoder(r.Body).Decode(&usersignup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	check, _ := rc.Get(context.Background(), "checkusername:"+usersignup.UserName).Result()

	if check == usersignup.UserName && usersignup.UserName != "" {
		responseError(w, "Username is used")
		return
	}

	if usersignup.UserName != "" && usersignup.Pwd != "" && usersignup.Name != "" && usersignup.SurName != "" {

		usersignup.Pwd = md5Encode(usersignup.Pwd)

		userid, _ := rc.Incr(context.Background(), "ID").Result()
		usersignup.ID = int(userid)

		sm := SuccessMessage{
			ID:       int(userid),
			UserName: usersignup.UserName,
		}

		response := jsonConvert(w, usersignup)
		ıd := strconv.Itoa(int(userid))

		_, err := rc.Set(context.Background(), "userinfo:"+ıd, response, 0).Result()
		if err != nil {
			fmt.Println("Redis set user info error:", err)
			return
		}

		_, er := rc.Set(context.Background(), "checkusername:"+usersignup.UserName, usersignup.UserName, 0).Result()
		if er != nil {
			fmt.Println("Redis set check username error:", er)
			return
		}

		data[usersignup.UserName] = usersignup
		dataint[usersignup.ID] = usersignup
		responseSuccess(w, sm)
		return
	}

	responseError(w, "Information cannot be empty")
}
