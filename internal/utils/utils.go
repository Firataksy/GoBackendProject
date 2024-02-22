package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func Md5Encode(input string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(input))
	md5 := hash.Sum(nil)
	return fmt.Sprintf("%x", md5)
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func JsonConvert(w http.ResponseWriter, input interface{}) []byte {
	Json, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Json Error", http.StatusInternalServerError)
		return nil
	}
	return Json
}

func GenerateToken() string {
	token := make([]byte, 8)
	rand.Read(token)
	return fmt.Sprintf("%x", token)
}
