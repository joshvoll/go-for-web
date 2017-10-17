package main

import (
	"net/http"
	"html/template"
	"log"
)

type hotdog int
var tpl *template.Template

// loading and interface base on hotdog int
func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// try to parse the form
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	// if everything is good we're going to load the template
	tpl.ExecuteTemplate(w, "form.html", r.Form)
}

func init() {
	// loading the templates
	tpl = template.Must(template.ParseFiles("form.html"))
}

func main() {
	// first les get the variable 
	var d hotdog
	http.ListenAndServe(":8080", d)
}