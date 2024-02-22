package api

import "net/http"

func (rc *RedisClient) win(w http.ResponseWriter, user *Sign) {
	user.Score += 3

	rc.redisSetJustData(w, user)
	rc.redisSetLeaderBoard(user)

}

func (rc *RedisClient) draw(w http.ResponseWriter, user1 *Sign, user2 *Sign) {
	user1.Score += 1
	user2.Score += 1

	rc.redisSetJustData(w, user1)
	rc.redisSetLeaderBoard(user1)

	rc.redisSetJustData(w, user2)
	rc.redisSetLeaderBoard(user2)

}
