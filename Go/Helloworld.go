package main

/*import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dataname string
var datapassword string

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Hesabınız Yok İse Kayıt Olmak İçin 1 Yazınız Var İse 2 Yazınız: ")
	asd, _ := reader.ReadString('\n')
	asd = strings.TrimSpace(asd)

	if asd == "1" {
		register()
	} else if asd == "2" {
		login()
	} else {
		fmt.Print("Yanlış sayı lütfen tekrar deneyiniz")
	}

}

func register() {

	registereader := bufio.NewReader(os.Stdin)
	fmt.Print("Kullanıcı Adı: ")
	registername, _ := registereader.ReadString('\n')
	registername = strings.TrimSpace(registername)

	fmt.Print("Şifre: ")
	registerpw, _ := registereader.ReadString('\n')
	registerpw = strings.TrimSpace(registerpw)

	dataname = registername
	datapassword = registerpw
	login()

}

func login() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Kullanıcı Adı: ")
	username, _ := reader.ReadString('\n')

	username = strings.TrimSpace(username)

	fmt.Print("Şifre: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if username == dataname && password == datapassword {
		fmt.Println("Giriş Başarılı! sayfaya aktarılıyorsunuz.")
		grade()
	} else if username != dataname && password != datapassword {
		fmt.Println("Hesabınız Bulunamadı Lütfen Kayıt Olunuz")

	}
}

func grade() {

	asd := bufio.NewReader(os.Stdin)

	fmt.Print("İsminizi giriniz: ")
	name, _ := asd.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Soyadınızı Giriniz: ")
	surname, _ := asd.ReadString('\n')
	surname = strings.TrimSpace(surname)

	fmt.Print(name, " ", surname, " Lütfen aldığınız notu giriniz: ")
	grade, _ := getgrade()

	var x string
	var error string

	if grade > 100 || grade <= -1 {
		x = "Lütfen Notunuzu Doğru Giriniz!\n"

		error = "Kullanıcı Girişi Sayfasına Gönderiliyorsunuz Lütfen Tekrardan Giriş Yapınız\n"
		fmt.Print(x)
		fmt.Print(error)
		main()

	} else if grade >= 100 {
		fmt.Print(name, " ", surname, " Tam Puan Takdir İle Geçtiniz")
	} else if grade >= 85 {
		fmt.Print(name, " ", surname, " Takdir İle Geçtiniz")
	} else if grade >= 70 {
		fmt.Print(name, " ", surname, " Teşekkür İle Geçtiniz")
	} else if grade >= 50 {
		fmt.Print(name, " ", surname, " Boş Belge İle Geçtiniz")
	} else {
		fmt.Print(name, " ", surname, " Kaldınız")
	}
}

func getgrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return num, nil

}

func getnamesurname() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("İsminizi giriniz: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Soyadınızı Giriniz: ")
	surname, _ := reader.ReadString('\n')
	surname = strings.TrimSpace(surname)
}
*/

//package main

/*import (
	"fmt"
	"net/http"
)

var userregister = []Users{
	{UName: "", Email: "", Pwd: "", PwdConfirm: ""},
}
var userlogin = []User{
	{UrName: "", Pswd: ""},
}

type Users struct {
	UName      string `json:"Name"`
	Email      string `json:"Email"`
	Pwd        string `json:"Password"`
	PwdConfirm string `json:"Passwordc"`
}

type User struct {
	UrName string `json:"Name"`
	Pswd   string `json:"Password"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	var userby Users

	if userby.UName == "" || userby.Email == "" || userby.Pwd == "" || userby.PwdConfirm == "" {
		userregister = append(userregister, Users{})
		fmt.Fprintf(w, "Information cannot be empty")
	} else if userby.Pwd != userby.PwdConfirm {
		fmt.Fprintf(w, "Passwords do not match")
	} else {
		fmt.Fprintf(w, "Succesful signup")
		return
	}

}
func login(w http.ResponseWriter, r *http.Request) {
	var userbyl User
	var userbyr Users
	if userbyr.UName == userbyl.UrName && userbyr.Pwd == userbyl.Pswd {
		fmt.Fprintf(w, "Succesful login")
	} else {
		fmt.Fprintf(w, "Wrong username or password")
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	err := http.ListenAndServe(":9999", mux)
	if err != nil {
		panic(err)
	}
}
*/
/*
------
ıd ıd ıd

var userids = []userid{
	{id: 1},
}

type userid struct {
	id int
}






-----------------
func signup(c *gin.Context) {
	var userby Users
	err := c.Bind(&userby)

	if err != nil || userby.UName != "" || userby.Pwd != "" || userby.Name != "" || userby.SName != "" {
		userregister = append(userregister, userby)
		fmt.Fprint(w, "true", "Succesful signup", userby.Name)
	}
	if userby.UName == "" || userby.Pwd == "" || userby.Name == "" || userby.SName == "" {
		fmt.Fprint(w, "false ", "Information cannot be empty")
		return
	}
	if err == nil && userby.UName != "" && userby.Pwd != "" && userby.Name != "" && userby.SName != "" {
		userregister = append(userregister, userby)
		c.JSON(http.StatusCreated, gin.H{"message": "true"})
		c.JSON(http.StatusCreated, gin.H{"message": "Succesful signup"})
		c.JSON(http.StatusCreated, gin.H{"message": userby.Name})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Student cannot be created "})
		return
	}
}

//func listusers(w http.ResponseWriter, r *http.Request) {
//var list Users
//fmt.Fprint(w, list)
//}

func main() {
	r := gin.Default()
	r.POST("/signup", signup)
	err := r.Run("localhost:9000")
	if err != nil {
		panic(err)
	}
}
--------------
func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Postmandan gelen ID'yi alın
	id := r.URL.Query().Get("id")

	// ID'yi integer'a dönüştür
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Kullanıcı ID'si usersid haritasında var mı kontrol edin
	user, exists := usersid[userID]
	if !exists {
		fmt.Fprint(w, "Durum: Başarısız", "\nMesaj: Kullanıcı bulunamadı")
		return
	}

	// Kullanıcının bilgilerini usersign'dan alın
	userInfo := usersign[userID]

	// Kullanıcı bilgilerini JSON formatında döndürün
	userJSON, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(userJSON)
}
-----------
userid.ID = currentID
		currentID++
		usersid[userid.ID] = userid
*/

//------------------------------
/*func login(w http.ResponseWriter, r *http.Request) {
	var userbyl Userlogin
	var userbys Usersign
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&userbyl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userbys.UName == userbyl.UName && userbys.Pwd == userbyl.Pwd {
		fmt.Fprint(w, `{"success": true, "message": "Successful login"}`")
		fmt.Fprint(w, userbyl, &userbys)
	} else {
		fmt.Fprint(w, `{"success": false, "message": "Wrong username or password"})
		fmt.Fprint(w, userbyl, &userbys)
	}
}*/

/*

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

var (
	users     = make(map[int]User)
	usersLock sync.Mutex
	currentID = 1
)

func main() {
	http.HandleFunc("/signup", SignupHandler)
	http.HandleFunc("/users", UsersHandler)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	usersLock.Lock()
	defer usersLock.Unlock()

	user.ID = currentID
	currentID++
	users[user.ID] = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	})
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	usersLock.Lock()
	defer usersLock.Unlock()

	userList := make([]User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userList)
}
*/
//-----USER SİGN------
/*usersign = append(usersign, userby)

if userby.UName != "" && userby.Pwd != "" && userby.Name != "" && userby.SName != "" {
	fmt.Fprint(w, "success:" ,"True", "message:", "Successful signup", userby.UName, "\n", userby.ID)
} else {
	fmt.Fprint(w, "success:", "False", "message:", " Information cannot be empty")
	return
}
-------------
func login(w http.ResponseWriter, r *http.Request) {
	var userbyl Userlogin
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&userbyl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Kullanıcıyı bulmak için username'e göre bir döngü oluştur
	var foundUser Usersign
	for _, user := range usersign {
		if user.UName == userbyl.UName {
			foundUser = user
			break
		}
	}

	if foundUser.UName == userbyl.UName && foundUser.Pwd == userbyl.Pwd {
		fmt.Fprint(w, "Status: True\nmessage: Successful login\nUsername: ", foundUser.UName)
	} else {
		fmt.Fprint(w, "Status: False\nmessage: Wrong username or password")
	}
}

----------------
func login(w http.ResponseWriter, r *http.Request) {
	var userbyl Userlogin
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&userbyl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Kullanıcı adı ve şifre kontrolü
	for _, user := range usersign {
		if user.UName == userbyl.UName && user.Pwd == userbyl.Pwd {
			// Kullanıcı doğrulandı, başarılı giriş yapın
			var userid Userid
			fmt.Fprint(w, "Status:", "True", "\nmessage:", "Successful login", "\nUserıd: ", userid.ID, "\nUsername: ", userbyl.UName)
			return
		}
	}

	// Kullanıcı adı veya şifre yanlışsa hata mesajı gönderin
	fmt.Fprint(w, "Status:", "False", "\nmessage:", "Wrong username or password")
}

----------------

if userby.UName != "" && userby.Pwd != "" && userby.Name != "" && userby.SName != "" {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":   "True",
			"message":  "User created successfully",
			"userıd":   userby.ID,
			"username": userby.UName,
		})
	} else if userby.UName == "" && userby.Pwd == "" && userby.Name == "" && userby.SName == "" {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "False",
			"message": "Information cannot be empty",
		})
		return
	}
*/

/*
package main

import (
    "fmt"
    "sync"
)

// Kullanıcıları tutmak için kullanılan veri yapısı
type UserRegistry struct {
    users map[string]bool
    mu    sync.Mutex
}

// Yeni bir kullanıcı kaydeder
func (ur *UserRegistry) Register(username string) error {
    ur.mu.Lock()
    defer ur.mu.Unlock()

    if ur.users[username] {
        return fmt.Errorf("Kullanıcı adı zaten kayıtlı: %s", username)
    }

    ur.users[username] = true
    fmt.Printf("Kullanıcı kaydedildi: %s\n", username)
    return nil
}

func main() {
    ur := &UserRegistry{
        users: make(map[string]bool),
    }

    usernames := []string{"user1", "user2", "user1", "user3"}

    for _, username := range usernames {
        if err := ur.Register(username); err != nil {
            fmt.Println(err)
        }
    }
}
-------------------------
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var usersign = []Usersign{
	{UName: "", Pwd: "", Name: "", SName: ""},
}

var userlogin = []Userlogin{
	{UName: "", Pwd: "", Name: "", SName: ""},
}

var (
	usersid   = make(map[int]Userid)
	user      = make(map[string]Usersign)
	currentID = 1
)

type Userid struct {
	ID int `json:"id"`
}

type Usersign struct {
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}

type Userlogin struct {
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}

func namecheck(UName string) bool{
	_,exists := user[UName]
	return !exists
}

func signup(w http.ResponseWriter, r *http.Request) (user string){
	w.Header().Set("Content-Type", "application/json")
	var userby Usersign
	var userl Userlogin
	var id Userid

	err := json.NewDecoder(r.Body).Decode(&userby)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	usersign = append(usersign, userby)
	if namecheck(UName){
		user[Uname] = true

	}else if userby.UName == userl.UName {
		fmt.Fprint(w, "Status: False", "Message: Username is used please try again")
		return
	} else if userby.UName != "" && userby.Pwd != "" && userby.Name != "" && userby.SName != "" {
		id.ID = currentID
		currentID++
		usersid[id.ID] = id
		fmt.Fprint(w, "Status: True", "\nMessage: Successful signup", "\nUserıd: ", id.ID, "\nUsername: ", userby.UName)
		userl.Name = userby.Name
		userl.SName = userby.SName
		return
	} else {
		fmt.Fprint(w, "Status: False", "\nMessage: Information cannot be empty")
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userbyl Userlogin
	var userid Userid

	err := json.NewDecoder(r.Body).Decode(&userbyl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userlogin = append(userlogin, userbyl)

	for id, user := range usersign {
		if user.UName == userbyl.UName && user.Pwd == userbyl.Pwd {
			userid.ID = id
			fmt.Fprint(w, "Status: True", "\nMessage: Successful login", "\nUserıd: ", userid.ID, "\nUsername: ", userbyl.UName)
			return
		}
	}
	fmt.Fprint(w, "Status: False", "\nMessage: Wrong username or password")
}

func getusers(w http.ResponseWriter, r *http.Request) {
	var userid Userid
	var status = false
	fmt.Fprint(w, "----Informations----\n")

	for id, user := range userlogin {
		if id != 1 && id != 0 {
			userid.ID = id
			fmt.Fprint(w, "Status: True", "\nUserId: ", userid.ID, "\nUsername: ", user.UName, "\nUserFirstName: ", user.Name, "\nUserLastName: ", user.SName, "\n----------------------", "\n")
			status = true
		} else if status ==  {
			fmt.Fprint(w, "User not found")
		}
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/list", getusers)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
------GETUSER------
func getusers(w http.ResponseWriter, r *http.Request) {
	var userid Userid
	fmt.Fprint(w, "----Informations----\n")

	for id, user := range usersign {
		if id != 0 {
			userid.ID = id
			fmt.Fprint(w, "Status: True", "\nUserId: ", userid.ID, "\nUsername: ", user.UName, "\nUserFirstName: ", user.Name, "\nUserLastName: ", user.SName, "\n----------------------", "\n")
		}

	}
	return
}

-------Signup---
func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userby Usersign
	var userid Userid

	err := json.NewDecoder(r.Body).Decode(&userby)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	usersign = append(usersign, userby)
	_, information := user[userby.UName]
	fmt.Fprint(w, information)

	if information != false {
		fmt.Fprint(w, "Username is used")
		return
	} else if information != true && userby.UName != "" && userby.Pwd != "" && userby.Name != "" && userby.SName != "" {
		userid.ID = currentID
		currentID++
		usersid[userid.ID] = userid
		user[userby.UName] = userby
		user[userby.Pwd] = userby
		fmt.Fprint(w, "Status: True", "\nMessage: Successful signup", "\nUserıd: ", userid.ID, "\nUsername: ", userby.UName)
		return
	} else {
		fmt.Fprint(w, "Status: False", "\nMessage: Information cannot be empty")
		return
	}
}

----------------------
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var usersign = []Usersign{
	{UName: "", Pwd: "", Name: "", SName: ""},
}

var userlogin = []Userlogin{
	{UName: "", Pwd: ""},
}

var userid = []Userbyid{
	{ID: 0},
}

var (
	user      = make(map[string]Usersign)
	userl     = make(map[string]Userlogin)
	useri     = make(map[int]Userbyid)
	currentID = 1
)

type Usersign struct {
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}

type Userlogin struct {
	ID    int    `json:"id"`
	UName string `json:"username"`
	Pwd   string `json:"password"`
	Name  string `json:"name"`
	SName string `json:"sname"`
}

type Userbyid struct {
	ID int `json:"id"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usersignup Usersign
	var userID Userbyid

	err := json.NewDecoder(r.Body).Decode(&usersignup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersign = append(usersign, usersignup)
	_, information := user[usersignup.UName]
	fmt.Fprint(w, information)

	if information != false {
		fmt.Fprint(w, "Username is used")
		return
	} else if information != true && usersignup.UName != "" && usersignup.Pwd != "" && usersignup.Name != "" && usersignup.SName != "" {
		userID.ID = currentID
		currentID++
		useri[userID.ID] = userID
		user[usersignup.UName] = usersignup
		fmt.Fprint(w, "Status: True", "\nMessage: Successful signup", "\nUserıd: ", userID.ID, "\nUsername: ", usersignup.UName, "\n")

	} else {
		fmt.Fprint(w, "Status: False", "\nMessage: Information cannot be empty")
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userbylogin Userlogin
	var userid Userbyid

	err := json.NewDecoder(r.Body).Decode(&userbylogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := useri[userid.ID]
	userl[userbylogin.UName] = userbylogin

	for _, user := range user {
		if user.UName == userbylogin.UName && user.Pwd == user.Pwd {
			fmt.Fprint(w, "Status: True", "\nMessage: Successful login", "\nUserıd: ", id, "\nUsername: ", userbylogin.UName)
			return
		} else {
			fmt.Fprint(w, "Status: False", "\nMessage: Wrong username or password")
			fmt.Fprint(w, user, userbylogin)
			return
		}
	}
}

func getusers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "----Informations----\n")
	for id, user := range userlogin {
		fmt.Fprint(w, "Status: True", "\nUserId: ", id, "\nUsername: ", user.UName, "\nUserFirstName: ", user.Name, "\nUserLastName: ", user.SName, "\n----------------------", "\n")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/list", getusers)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}

-------------- User datası çekme --------------
for _, user := range user {
		fmt.Println(user.UName, " ", userbylogin.UName, " ", userbylogin.Pwd, " ", user.Pwd)
		if user.UName == userbylogin.UName && user.Pwd == userbylogin.Pwd {
			fmt.Fprint(w, "Status: True", "\nMessage: Successful login", "\nUserıd: ", user.ID, "\nUsername: ", userbylogin.UName)

		} else {
			fmt.Fprint(w, "Status: False", "\nMessage: Wrong username or password")
			return
		}
	}
------------- USER LİSTELEME -----------------
	func getUserByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // URL'den kullanıcı ID'sini alın
    userID := r.URL.Query().Get("id")

    // Kullanıcı ID'sini tamsayıya çevirin
    id, err := strconv.Atoi(userID)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    // Kullanıcıyı ID'ye göre bulun
    user, found := userByID(id)
    if !found {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Kullanıcı bilgilerini JSON formatına çevirip yanıt verin
    userJSON, err := json.Marshal(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write(userJSON)
}

// Kullanıcıyı ID'ye göre bulan yardımcı bir fonksiyon
func userByID(id int) (Sign, bool) {
    for _, u := range usersign {
        if u.ID == id {
            return u, true
        }
    }
    return Sign{}, false
}

*/
