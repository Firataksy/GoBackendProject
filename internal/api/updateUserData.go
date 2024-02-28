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
	var updatedNewUserData UpdateNewUserData

	headerUserID := r.Header.Get("userid")
	er := json.NewDecoder(r.Body).Decode(&updatedNewUserData)
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

	if updatedNewUserData.UserName == userData.UserName {
		ResponseFail(w, "You already use this username")
		return
	}

	checkUserID, _ := rc.Client.Get(context.Background(), updatedNewUserData.UserName).Result()

	if checkUserID != "" {
		ResponseFail(w, "Username is used")
		return
	}

	if updatedNewUserData.UserName != "" {
		_, userRenameErr := rc.Client.Rename(context.Background(), userData.UserName, updatedNewUserData.UserName).Result()
		if userRenameErr != nil {
			log.Fatal("Redis Not Renamed Username", userRenameErr)
		}
		rc.Client.Del(context.Background(), "userID:"+userData.UserName)
		userData.UserName = updatedNewUserData.UserName
		updatedUser.UserName = updatedNewUserData.UserName

	}

	if updatedNewUserData.Password != "" {
		hashPWD := utils.Md5Encode(updatedNewUserData.Password)
		userData.Password = hashPWD
	}

	if updatedNewUserData.Name != "" {
		userData.Name = updatedNewUserData.Name
		updatedUser.Name = updatedNewUserData.Name
	}

	if updatedNewUserData.SurName != "" {
		userData.SurName = updatedNewUserData.SurName
		updatedUser.SurName = updatedNewUserData.SurName
	}

	rc.redisSetDataAndID(userData)
	ResponseSuccess(w, updatedUser)
}
