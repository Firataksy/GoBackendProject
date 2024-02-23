package api

func (rc *RedisClient) win(user *Sign) {
	user.Score += 3

	rc.redisSetJustData(user)
	rc.redisSetLeaderBoard(user)

}

func (rc *RedisClient) draw(user1 *Sign, user2 *Sign) {
	user1.Score += 1
	user2.Score += 1

	rc.redisSetJustData(user1)
	rc.redisSetLeaderBoard(user1)

	rc.redisSetJustData(user2)
	rc.redisSetLeaderBoard(user2)

}
