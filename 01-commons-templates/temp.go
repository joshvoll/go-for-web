package main

import (
	"text/template"
	"log"
	"os"
)

// global properties
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.html"))
}

func main() {
	// sending the data
	err := tpl.ExecuteTemplate(os.Stdout, "main.html", ` Focus on your jobs `)

	if err != nil {
		log.Fatal(err)
	}


}