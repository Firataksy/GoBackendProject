package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func updateUserData(w http.ResponseWriter, r *http.Request) {
	var updateuserdata Sign
	var updatenewuserdata UpdateNewUserData
	var updatelastuserdata UpdateLastUserData

	idurl := r.URL.Query().Get("id")

	er := json.NewDecoder(r.Body).Decode(&updatenewuserdata)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	checkuser, _ := rc.Get(context.Background(), "user:"+idurl).Result()
	checkuserid, _ := rc.Get(context.Background(), "user:"+updatenewuserdata.UserName).Result()
	checkusername, _ := rc.Get(context.Background(), "user:"+checkuserid).Result()

	if checkusername != "" {
		responseError(w, "Username is used")
		return
	}

	if checkuser == "" {
		responseError(w, "User not found")
		return
	}

	json.Unmarshal([]byte(checkuser), &updateuserdata)

	if updatenewuserdata.UserName == updateuserdata.UserName {
		responseError(w, "You already use this username")
		return
	}
	updatenewuserdata.ID = updateuserdata.ID

	if updatenewuserdata.UserName == "" {
		updatenewuserdata.UserName = updateuserdata.UserName
		updatelastuserdata.UserName = updateuserdata.UserName
	}
	if updatenewuserdata.Password == "" {
		updatenewuserdata.Password = updateuserdata.Password
	}
	if updatenewuserdata.Name == "" {
		updatenewuserdata.Name = updateuserdata.Name
		updatelastuserdata.Name = updateuserdata.Name
	}
	if updatenewuserdata.SurName == "" {
		updatenewuserdata.SurName = updateuserdata.SurName
		updatelastuserdata.SurName = updateuserdata.SurName
	}

	hashedpwd := md5Encode(updatenewuserdata.Password)

	updatenewuserdata.Password = hashedpwd

	ud := &UpdateLastUserData{
		ID:       updateuserdata.ID,
		UserName: updatenewuserdata.UserName,
		Name:     updatenewuserdata.Name,
		SurName:  updatenewuserdata.SurName,
	}

	_, userinfoerr := rc.Rename(context.Background(), "user:"+updateuserdata.UserName, "user:"+updatenewuserdata.UserName).Result()
	if userinfoerr != nil {
		log.Fatal("Redis rename error:", userinfoerr)
		return
	}

	jsonresponse := jsonConvert(w, updatenewuserdata)
	_, userseterr := rc.Set(context.Background(), "user:"+idurl, jsonresponse, 0).Result()
	if userseterr != nil {
		log.Fatal("Redis set update all user data sign error:", userseterr)
		return
	}
	_, usererr := rc.Set(context.Background(), "user:"+updatenewuserdata.UserName, idurl, 0).Result()
	if usererr != nil {
		log.Fatal("Redis set update user id data sign error:", usererr)
		return
	}

	responseSuccess(w, ud)
}
