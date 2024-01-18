package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func updateUserData(w http.ResponseWriter, r *http.Request) {
	var userData *Sign
	var updatedUser UpdatedUser
	var updateNewUserData UpdateNewUserData

	idUrl := r.URL.Query().Get("id")

	er := json.NewDecoder(r.Body).Decode(&updateNewUserData)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	checkUser, _ := rc.Get(context.Background(), "player_"+idUrl).Result()

	if checkUser == "" {
		responseError(w, "User not found")
		return
	}

	json.Unmarshal([]byte(checkUser), &userData)
	json.Unmarshal([]byte(checkUser), &updatedUser)

	if updateNewUserData.UserName == userData.UserName {
		responseError(w, "You already use this username")
		return
	}

	checkUserID, _ := rc.Get(context.Background(), updateNewUserData.UserName).Result()

	if checkUserID != "" {
		responseError(w, "Username is used")
		return
	}

	if updateNewUserData.UserName != "" {
		_, userRenameErr := rc.Rename(context.Background(), userData.UserName, updateNewUserData.UserName).Result()
		if userRenameErr != nil {
			log.Fatal("Redis Not Renamed Username", userRenameErr)
		}
		userData.UserName = updateNewUserData.UserName
		updatedUser.UserName = updateNewUserData.UserName
	}

	if updateNewUserData.Password != "" {
		hashPWD := md5Encode(updateNewUserData.Password)
		userData.Password = hashPWD
	}

	if updateNewUserData.Name != "" {
		userData.Name = updateNewUserData.Name
		updatedUser.Name = updateNewUserData.Name
	}

	if updateNewUserData.SurName != "" {
		userData.SurName = updateNewUserData.SurName
		updatedUser.SurName = updateNewUserData.SurName
	}
	redisSetJustData(w, userData)
	w.Header().Add("Token", userData.Token)
	responseSuccess(w, updatedUser)
}
