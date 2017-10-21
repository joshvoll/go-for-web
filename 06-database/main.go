package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"io"
)

func main() {
	// connectin to the database 
	db, err := sql.Open("mysql", "root:sandals01@tcp(10.102.1.56:3306)/database_test?charset=utf8")

	if err != nil {
		log.Fatal("DATABASE ERROR: ", err)
	}

	defer db.Close()

	// ping de database
	if err := db.Ping(); err != nil {
		log.Fatal("PING ERROR:" ,  err)
	}

	// createing route to send the data to the browser
	http.HandleFunc("/", Home)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func Home(w http.ResponseWriter, r *http.Request) {

	_, err := io.WriteString(w, "Succesfully completed connected")

	if err != nil {
		log.Fatal(err)
	}
}