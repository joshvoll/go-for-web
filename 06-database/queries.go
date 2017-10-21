package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"io"
)

 
var db *sql.DB
var err error

func main() {
	// creating the cpnnecton from the database
	db, err = sql.Open("mysql", "root:sandals01@tcp(10.102.1.56:3306)/database_test?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	// defer the database
	defer db.Close()

	// ping the database 
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// creating the web server
	http.HandleFunc("/", home)
	http.HandleFunc("/users", users)

	// turn on the servsers
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	// just say everything is fine
	io.WriteString(w, "Conneciton to the database complete")
}

func users(w http.ResponseWriter, r *http.Request) {
	// creating the query
	var s, name, lastname, email string

	// creating the string query
	rows, err := db.Query("SELECT name, lastname, email FROM Users;")

	if err != nil {
		log.Fatal(err)
	}

	// defer the query
	defer rows.Close()


	// for loop the query
	for rows.Next() {
		if err := rows.Scan(&name, &lastname, &email); err != nil {
			log.Fatal(err)
		}

		s += name + lastname + email + "\n"
	}

	fmt.Fprintln(w, s)


}