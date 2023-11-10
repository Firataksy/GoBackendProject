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
<<<<<<< HEAD
	Status bool   `json:"status"`
	ID     int    `json:"id"`
	UName  string `json:"username"`
	Pwd    string `json:"password"`
}

type Signw struct {
	Status      bool `json:"status"`
	Information struct {
		ID    int    `json:"id"`
		Uname string `json:"uname"`
	}
}

type Loginw struct {
	Status      bool `json:"status"`
	Information struct {
		ID    int    `json:"id"`
		Uname string `json:"uname"`
	}
}

type Listw struct {
	Status      bool `json:"status"`
	Information struct {
		ID    int    `json:"id"`
		Uname string `json:"uname"`
		Name  string `json:"name"`
		SName string `json:"sname"`
	}
=======
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
>>>>>>> c7acc5e1344126a1e92b5362a7a2c9a0dfc2a104
}

type Message struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
