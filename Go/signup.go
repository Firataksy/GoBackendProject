package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var usersignup Sign
	var user User

	err := json.NewDecoder(r.Body).Decode(&usersignup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkusername, _ := rc.Get(context.Background(), "user:"+usersignup.UserName).Result()

	check, _ := rc.Get(context.Background(), "user:"+checkusername).Result()

	json.Unmarshal([]byte(check), &user)

	if user.UserName == usersignup.UserName {
		responseError(w, "Username is used")
		return
	}

	if usersignup.UserName != "" && usersignup.Password != "" && usersignup.Name != "" && usersignup.SurName != "" {

		usersignup.Password = md5Encode(usersignup.Password)

		userIncrID, _ := rc.Incr(context.Background(), "userIncrId").Result()

		sm := SuccessMessage{
			ID:       int(userIncrID),
			UserName: usersignup.UserName,
		}

		usersignup.ID = int(userIncrID)
		redisalldata := jsonConvert(w, usersignup)
		stringid := strconv.Itoa(int(userIncrID))

		_, userinfoerr := rc.Set(context.Background(), "user:"+stringid, redisalldata, 0).Result()
		if userinfoerr != nil {
			log.Fatal("Redis set user data sign error:", userinfoerr)
			return
		}

		_, usererr := rc.Set(context.Background(), "user:"+usersignup.UserName, userIncrID, 0).Result()
		if usererr != nil {
			log.Fatal("Redis set user id sign error:", usererr)
			return
		}

		responseSuccess(w, sm)
		return
	}
	responseError(w, "Information cannot be empty")
}
