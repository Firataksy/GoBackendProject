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
	var user User
	var id UserId
	err := json.NewDecoder(r.Body).Decode(&usersignup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	check, _ := rc.Get(context.Background(), "user:"+id.ID).Result()
	fmt.Println(id.ID)
	json.Unmarshal([]byte(check), &user)

	if user.UserName == usersignup.UserName {
		responseError(w, "Username is used")
		return
	}

	if usersignup.UserName != "" && usersignup.Pwd != "" && usersignup.Name != "" && usersignup.SurName != "" {

		usersignup.Pwd = md5Encode(usersignup.Pwd)

		userıncrid, _ := rc.Incr(context.Background(), "userIncrId").Result()
		usersignup.ID = int(userıncrid)

		sm := SuccessMessage{
			ID:       int(userıncrid),
			UserName: usersignup.UserName,
		}

		response := jsonConvert(w, usersignup)
		stringid := strconv.Itoa(int(userıncrid))
		_, userinfoerr := rc.Set(context.Background(), "user:"+stringid, response, 0).Result()
		if userinfoerr != nil {
			fmt.Println("Redis set user sign error:", userinfoerr)
			return
		}
		responseSuccess(w, sm)
		return
	}
	responseError(w, "Information cannot be empty")
}
