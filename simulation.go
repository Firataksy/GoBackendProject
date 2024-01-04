package main

import "strconv"

func registerUser() *Sign {
	id := idCreate()
	strID := strconv.Itoa(int(id))
	sn := &Sign{
		UserName: "player_" + strID,
		Password: "12345",
		Name:     RandStringRunes(5),
		SurName:  RandStringRunes(5),
	}
	return sn
}
