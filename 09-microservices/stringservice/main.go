package main

import(
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// global properties
// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// StringService provides operations on strings.
type StringService interface {
	UpperCase(string) (string, error)

}


// defining struct
type stringservice struct {}

// call the interfaces
func (stringservice) UpperCase(s string) (string, error) {
	// handling error
	if s == "" {
		return "", ErrEmpty
	}

	return strings.ToUpper(s), nil
}


func main() {
	svc := stringservice{}

	// handler for the interfaces
	upperCaseHandler := httptransport.NewServer(
		makeUpperCaseEndPoint(svc),
		decodeUpperCaseRequest,
		encodeResponse,
	)

	// creating the http router
	http.Handle("/uppercase", upperCaseHandler)


	// run the server
	log.Fatal(http.ListenAndServe(":8000", nil))


}

// handler the routes
func makeUpperCaseEndPoint(svc stringservice) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.UpperCase(req.S)

		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}

		return uppercaseResponse{v, ""}, nil
	}
}

func decodeUpperCaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// encoding the response for everyting
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}


type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err, omitempty"`
}






