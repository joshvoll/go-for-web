package main

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"encoding/xml"
	"io/ioutil"

)

type QuoteRes struct {
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
	// this will run on a go rutine
	startTime := time.Now()

	// get all the stock id from companies
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

	// make the request to the url, we need to loop throw the 
	for _, symbol := range stockSymbols {
		res, err := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)

		if err != nil {
			log.Fatal("URL ERROR: ", err)
		}
		// defer evertyhing
		defer res.Body.Close()
		// read with io util
		body, _ := ioutil.ReadAll(res.Body)

		quote := new(QuoteRes)

		xml.Unmarshal(body, &quote)

		fmt.Printf("%s: %.2f \n", quote.Name, quote.LastPrice)
		
	}

	elapse := time.Since(startTime)
	fmt.Printf("Execution time: %s \n", elapse)
	
}


