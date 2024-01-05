package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func registerUser() *Sign {

	id := idCreate()
	strID := strconv.Itoa(int(id))
	sn := &Sign{
		ID:       int(id),
		UserName: "player_" + strID,
		Password: "12345",
		Name:     RandStringRunes(5),
		SurName:  RandStringRunes(5),
	}
	return sn
}

func simulation(w http.ResponseWriter, r *http.Request) {
	var sim Simulation
	err := json.NewDecoder(r.Body).Decode(&sim)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i := 0; i < sim.Count; i++ {
		ru := registerUser()
		_, er := rc.Set(context.Background(), "ID:"+ru.UserName, ru.ID, 0).Result()
		if er != nil {
			log.Fatal("Set User ID err: ", er)
		}

		id := strconv.Itoa(ru.ID)
		hashPwd := md5Encode(ru.Password)
		ru.Password = hashPwd
		ruJson := jsonConvert(w, ru)
		_, error := rc.Set(context.Background(), "player_"+id, ruJson, 0).Result()
		if error != nil {
			log.Fatal("Set User Data err: ", error)
		}
	}
}
