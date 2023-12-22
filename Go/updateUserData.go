package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func updateUserData(w http.ResponseWriter, r *http.Request) {
	var userdata UserData
	var updatenewuserdata UpdateNewUserData

	idurl := r.URL.Query().Get("id")

	er := json.NewDecoder(r.Body).Decode(&updatenewuserdata)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	checkuser, _ := rc.Get(context.Background(), "user:"+idurl).Result()

	if checkuser == "" {
		responseError(w, "User not found")
		return
	}

	json.Unmarshal([]byte(checkuser), &userdata)

	if updatenewuserdata.UserName == userdata.UserName {
		responseError(w, "You already use this username")
		return
	}

	checkuserid, _ := rc.Get(context.Background(), "user:"+updatenewuserdata.UserName).Result()

	if checkuserid != "" {
		responseError(w, "Username is used")
		return
	}

	if updatenewuserdata.UserName != "" {
		_, userrenameerr := rc.Rename(context.Background(), "user:"+userdata.UserName, "user:"+updatenewuserdata.UserName).Result()
		if userrenameerr != nil {
			log.Fatal("Redis rename error:", userrenameerr)
			return
		}
	}

	if updatenewuserdata.UserName != "" {
		userdata.UserName = updatenewuserdata.UserName
	}
	if updatenewuserdata.Password != "" {
		userdata.Password = updatenewuserdata.Password
	}
	if updatenewuserdata.Name != "" {
		userdata.Name = updatenewuserdata.Name
	}
	if updatenewuserdata.SurName != "" {
		userdata.SurName = updatenewuserdata.SurName
	}

	jsonresponse := jsonConvert(w, userdata)
	userdata.Password = ""

	_, userseterr := rc.Set(context.Background(), "user:"+idurl, jsonresponse, 0).Result()
	if userseterr != nil {
		log.Fatal("Redis set update all new user data sign error:", userseterr)
		return
	}

	responseSuccess(w, userdata)
}
