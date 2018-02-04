package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"crypto/sha256"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "src/me/simple_database_login/helpers.go"
)
const BlockSize = 64

type User struct {
	id int
	email    string
	password string
}




func do_index(RESPONSE http.ResponseWriter, REQUEST *http.Request) {

	REQUEST.ParseForm()

	fmt.Println(REQUEST.Form["url_long"])

	for k, v := range REQUEST.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))

	}
}

func do_login(RESPONSE http.ResponseWriter, REQUEST *http.Request) {
	fmt.Println("method:", REQUEST.Method) //get request method
	if REQUEST.Method == "GET" {
		TEMPLATE, _ := template.ParseFiles("views/do_login.html")
		TEMPLATE.Execute(RESPONSE, nil)
	} else {
		REQUEST.ParseForm()

		//Let's make us some strings that we can use and pass out.
		

		// logic part of log in
		fmt.Println("email:", REQUEST.Form["email"][0])
		fmt.Println("password:", REQUEST.Form["password"][0])

		email := REQUEST.Form["email"][0]
		hash := hash_password(REQUEST.Form["password"][0])
		fmt.Println(hash)

		TEMPLATE, _ := template.ParseFiles("views/secure.html")
		TEMPLATE.Execute(RESPONSE, REQUEST.Form)

	}
}

func main() {

	//Let's do our routing.

	http.HandleFunc("/", do_index) // setting router rule
	http.HandleFunc("/login", do_login)
	err := http.ListenAndServe(":4000", nil) // setting listening port
	if err != nil {
		//TEMPLATE, _ := template.ParseFiles("views/error.html")
		//TEMPLATE.Execute(err, nil)
		log.Fatal("ListenAndServe: ", err)
	}
}
