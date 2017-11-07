package main

import (
	"log"
	"time"
	"errors"
	"net/http"
)


func HttpPing(url string , protocol string) (int, error) {
	// check if url works

	if url == "" {
		error := errors.New("URL is NOT VALID")
		return -1, error
	}

	log.Printf("Pinging %v ", url)

	// local properties
	var duration int
	timeNow := time.Now().Nanosecond()

	getUrl := protocol + "://" + url

	// get the url first data
	_, err := http.Get(getUrl)
	if err != nil {
		log.Printf("ERROR %v", err)
		return duration, err
	}

	duration = time.Now().Nanosecond() - timeNow
	duration = duration / 1000

	log.Printf("TOTAL DURATION: %vms", duration)

	return duration, err

}

func main() {
	ping, err := HttpPing("obe.sandals.com", "http")

	if err != nil {
		log.Fatal("ERROR:", err)
	}

	log.Printf("TOTAL:%vms", ping)
}

