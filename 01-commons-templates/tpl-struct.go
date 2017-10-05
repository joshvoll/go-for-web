package main

import (
	"log"
	"os"
	"text/template"
)

// global propertes
var tpl *template.Template

// creating the struct
type saga struct{
	Name string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("struct.html"))
}

func main() {
	// creating teh struct
	mohamed := saga{
		Name: "Mohamed",
		Motto: "The belief or not belief",
	}

	morazan := saga{
		Name: "Francisco Morazan",
		Motto: "tegucigalpa es para siempre",
	}

	trump := saga{
		Name: "Donald J. Trump",
		Motto: "Sorry to complicate bussiness",
	}

	// adding all elemnts to a variable
	sagas := []saga{mohamed, morazan, trump}
	// executing the templates
	if err := tpl.Execute(os.Stdout, sagas); err != nil {
		log.Fatal(err)
	}
}

