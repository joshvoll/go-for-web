package main 


import (
 	"time"
	"net/http"
	"fmt"
	"io"
	"encoding/json"
	"os"
	"bytes"

)

type sizeWriter int


// defining the interface
type Response interface {
	Status() int
	Header() http.Header
	HeaderSize() int
	BodySize() int
	TimeResponse(time.Time) time.Duration
	Stats() *Stats
}

// stats strct this will be in json format
type Stats struct {
	Status       int           `json:"status, omitempty"`
	Header       http.Header   `json:"header, omitempty"`
	HeaderSize   int           `json:"header_size, omitempty"`
	BodySize     int           `json:"body_size, omitempty"` 
	TimeResponse time.Duration `json:"time_response"`

}

// struct of the projects
type response struct {
	status     int
	header     http.Header
	headerSize int
	bodySize   sizeWriter
}

// global properties
var DefaultClient = &http.Client{}

// declaring interfaces
func (w *sizeWriter) Write(b []byte) (int, error) {
	*w += sizeWriter(len(b))
	return len(b), nil
}

func (r *response) BodySize() int {
	return int(r.bodySize)
}

func (r *response) Header() http.Header {
	return r.header
}

func (r *response) HeaderSize() int {
	return r.headerSize
}

func (r *response) Status() int {
	return r.status
}

// implementing time resposne
func (r *response) TimeResponse(now time.Time) time.Duration {
	// returntin the time to connecitons
	return now.Sub(time.Now())
}

func (r *response) Stats() *Stats {

	// local properties
	now := time.Now()
	// return the objects
	return &Stats{
		Status: r.Status(),
		Header: r.Header(),
		HeaderSize: r.HeaderSize(),
		BodySize: r.BodySize(),
		TimeResponse: r.TimeResponse(now),
	}
}



func RequestWithClient(client *http.Client, method string, uri string, header http.Header, body io.Reader) (Response, error) {
	// get the url 
	req, err := http.NewRequest(method, uri, body)

	if err != nil {
		return nil, err
	}

	// looop throw the header
	for name, field := range header {
		for _, v := range field {
			req.Header.Set(name, v)
		}
	}

	// local propiertis
	var out response
	req = req.WithContext(req.Context())

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// defer everything 
	defer res.Body.Close()

	// filling all variables
	out.status = res.StatusCode
	// get teh body size
	if _, err := io.Copy(&out.bodySize, res.Body); err != nil {
		return nil, err
	}
	var resHeader bytes.Buffer
	res.Header.Write(&resHeader)
	out.header = res.Header
	out.headerSize = resHeader.Len()

	return &out, nil


}

func Request(method string, uri string, header http.Header, body io.Reader) (Response, error) {
	return RequestWithClient(DefaultClient, method, uri, header, body)
}

func main() {

	res, err := Request("GET", "https://www.sandalsselect.com", nil, nil)

	if err != nil {
		fmt.Println("ERROR URL:", err)
	}

	enc := json.NewEncoder(os.Stderr)
	enc.SetIndent("", "  ")
	enc.Encode(res.Stats())
}

