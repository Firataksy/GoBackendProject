package main

var usersign = []Sign{
	{UName: "", Pwd: "", Name: "", SName: ""},
}

var userlogin = []Login{
	{UName: "", Pwd: ""},
}

var (
	user      = make(map[string]Sign)
	userl     = make(map[string]Login)
	currentID = 0
)

type Sign struct {
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}

type Login struct {
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
}

type Message struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
