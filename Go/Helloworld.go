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
