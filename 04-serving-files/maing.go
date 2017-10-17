package main  


import (
	"net/http"
	"log"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir("."))))
}