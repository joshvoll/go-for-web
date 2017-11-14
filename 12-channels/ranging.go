package main

import (
	"strings"
	"fmt"
)

func main() {
	
	phrase := "There are the time that make can make souls \n"
	words := strings.Split(phrase, " ")

	// creating the channel
	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	// creating an infitine loop
	for msg := range ch {
		fmt.Print(msg  + " ")
	}
}