package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/my/repo/internal/utils"
)

func (rc *RedisClient) UpdateUserData(w http.ResponseWriter, r *http.Request) {
	var userData *Sign
	var updatedUser UpdatedUser
	var updateNewUserData UpdateNewUserData

	headerUserID := r.Header.Get("userid")
	er := json.NewDecoder(r.Body).Decode(&updateNewUserData)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	check, _ := rc.Client.Get(context.Background(), "user:"+headerUserID).Result()
	checkUser, _ := rc.Client.Get(context.Background(), check).Result()

	if checkUser == "" {
		ResponseFail(w, "User not found")
		return
	}

	json.Unmarshal([]byte(checkUser), &userData)
	json.Unmarshal([]byte(checkUser), &updatedUser)

	if updateNewUserData.UserName == userData.UserName {
		ResponseFail(w, "You already use this username")
		return
	}

	checkUserID, _ := rc.Client.Get(context.Background(), updateNewUserData.UserName).Result()

	if checkUserID != "" {
		ResponseFail(w, "Username is used")
		return
	}

	if updateNewUserData.UserName != "" {
		_, userRenameErr := rc.Client.Rename(context.Background(), userData.UserName, updateNewUserData.UserName).Result()
		if userRenameErr != nil {
			log.Fatal("Redis Not Renamed Username", userRenameErr)
		}
		rc.Client.Del(context.Background(), "userID:"+userData.UserName)
		userData.UserName = updateNewUserData.UserName
		updatedUser.UserName = updateNewUserData.UserName

	}

	if updateNewUserData.Password != "" {
		hashPWD := utils.Md5Encode(updateNewUserData.Password)
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

	rc.redisSetDataAndID(w, userData)
	ResponseSuccess(w, updatedUser)
}
