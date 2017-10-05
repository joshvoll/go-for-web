package main

import (
	"log"
	"os"
	"text/template"
)

// globla properties
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("compose.html"))
}

func main() {
	// creating a matp of data
	saga := map[string]string {
		"India":"Ghandi",
		"Honduras":"Lempira",
		"USA":"Reagan",
		"Italia":"Mussulini",
		"Germany":"Hitler",
	}

	// executing the template
	if err := tpl.Execute(os.Stdout, saga); err != nil {
		log.Fatal(err)
	}
}