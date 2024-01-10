package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	var userSignUp Sign
	var user User

	err := json.NewDecoder(r.Body).Decode(&userSignUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UserID, _ := rc.Get(context.Background(), "userID:"+userSignUp.UserName).Result()
	check, _ := rc.Get(context.Background(), "user:player_"+UserID).Result()
	json.Unmarshal([]byte(check), &user)
	if user.UserName == userSignUp.UserName {
		responseError(w, "Username is used")
		return
	}

	if userSignUp.UserName != "" && userSignUp.Password != "" && userSignUp.Name != "" && userSignUp.SurName != "" {

		userSignUp.Password = md5Encode(userSignUp.Password)
		id := idCreate()
		sm := SuccessMessage{
			ID:       int(id),
			UserName: userSignUp.UserName,
		}

		userSignUp.ID = int(id)
		redisAllData := jsonConvert(w, userSignUp)
		stringID := strconv.Itoa(int(id))

		_, userInfoErr := rc.Set(context.Background(), "user:"+stringID, redisAllData, 0).Result()
		if userInfoErr != nil {
			log.Fatal("Redis set user data sign error:", userInfoErr)
			return
		}

		_, userErr := rc.Set(context.Background(), "user:"+userSignUp.UserName, int(id), 0).Result()
		if userErr != nil {
			log.Fatal("Redis set user id sign error:", userErr)
			return
		}

		responseSuccess(w, sm)
		return
	}
	responseError(w, "Information cannot be empty")
}
