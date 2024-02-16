package main

type Sign struct {
	Token    string `json:"token"`
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
	Score    int    `json:"score"`
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

type TokenUsername struct {
	Token    string `json:"token"`
	UserName string `json:"username"`
}

type FailMessage struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessMessage struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
}

type Success struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type User struct {
	Token    string `json:"token"`
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Pwd      string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
	Score    int    `json:"score"`
}

type UpdatedUser struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
}

type UpdateNewUserData struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
}

type Match struct {
	UserID1 int `json:"userid1"`
	UserID2 int `json:"userid2"`
	Score1  int `json:"score1"`
	Score2  int `json:"score2"`
}

type LeaderBoard struct {
	Page  int `json:"page"`
	Count int `json:"count"`
}

type UserLeaderBoard struct {
	UserName string `json:"username"`
	UserID   int    `json:"id"`
	Score    int    `json:"score"`
}

type Simulation struct {
	Count int `json:"count"`
}

type FriendRequestList struct {
	ID       int     `json:"ID"`
	UserName string  `json:"username"`
	Date     float64 `json:"date"`
}

type FriendList struct {
	ID       int    `json:"ID"`
	UserName string `json:"username"`
}

type UserSearchID struct {
	ID int `json:"ID"`
}

type AcceptReject struct {
	Status string `json:"status"`
	ID     int    `json:"ID"`
}
