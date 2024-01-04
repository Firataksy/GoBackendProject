package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func updateUserData(w http.ResponseWriter, r *http.Request) {
	var userData UserData
	var updateNewUserData UpdateNewUserData

	idUrl := r.URL.Query().Get("id")

	er := json.NewDecoder(r.Body).Decode(&updateNewUserData)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	checkUser, _ := rc.Get(context.Background(), "user:"+idUrl).Result()

	if checkUser == "" {
		responseError(w, "User not found")
		return
	}

	json.Unmarshal([]byte(checkUser), &userData)

	if updateNewUserData.UserName == userData.UserName {
		responseError(w, "You already use this username")
		return
	}

	checkUserID, _ := rc.Get(context.Background(), "user:"+updateNewUserData.UserName).Result()

	if checkUserID != "" {
		responseError(w, "Username is used")
		return
	}

	if updateNewUserData.UserName != "" {
		_, userRenameErr := rc.Rename(context.Background(), "user:"+userData.UserName, "user:"+updateNewUserData.UserName).Result()
		if userRenameErr != nil {
			log.Fatal("Not Renamed Username", userRenameErr)
		}
		userData.UserName = updateNewUserData.UserName
	}
	if updateNewUserData.Password != "" {
		hashPWD := md5Encode(updateNewUserData.Password)
		userData.Password = hashPWD
	}
	if updateNewUserData.Name != "" {
		userData.Name = updateNewUserData.Name
	}
	if updateNewUserData.SurName != "" {
		userData.SurName = updateNewUserData.SurName
	}
	jsonResponse := jsonConvert(w, userData)
	userData.Password = ""

	_, userSetErr := rc.Set(context.Background(), "user:"+idUrl, jsonResponse, 0).Result()
	if userSetErr != nil {
		log.Fatal("Redis set update all new user data sign error:", userSetErr)
		return
	}

	responseSuccess(w, userData)
}
