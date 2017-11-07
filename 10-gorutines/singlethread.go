package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"log"
	"fmt"
	"time"
)


// creating the struct of the document
type QuoteResponse struct {
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
	// get the url form webiste
	startTime := time.Now()

	res, err := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=googl")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	quote := new(QuoteResponse)

	xml.Unmarshal(body, &quote)

	fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)

	elapse := time.Since(startTime)

	fmt.Printf("Execution time %s: \n", elapse)

}



