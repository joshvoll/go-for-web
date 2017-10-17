package main

import(
	"html/template"
	"log"
	"net/http"
)
// global properties
var tpl *template.Template

// creating a struct from the form 
type Person struct {
	FirstName string
	LastName string
	Subscribed bool
}

// initialize the object
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))

}

func main() {
	// getting server start it 
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	// get the variablos from form
	f := r.FormValue("first")
	l := r.FormValue("last")
	s := r.FormValue("subscribe") == "on"

	// render the information to the template
	if err := tpl.ExecuteTemplate(w, "index.html", Person{f,l,s}); err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
	}


}