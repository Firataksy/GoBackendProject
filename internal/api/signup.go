package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/my/repo/internal/utils"
)

func (rc *RedisClient) SignUp(w http.ResponseWriter, r *http.Request) {
	var userSignUp *Sign
	var user User

	err := json.NewDecoder(r.Body).Decode(&userSignUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UserID, _ := rc.Client.Get(context.Background(), "userID:"+userSignUp.UserName).Result()
	check, _ := rc.Client.Get(context.Background(), "player_"+UserID).Result()
	json.Unmarshal([]byte(check), &user)
	if user.UserName == userSignUp.UserName {
		responseFail(w, "Username is used")
		return
	}

	if userSignUp.UserName != "" && userSignUp.Password != "" && userSignUp.Name != "" && userSignUp.SurName != "" {

		userSignUp.Password = utils.Md5Encode(userSignUp.Password)
		id := rc.IDCreate()
		userSignUp.ID = int(id)

		token := utils.GenerateToken()
		userSignUp.Token = token
		rc.redisSetToken(userSignUp)

		sm := TokenUsername{
			Token:    userSignUp.Token,
			UserName: userSignUp.UserName,
		}

		rc.redisSetJustData(w, userSignUp)
		rc.redisSetUserNameAndID(w, userSignUp.UserName, int(id))

		responseSuccess(w, sm)
		return
	}
	responseFail(w, "Information cannot be empty")
}

func (rc *RedisClient) redisSetToken(sign *Sign) {
	_, err := rc.Client.Set(context.Background(), "token:"+sign.Token, sign.ID, 0).Result()
	if err != nil {
		log.Fatal("Redis Set Token err:", err)
		return
	}
}
