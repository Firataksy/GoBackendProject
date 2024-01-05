package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func registerUser(w http.ResponseWriter) *Sign {

	id := idCreate()
	strID := strconv.Itoa(int(id))
	sn := &Sign{
		ID:       int(id),
		UserName: "player_" + strID,
		Password: "12345",
		Name:     RandStringRunes(5),
		SurName:  RandStringRunes(5),
	}
	sm := SuccessMessage{
		ID:       sn.ID,
		UserName: sn.UserName,
	}
	responseSuccess(w, sm)
	return sn
}

/* func autoScore(w http.ResponseWriter, user *Sign) {
	const SCORE = 3
	user.

}
*/
/* func autoMatch(users []*Sign) {
	var match Match
	for _, user1 := range users {
		for _, user2 := range users {
			match.Score1 = RandIntRunes(3)
			match.Score2 = RandIntRunes(3)
			if match.Score1 > match.Score2 {
				autoScore(w ,user1)

			}
			if match.Score1 < match.Score2 {
				autoScore(w, user2)
			}
			if match.Score1 == match.Score2 {

			}
		}
	}
} */

func simulation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sim Simulation
	err := json.NewDecoder(r.Body).Decode(&sim)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users := make([]*Sign, sim.Count)
	for i := 0; i < sim.Count; i++ {
		ru := registerUser(w)
		users[i] = ru
		/* autoMatch(users) */
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
