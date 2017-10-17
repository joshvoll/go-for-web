package main

import(
	"html/template"
	"net/http"
	"net/url"
	"log"
)

// global properties
type hotdog int
var tpl *template.Template


// interface for the http servers
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handler the request and response
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	// data to collect and pass to the templates
	data := struct {
		Method string
		Submissions url.Values
	}{
		r.Method,
		r.Form,
	}

	// execution the template and passing the data
	tpl.ExecuteTemplate(w, "methods.html", data)
}

// initialize the template
func init() {
	// parsing the html files
	tpl = template.Must(template.ParseFiles("methods.html"))
}

func main() {
	// defining the hotdog variables
	var d hotdog
	http.ListenAndServe(":8080", d)
}

