package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var (
		id            int
		business_name string
	)

	db, err := sql.Open("sqlite3", "./idium.db")
	checkErr(err)

	var sql string = "SELECT * FROM customers"
	rows, err := db.Query(sql)

	for rows.Next() {
		err := rows.Scan(&id, &business_name)
		checkErr(err)
		fmt.Println(id, business_name)
	}
	rows.Close()
	db.Close()

}
