package main


import(
	"net/http"
	"io"
)

type dog int
type cat int

func (d dog) ServeHTTP(r http.ResponseWriter, w *http.Request) {
	io.WriteString(r, "dog, dog, dog")
}

func (c cat) ServeHTTP(r http.ResponseWriter, w *http.Request) {
	io.WriteString(r, "cat, cat, cat")
}

func main() {

	var d dog
	var c cat

	mux := http.NewServeMux()

	mux.Handle("/dog", d)
	mux.Handle("/cat", c)

	// run the servers
	http.ListenAndServe(":8080", mux)
}