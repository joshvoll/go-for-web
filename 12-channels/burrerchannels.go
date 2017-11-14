package main

import (
	"fmt"
	"strings"
)

func main() {

	phrase := "These are the time that man can make souls"
	words  := strings.Split(phrase, " ")

	// creating the buffered channel
	ch := make(chan string, len(words))

	// loop throw the words and set the letter into the channel
	for _, word := range words {
		ch <- word
	}

	// loop throw the channel to get all the letters
	for i := 0; i < len(words); i++ {
		fmt.Print(<- ch + "")
	}
	


}