package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	var userSignUp *Sign
	var user User

	err := json.NewDecoder(r.Body).Decode(&userSignUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UserID, _ := rc.Get(context.Background(), "userID:"+userSignUp.UserName).Result()
	check, _ := rc.Get(context.Background(), "player_"+UserID).Result()
	json.Unmarshal([]byte(check), &user)
	if user.UserName == userSignUp.UserName {
		responseFail(w, "Username is used")
		return
	}

	if userSignUp.UserName != "" && userSignUp.Password != "" && userSignUp.Name != "" && userSignUp.SurName != "" {

		userSignUp.Password = md5Encode(userSignUp.Password)
		id := idCreate()
		userSignUp.ID = int(id)

		token := generateToken()
		userSignUp.Token = token
		redisSetToken(userSignUp)

		sm := TokenUsername{
			Token:    userSignUp.Token,
			UserName: userSignUp.UserName,
		}

		redisSetJustData(w, userSignUp)
		redisSetUserNameAndID(w, userSignUp.UserName, int(id))

		responseSuccess(w, sm)
		return
	}
	responseFail(w, "Information cannot be empty")
}
