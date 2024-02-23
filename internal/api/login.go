package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/my/repo/internal/utils"
)

func (rc *RedisClient) Login(w http.ResponseWriter, r *http.Request) {
	var userLogin Login
	var user User

	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	val, err := rc.Client.Get(context.Background(), userLogin.UserName).Result()
	if err != nil {
		log.Fatal("login get user data err :", err)
		return
	}

	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		log.Fatal("login unmarshal err :", err)
		return
	}

	userLogin.Password = utils.Md5Encode(userLogin.Password)
	if user.Pwd == userLogin.Password {

		sm := TokenUsername{
			Token:    user.Token,
			UserName: userLogin.UserName,
		}

		ResponseSuccess(w, sm)
		return
	}

	ResponseFail(w, "Wrong username or password")
}
