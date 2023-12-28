package main

type Sign struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
}

type Login struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type SuccessData struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
}

type SuccessMessage struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

type FailMessage struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Pwd      string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
	Puan     int    `json:"puan"`
}

type RedisControlData struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UpdateNewUserData struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
}

type UserData struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
}

type Match struct {
	Userid1 int `json:"userid1"`
	Userid2 int `json:"userid2"`
	Score1  int `json:"score1"`
	Score2  int `json:"score2"`
}

type User1 struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Pwd      string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
	Puan     int    `json:"puan"`
}

type User2 struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Pwd      string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
	Puan     int    `json:"puan"`
}

type LeaderBoard struct {
	Page  int `json:"page"`
	Count int `json:"count"`
}

type UserLeaderBoard struct {
	UserName string `json:"username"`
	ID       int    `json:"id"`
	Puan     int    `json:"puan"`
}
