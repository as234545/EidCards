// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (

	// "fmt" has methods for formatted I/O operations (like printing to the console)

	"log"

	// The "net/http" library has methods to implement HTTP clients and servers
	"html/template"
	"net/http"
)

type Student struct {
	Img       string
	Size      string
	Colors    string
	TextHight int
	TextWidth int
	Fonts     string
}

type User struct {
	Name string
}

func main() {

	// The "HandleFunc" method accepts a path and a function as arguments
	// (Yes, we can pass functions as arguments, and even treat them like variables in Go)
	http.HandleFunc("/", index)
	// However, the handler function has to have the appropriate signature (as described by the "handler" function below)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/card1", card1)
	http.HandleFunc("/card2", card2)
	http.HandleFunc("/card3", card3)
	http.HandleFunc("/card4", card4)
	http.HandleFunc("/card5", card5)
	http.HandleFunc("/card6", card6)
	http.HandleFunc("/card7", card7)
	http.HandleFunc("/card8", card8)
	http.HandleFunc("/card9", card9)

	// After defining our server, we finally "listen and serve" on port 8080
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	t, err := template.ParseFiles("templates/index5.html") //index4
	t.Execute(w, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// public serves static assets such as CSS and JavaScript to clients.
func public() http.Handler {
	return http.StripPrefix("/css/", http.FileServer(http.Dir("./css")))
}

func card1(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Img:       "/static/img/eidcards9_noname.png",
		Size:      "40px",
		Colors:    "rgb(99, 127, 156)",
		TextHight: 115,
		TextWidth: 5,
		Fonts:     "/static/fonts/FontsFree-Net-din-next-arabic-bold.ttf",
	}
	parsedTemplate, _ := template.ParseFiles("templates/testpreview2.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func card2(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Img:       "/static/img/eidcards1_noname-01.png",
		Size:      "70px",
		Colors:    "rgb(255, 255, 255)",
		TextHight: 400,
		TextWidth: 5,
	}
	parsedTemplate, _ := template.ParseFiles("templates/testpreview2.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func card3(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Img:       "/static/img/eidcards4_noname.png",
		Size:      "40px",
		Colors:    "rgb(86, 134, 135)",
		TextHight: 150,
		TextWidth: 0,
	}
	parsedTemplate, _ := template.ParseFiles("templates/testpreview2.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func card4(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Img:       "/static/img/eidcards10_noname.png",
		Size:      "40px",
		Colors:    "rgb(131, 158, 188)",
		TextHight: 300,
		TextWidth: 200,
	}
	parsedTemplate, _ := template.ParseFiles("templates/testpreview2.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func card5(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Img:       "/static/img/eidcards6_noname.png",
		Size:      "40px",
		Colors:    "rgb(255, 255, 255)",
		TextHight: 170,
		TextWidth: 5,
	}
	parsedTemplate, _ := template.ParseFiles("templates/testpreview2.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func card6(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Img:       "/static/img/eidcards7_noname-01.png",
		Size:      "70px",
		Colors:    "rgb(123, 106, 98)",
		TextHight: 400,
		TextWidth: 5,
	}
	parsedTemplate, _ := template.ParseFiles("templates/testpreview2.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
func card7(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Img:       "/static/img/eidcards8_noname-01.png",
		Size:      "70px",
		Colors:    "rgb(0, 0, 0)",
		TextHight: 300,
		TextWidth: 5,
	}
	parsedTemplate, _ := template.ParseFiles("templates/testpreview2.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func card8(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Img:       "/static/img/eidcards3_noname.png",
		Size:      "40px",
		Colors:    "rgb(255, 255, 255)",
		TextHight: 150,
		TextWidth: 5,
	}
	parsedTemplate, _ := template.ParseFiles("templates/testpreview2.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func card9(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Img:       "/static/img/eidcards5_noname-01.png",
		Size:      "70px",
		Colors:    "rgb(82, 88, 114)",
		TextHight: 380,
		TextWidth: 5,
	}
	parsedTemplate, _ := template.ParseFiles("templates/testpreview2.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
