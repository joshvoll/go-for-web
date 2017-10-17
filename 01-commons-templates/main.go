package main
 
import(
	"fmt"
	"text/template"
	"net/http"
	"log"
)

// global propierties
var tpl *template.Template

// createin a struct
type person struct{
    FirstName string
    LastName string
    Age int
}

// pre-defined methods
func (software person) DoubleAge() int {
	fmt.Println("AGE:", software.Age)
	return software.Age * 2
}

func (software person) YearBirth(x int) int {
	return software.Age - 2017
}

func init() {
	// loadging all the templates
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
    // loading the http servers
	http.HandleFunc("/", Index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8000", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {
	// setting data
	software := person{
		FirstName: "Josue",
		LastName: "Rodriguez",
		Age: 38,
	}



	// loading the templates
	if err := tpl.ExecuteTemplate(w, "index.html", software); err != nil {
		log.Fatal(err)
	}
}

