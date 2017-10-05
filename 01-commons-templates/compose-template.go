package main

import (
	"log"
	"text/template"
	"os"
)

// global properties
var tpl *template.Template

func init() {
	// assighe the variable to the template
	tpl = template.Must(template.ParseFiles("compose.html"))
}

func main() {
	// assign values 
	datos := []string{"Josue", "Manuel", "Rodriguez", "Castro"}

	// applying those names to the output
	if err := tpl.Execute(os.Stdout, datos); err != nil {
		log.Fatal(err)
	}

}