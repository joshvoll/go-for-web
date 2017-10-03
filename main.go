/* first web dev servers */
package main 


import (
	
	"fmt"
	"net/http"
	"log"

	"html/template"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"encoding/json"
)

// creating struct for displaying
type Page struct{
	Name string
	DBStatus bool
}


func main() {

	// looading the templates
	templates := template.Must(template.ParseFiles("templates/index.html"))

	// creating the connection to the database 
	db, err := sql.Open("sqlite3", "dev.db")

	if err != nil {
	    log.Fatal(err)
	}

	// creating the handler for the servers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		p := Page{Name: "Gopher"}

		// pasing the object to the pages
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}

		// send the connection to the page
		p.DBStatus = db.Ping() == nil
		// execution the templates object
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	// creating the second route for the comming request handerl
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// tesgint the resulst of the dummy data
		var results []SearchResult
		var err error

		if results, err = Search(r.FormValue("search")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// encoding json the results
		encoder := json.NewEncoder(w)

		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	// route for saving the books
	http.HandleFunc("/books/add", func(w http.ResponseWriter, r *http.Request) {

		// printin the result ID
		fmt.Println("RESULT ID: ", r.FormValue("id"))
		// now save everything to the database
		book, err := Find(r.FormValue("id"))

		if err != nil {
			fmt.Println("ERROR: ", err)
		    http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// ping the db to check if onneciton is app an running
		if err = db.Ping(); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// save everythng to the db
		_, err = db.Exec("insert into books (pk, title, author, id, classification) values (?, ?, ?, ?, ?)", 
			              nil, book.BookData.Title, book.BookData.Author, book.BookData.ID, book.Classification.MostPopular)

		if err != nil {
			fmt.Println("SQL ERROR: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//loading the srevers
	fmt.Println(http.ListenAndServe(":8080", nil))
}







