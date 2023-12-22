package main

import (
	"encoding/json"
	"net/http"
)

func match(w http.ResponseWriter, r *http.Request) {
	var match Match

	er := json.NewDecoder(r.Body).Decode(&match)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}
	//struserid1 := strconv.Itoa(match.userid1)
	//struserid2 := strconv.Itoa(match.userid2)
	//checkuserid1, _ := rc.Get(context.Background(), "user:"+struserid1).Result()
	//checkuserid2, _ := rc.Get(context.Background(), "user:"+struserid2).Result()

}
