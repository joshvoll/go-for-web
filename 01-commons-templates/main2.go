package main

import(
	"fmt"
	"os"
	"text/template"
)
	
func main() {
	tpl, err := template.ParseFiles("letter.html")

	if err != nil {
		fmt.Println("There was an error parsing the file")
	}

	friends := []string{"Alex", "Conor", "Ken", "Cindy", "Alanis"}

	err = tpl.Execute(os.Stdout, friends)

	if err != nil {
		fmt.Println("erro parsin ghe file")
	}
}