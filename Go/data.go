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
	Status bool `json:"status"`
	Data   interface{}
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Pwd      string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
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
	userid1 int `json:"userid1"`
	userid2 int `json:"userid2"`
	score1  int `json:"score1"`
	score2  int `json:"score2"`
}
