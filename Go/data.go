package main

var (
	user      = make(map[string]Sign)
	userl     = make(map[string]Login)
	currentID = 0
)

type Sign struct {
	Status   bool   `json:"status"`
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Pwd      string `json:"password"`
	Name     string `json:"name"`
	SurName  string `json:"surname"`
}

type Signw struct {
	Status bool `json:"status"`
	Data   Sign
}

type Login struct {
	Status   bool   `json:"status"`
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Listw struct {
	Status bool `json:"status"`
	Data   struct {
		ID       int    `json:"id"`
		UserName string `json:"username"`
		Name     string `json:"name"`
		SurName  string `json:"surname"`
	}
}

func Signr(status bool, ID int, UserName string) Sign {
	return Sign{
		Status:   status,
		ID:       ID,
		UserName: UserName,
	}
}

func Loginr(status bool, ID int, UserName string) Login {
	return Login{
		Status:   status,
		ID:       ID,
		UserName: UserName,
	}
}

func Statustrue() Message {
	var mes Message
	mes.Status = true
	return mes
}

func Statusfalse() Message {
	var mes Message
	mes.Status = false
	mes.Message = "Error, try again"
	return mes
}

type Message struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
