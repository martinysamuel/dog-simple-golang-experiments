package main

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

var (
	email   string
	id      int
	hash    string
	message string
)

func do_index(RESPONSE http.ResponseWriter, REQUEST *http.Request) {
	TEMPLATE, _ := template.ParseFiles("./views/login.html")
	TEMPLATE.Execute(RESPONSE, nil)
}

func do_login(RESPONSE http.ResponseWriter, REQUEST *http.Request) {

	//fmt.Println("Login function has been hit.")

	const BlockSize = 64

	//fmt.Println("method:", REQUEST.Method) //get request method
	if REQUEST.Method == "GET" {

		message = "Method issue. We should probably do a redirect on this piece. I'll do that in the next one."
		out := bytes.NewBufferString(message)
		TEMPLATE, _ := template.ParseFiles("./views/error.html")
		TEMPLATE.Execute(RESPONSE, out)

	} else {

		//We have a post state! Excellent.

		DATABASE, err := sql.Open("sqlite3", "./data/idium.db")
		do_error_checking(err)
		REQUEST.ParseForm()

		email = REQUEST.FormValue("email")
		hash = hash_password(REQUEST.FormValue("password"))

		var statement string = fmt.Sprintf("SELECT id,email FROM users where email='%s' and password ='%s'", email, hash)

		data := DATABASE.QueryRow(statement, 3)

		//Here, we can't take the shortcut with the easy error checking,
		//because we can have a variety of errors here. So we switch it.

		switch err := data.Scan(&id, &email); err {
		case sql.ErrNoRows:
			//No such user.
			message = ("Not a valid user, sorry. Please try again.")
			out := bytes.NewBufferString(message)
			TEMPLATE, _ := template.ParseFiles("./views/error.html")
			TEMPLATE.Execute(RESPONSE, out)
		case nil:
			//Valid user
			TEMPLATE, _ := template.ParseFiles("./views/secure.html")
			TEMPLATE.Execute(RESPONSE, REQUEST.Form)
		default:
			//Method or database issue.
			message = "Probably a database issue."
			out := bytes.NewBufferString(message)
			TEMPLATE, _ := template.ParseFiles("./views/error.html")
			TEMPLATE.Execute(RESPONSE, out)
		}

	}
}

func main() {

	//Let's do our routing.

	http.HandleFunc("/", do_index) // setting router rule
	http.HandleFunc("/login", do_login)
	err := http.ListenAndServe(":4000", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//Helper functions, here, for now.

func hash_password(password string) string {
	sum := sha256.Sum256([]byte(password))
	out := fmt.Sprintf("%x", sum)
	return out
}
func do_error_checking(err error) {
	if err != nil {
		panic(err)
	}
}
