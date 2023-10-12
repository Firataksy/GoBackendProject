/*import "fmt"

func main() {

	/*var (
		firstname      = "Fırat"
		age       int  = 17

		weight float32 = 90.5
		height int     = 170
	)
	//var firstname, age, weight, height = "Fırat", 17, 90.5, 170

	firstname, age, weight, height := "Fırat", 17, 90.5, 170

	fmt.Println(firstname)
	fmt.Println(age)
	fmt.Println(weight)
	fmt.Println(height)

	studentname, grade, isPassed := "Fırat", 77, true

	fmt.Println(studentname)
	fmt.Println(grade)
	fmt.Println(isPassed)
}*/
//----------------------------------------------------------------------------
/*package main

import "fmt"

func main() {
	var x, y int

	for x = 1; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}
		if y > (x / y) {
			fmt.Printf("%d Asal Sayıdır +++++ \n", x)
			fmt.Println("---------------------------")
		} else {
			fmt.Printf("%d Asal Sayı Değildir !!!!! \n", x)
			fmt.Println("---------------------------")
		}
	}
}*/
//----------------------------------------------------------------------------
/*package main

import "fmt"

func main() {
	x := 0

	for x < 10 {
		fmt.Println(x)
		x++
	}
}*/
//----------------------------------------------------------------------------
/*package main

import "fmt"

func main() {
	switch x := 25; {
	case x < 20:
		fmt.Printf("%d küçüktür 20 \n", x)
		fallthrough

	case x < 50:
		fmt.Printf("%d küçüktür 50 \n", x)
		fallthrough

	case x < 100:
		fmt.Printf("%d küçüktür 100 \n", x)
		fallthrough

	case x < 200:
		fmt.Printf("%d küçüktür 200 \n", x)
	}
}*/
//----------------------------------------------------------------------------
/*package main

import "fmt"

func main() {

	var x, y int

	for x = 1; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}

		if x%2 == 0 {
			fmt.Printf("%d Çift!!\n", x)
			fmt.Println("--------------------------------------------------------")

		} else {
			fmt.Printf("%d Tek!\n", x)
			fmt.Println("--------------------------------------------------------")
		}
	}

}*/
//----------------------------------------------------------------------------
/*package main

import "fmt"

func main() {
	x, y := 10, 4
	sum1, dif1, prod1 := calculation(x, y)
	fmt.Println("Toplam: ", sum1)
	fmt.Println("Fark: ", dif1)
	fmt.Println("Çarpım: ", prod1)
}

func calculation(num1, num2 int) (int, int, int) {
	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2

	return sum, dif, prod

}*/
//----------------------------------------------------------------------------
/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Lütfen Aldığınız Notu giriniz: ")
	grade, _ := getgrade()

	var result string

	if grade >= 50 {
		result = "Geçtin"
	} else {
		result = "Kaldın"
	}

	fmt.Println(result)

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

}*/
//----------------------------------------------------------------------------

/*
import (

	"bufio"
	"fmt"
	"os"
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
			fmt.Println("Giriş Başarılı!")
		} else {
			fmt.Println("Kullanıcı Adı veya Şifre Hatalı Lütfen Tekrar Deneyiniz!")
		}
	}
*/
//---------- Pointers ----------

/*import "fmt"

func main() {
	a := 8
	ekle(&a)
	fmt.Println(a) //Çıktımız: 13
}
func ekle(v *int) {
	*v += 5
}*/

//---------- Struct ----------

/*import "fmt"

type kişi struct {
	isim    string
	soyİsim string
	yaş     int
}

func main() {

	kişi1 := kişi{"Kaan", "Kuşcu", 23}

	fmt.Println(kişi1)

}*/
//---------- Struct (method)----------
/*import "fmt"

type insan struct {
	isim string
	yaş  int
}

func (i insan) tanıt() {
	fmt.Printf("Merhaba, Ben %s. %d yaşındayım.", i.isim, i.yaş)
}
func main() {
	kişi := insan{"Kaan", 23}
	kişi.tanıt()
}*/
//---------- Slices ---------
/*import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	fmt.Println(a)
	a = append(a, "Ali")
	a = append(a, "Veli")
	fmt.Println(a)
}*/

/*import "fmt"

var isimler = []string{"Ali", "Veli", "Hasan", "Ahmet", "Mehmet"}

func main() {
	for a, b := range isimler {
		fmt.Printf("%d. indeks = %s\n", a, b)
	}
}*/
/*mport "fmt"

type kişi struct {
	isim    string
	soyisim string
	yaş     int
}

func main() {
	kişi1 := kişi{"Fırat", "Aksoy", 17}

	fmt.Println(kişi1)
}*/
//---------- Struct ----------
/*import (
	"fmt"
)

type kisi struct {
	ad     string
	soyad  string
	yas    int
	email  string
	numara int
}

func main() {
	k := kisi{"Fırat", "Aksoy", 17, "asdqwegas", 1243124542331}
	fmt.Println("Adınız: ", k.ad)
	fmt.Println("Soyadınız: ", k.soyad)
	fmt.Println("Yaşınız: ", k.yas)
	fmt.Println("E-Mail: ", k.email)
	fmt.Println("Numaranız: ", k.numara)
}*/
//---------- Range ---------
/*
import "fmt"

var isimler = []string{"Ali", "Veli", "Hasan", "Ahmet", "Mehmet"}

func main() {
	for a, b := range isimler {
		fmt.Printf("%d. indeks = %s\n", a, b)
	}
}*/

/*import (
	"fmt"
)

type dikdortgen struct {
	a, b float64
}

func (d dikdortgen) alan() float64 {
	return d.a * d.b
}

func (d dikdortgen) cevre() float64 {
	return 2 * (d.a + d.b)
}

type sekil interface {
	alan() float64
	cevre() float64
}

func interfacefunc(i sekil) {
	fmt.Println(i)
	fmt.Println(i.alan())
	fmt.Println(i.cevre())
	fmt.Printf("%T", i)
	fmt.Println()

}

func main() {
	r1 := dikdortgen{3, 8}
	fmt.Println("Alan: ", r1.alan())
	fmt.Println("Cevre: ", r1.cevre())

	fmt.Println()
	interfacefunc(r1)

}*/
////---------- Generics ---------
/*import "fmt"

func main() {
	var sayı1 int = 5
	var sayı2 float64 = 5.3
	fmt.Println(arttır(sayı1))
	fmt.Println(arttır(sayı2))
}

func arttır[n int | float64](sayı n) n {
	return sayı + 1
}*/

//---------- Interface deneme ---------

/*
import (

	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

)

	type dikhesap struct {
		a, b float64
	}

	type karehesap struct {
		a float64
	}

	type ucgencevrehesap struct {
		a, b, c float64
	}

	type ucgenalanhesap struct {
		a, b float64
	}

	func (d dikhesap) alan() float64 {
		return d.a * d.b
	}

	func (d dikhesap) cevre() float64 {
		return 2 * (d.a + d.b)
	}

	func (k karehesap) alan() float64 {
		return k.a * k.a
	}

	func (k karehesap) cevre() float64 {
		return k.a * 4
	}

	func (u ucgencevrehesap) ucgencevre() float64 {
		return u.a + u.b + u.c
	}

	func (uc ucgenalanhesap) ucgenalan() float64 {
		ucgencevre1()
		return (uc.a * uc.b) / 2

}

	type sekil interface {
		alan() float64
		cevre() float64
		ucgenalan() float64
	}

	func name() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Başlamadan önce adınızı giriniz: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
	}

	func dikdortgen() {
		dikkare := bufio.NewReader(os.Stdin)

		fmt.Print("Uzun Kenarı Giriniz: ")
		num1, _ := dikkare.ReadString('\n')
		num1 = strings.TrimSpace(num1)

		fmt.Print("Kısa Kenarı Giriniz: ")
		num2, _ := dikkare.ReadString('\n')
		num2 = strings.TrimSpace(num2)

		sayi1, _ := strconv.ParseFloat(num1, 64)
		sayi2, _ := strconv.ParseFloat(num2, 64)

		r1 := dikhesap{sayi1, sayi2}
		fmt.Println("Alan: ", r1.alan())
		fmt.Println("Cevre: ", r1.cevre())
		cıkgir()

}

	func kare() {
		kare := bufio.NewReader(os.Stdin)

		fmt.Print(" Bir Kenarını giriniz: ")
		num1, _ := kare.ReadString('\n')
		num1 = strings.TrimSpace(num1)

		sayi1, _ := strconv.ParseFloat(num1, 64)

		r1 := karehesap{sayi1}
		fmt.Println("Alan: ", r1.alan())
		fmt.Println("Cevre: ", r1.cevre())
		cıkgir()

}

	func ucgencevre1() {
		fmt.Print(" Üçgenin çevresini hesaplamak için uzunlukları giriniz")
		asd := bufio.NewReader(os.Stdin)

		fmt.Println("İlk uzunluk: ")
		num1, _ := asd.ReadString('\n')
		num1 = strings.TrimSpace(num1)

		fmt.Print("İkinci uzunluk: ")
		num2, _ := asd.ReadString('\n')
		num2 = strings.TrimSpace(num2)

		fmt.Print("Üçüncü uzunluk: ")
		num3, _ := asd.ReadString('\n')
		num3 = strings.TrimSpace(num3)

		sayi1, _ := strconv.ParseFloat(num1, 64)
		sayi2, _ := strconv.ParseFloat(num2, 64)
		sayi3, _ := strconv.ParseFloat(num3, 64)

		r1 := ucgencevrehesap{sayi1, sayi2, sayi3}
		fmt.Println("Çevre: ", r1.ucgencevre())
		cıkgir()

}

	func ucgenalan1() {
		fmt.Println(" Uçgenin alanını hesaplamak için uzunlukları giriniz")
		ucgen := bufio.NewReader(os.Stdin)

		fmt.Print("Üçgeninin taban uzunluğunu giriniz: ")
		num1, _ := ucgen.ReadString('\n')
		num1 = strings.TrimSpace(num1)

		fmt.Print("Üçgeninin yüksekliğini giriniz:  ")
		num2, _ := ucgen.ReadString('\n')
		num2 = strings.TrimSpace(num2)

		sayi1, _ := strconv.ParseFloat(num1, 64)
		sayi2, _ := strconv.ParseFloat(num2, 64)

		r1 := ucgenalanhesap{sayi1, sayi2}
		fmt.Println("Alan: ", r1.ucgenalan())

}

	func cıkgir() {
		cıkgir := bufio.NewReader(os.Stdin)

		fmt.Print("Başka hesaplama yapmak için 1 çıkış yapmak için 2:  ")
		asama, _ := cıkgir.ReadString('\n')
		asama = strings.TrimSpace(asama)

		if asama == "1" {
			main()
		} else if asama == "2" {
			fmt.Print("Program Kapatılıyor !!!! ")
			fmt.Print("Program Kapatıldı! ")
		} else {
			fmt.Print("Böyle bir komut bulunmamakta program kapatıldı")
		}

}

	func main() {
		fmt.Println("Alan ve çevre hesapla programına hoş geldiniz.")
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Dikdörtgen için 1 kare için 2 yazınız ve üçgen için 3 yazınız:  ")
		bilgi, _ := reader.ReadString('\n')
		bilgi = strings.TrimSpace(bilgi)

		if bilgi == "1" {
			dikdortgen()
		} else if bilgi == "2" {
			kare()

		} else if bilgi == "3" {
			ucgenalan1()
			ucgencevre1()

		} else {
			fmt.Println("Böyle bir komut bulunmamakta")
		}
	}
*/

package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var students = []Student{
	{Name: "Fırat", Class: "1-a", Teacher: "Ali"},
	{Name: "Ahmet", Class: "2-a", Teacher: "Mustafa"},
}

type Student struct {
	Name    string `json:"name"`
	Class   string `json:"class"`
	Teacher string `json:"teacher"`
}

func listStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

func createStudents(c *gin.Context) {
	var studentByUser Student
	err := c.Bind(&studentByUser)

	if err == nil && studentByUser.Name != "" && studentByUser.Class != "" && studentByUser.Teacher != "" {
		students = append(students, studentByUser)
		c.JSON(http.StatusCreated, gin.H{"message": "Student has been created", "studentname": studentByUser.Name})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Student cannot be created "})
		return
	}
}

func getStudentByname(name string) (*Student, error) {
	for n, s := range students {
		if s.Name == name {
			return &students[n], nil
		}
	}
	return nil, errors.New("student cannot be found")
}

func getStudent(c *gin.Context) {
	nameget := c.Param("name")

	student, err := getStudentByname(nameget)
	if err == nil {
		c.JSON(http.StatusOK, student)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "student cannot be found"})
	}
}

func main() {
	r := gin.Default()
	r.POST("/students", createStudents)
	r.GET("/students", listStudents)
	r.GET("/students/:name", getStudent)
	err := r.Run("localhost:9090")
	if err != nil {
		panic(err)
	}

}
