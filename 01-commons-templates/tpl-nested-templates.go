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
	Age int
}

// pre-defined methods
func (sofware person) DoubleAge() int {
	return software.Age * 2
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	// generating the files
	software := person{
		FirstName: "Luis",
		LastName: "Matute",
		Age: 32,
	}

	qa := person{
		FirstName: "Fernando",
		LastName: "Lontero",
		Age: 26,
	}

	pm := person{
		FirstName: "Arlene",
		LastName: "Espana",
		Age: 23,
	}

	companyPositions := []person{software, qa, pm}

	// executions the templates
	if err := tpl.ExecuteTemplate(os.Stdout, "index.html", companyPositions	); err != nil {
		log.Fatal(err)
	}
}