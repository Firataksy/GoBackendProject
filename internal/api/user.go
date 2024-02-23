package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/my/repo/internal/utils"
)

func (rc *RedisClient) RegisterUser() *Sign {
	token := utils.GenerateToken()
	id := rc.IDCreate()
	strID := strconv.Itoa(int(id))
	sn := &Sign{
		Token:    token,
		ID:       int(id),
		UserName: "player_" + strID,
		Password: "12345",
		Name:     utils.RandStringRunes(5),
		SurName:  utils.RandStringRunes(5),
	}

	hashPwd := utils.Md5Encode(sn.Password)
	sn.Password = hashPwd
	rc.redisSetToken(sn)
	return sn
}

func ResponseSuccessMessage(w http.ResponseWriter, input interface{}) {
	w.Header().Add("Content-Type", "application/json")
	rp := SuccessMessage{
		Status:  true,
		Message: input,
	}

	response := utils.JsonConvert(rp)
	w.Write(response)
}

func ResponseSuccess(w http.ResponseWriter, input interface{}) {
	w.Header().Add("Content-Type", "application/json")
	rp := Success{
		Status: true,
		Data:   input,
	}

	response := utils.JsonConvert(rp)
	w.Write(response)
}

func ResponseFail(w http.ResponseWriter, input string) {
	w.Header().Add("Content-Type", "application/json")
	ms := FailMessage{
		Status:  false,
		Message: input,
	}

	response := utils.JsonConvert(ms)
	w.Write(response)
}

func (rc *RedisClient) IDCreate() int64 {
	userIncrID, _ := rc.Client.Incr(context.Background(), "userIncrId").Result()
	return userIncrID
}
