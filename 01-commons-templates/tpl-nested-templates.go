package main

import(
	"log"
	"os"
	"text/template"
)

// global propierties
var tpl  *template.Template

// defingin struct files
type person struct{
	FirstName string
	LastName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	// generating the files
	software := person{
		FirstName: "Luis",
		LastName: "Matute",
	}

	qa := person{
		FirstName: "Fernando",
		LastName: "Lontero",
	}

	pm := person{
		FirstName: "Arlene",
		LastName: "Espana",
	}

	companyPositions := []person{software, qa, pm}

	// executions the templates
	if err := tpl.ExecuteTemplate(os.Stdout, "index.html", companyPositions	); err != nil {
		log.Fatal(err)
	}
}