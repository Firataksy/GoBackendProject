package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signUp)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/data", getUserData)
	mux.HandleFunc("/updateuser", updateUserData)
	mux.HandleFunc("/match", match)
	mux.HandleFunc("/leaderboard", listLeaderBoard)
	mux.HandleFunc("/simulation", simulation)
	fmt.Println("http listen started")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		fmt.Println(err)
	}
}

var rc *redis.Client

func redisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}

func init() {
	rc = redisConnect()
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func idCreate() int64 {
	userIncrID, _ := rc.Incr(context.Background(), "userIncrId").Result()
	return userIncrID
}

func jsonConvert(w http.ResponseWriter, input interface{}) []byte {
	Json, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Json Error", http.StatusInternalServerError)
		return nil
	}
	return Json
}

func responseSuccess(w http.ResponseWriter, input interface{}) {
	w.Header().Add("Content-Type", "application/json")
	rp := Response{
		Status: true,
		Data:   input,
	}

	response := jsonConvert(w, rp)
	w.Write(response)
}

func responseError(w http.ResponseWriter, input string) {
	ms := &FailMessage{
		Status:  false,
		Message: input,
	}

	response := jsonConvert(w, ms)
	w.Write(response)
}

func redisSetJustData(w http.ResponseWriter, data *Sign) {
	jsonData := jsonConvert(w, data)
	strID := strconv.Itoa(data.ID)
	_, er := rc.Set(context.Background(), "player_"+strID, jsonData, 0).Result()
	if er != nil {
		log.Fatal("Set User data err: ", er)
	}
}

func redisSetJustID(username string, id int) {
	_, er := rc.Set(context.Background(), "userID:"+username, id, 0).Result()
	if er != nil {
		log.Fatal("Set User ID err: ", er)
	}
}

func redisSetDataAndID(w http.ResponseWriter, data *Sign) {
	redisSetJustData(w, data)
	redisSetJustID(data.UserName, data.ID)
}

func redisSetLeaderBoard(user *Sign) {

	z := &redis.Z{
		Score:  float64(user.Score),
		Member: user.ID,
	}

	rc.ZAdd(context.Background(), "leaderboard", *z).Result()
}

// func redisGetAllLeaderBoardData() []Sign {
// 	var user User
// 	allDataList, err := rc.ZRevRangeWithScores(context.Background(), "leaderboard", 0, -1).Result()
// 	if err != nil {
// 		log.Fatal("ERR list leaderboard", err)
// 		return nil
// 	}

// 	leaderBoardSlice := make([]Sign, len(allDataList))
// 	for i, data := range allDataList {
// 		s, _ := rc.Get(context.Background(), "player_"+data.Member.(string)).Result()
// 		err := json.Unmarshal([]byte(s), &user)
// 		if err != nil {
// 			log.Fatal("Unmarshal err:", err)
// 		}

// 		leaderBoardSlice[i] = Sign{
// 			ID:       user.ID,
// 			Score:    user.Score,
// 			UserName: user.UserName,
// 		}
// 	}
// 	return leaderBoardSlice
// }

func generateToken() string {
	token := make([]byte, 8)
	rand.Read(token)
	return fmt.Sprintf("%x", token)
}

func redisSetToken(sign *Sign) {
	_, err := rc.Set(context.Background(), "token:"+sign.Token, sign.ID, 0).Result()
	if err != nil {
		log.Fatal("Redis Set Token err:", err)
	}
}
