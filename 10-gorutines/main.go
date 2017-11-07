package main

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"encoding/xml"
	"io/ioutil"
	"runtime"
)

type QuoteResp struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChnageYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}

func main() {

	// we're going to give the application 4 cores
	runtime.GOMAXPROCS(4)
	// we're going to create with gorotune
	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmuz",
		"s",
	}

	// we're going to call a variable to know when the applicaiton complete
	numComplete := 0

	// looping to the http request
	for _, symbol := range stockSymbols {
		// stating the gorutine
		go func(symbol string) {
			// make the http request
			res, err := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)

			if err != nil {
				log.Fatal("URL ERROR:", err)
			}
			// defer evetyhing
			defer res.Body.Close()

			body, _ := ioutil.ReadAll(res.Body)
			quote := new(QuoteResp)
			// parse the xml
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: %.2f \n", quote.Name, quote.LastPrice)

			// incremenet the numComplete 
			numComplete++

		}(symbol)
	}

	// check if numCompelte is equal to the size of the stockSymbols slice of string
	for numComplete < len(stockSymbols) {
		// we wait for 10 miliseconds
		time.Sleep(10 * time.Millisecond)
	}

	elpase := time.Since(start)

	fmt.Printf("Execution time: %s \n", elpase)
}



