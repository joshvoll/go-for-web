package main

import (
	"fmt"

)

// creating a struct for channls
type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}


func main() {
	// defining 2 channels base on the struct
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	// creating a basic struct with data for the message struct
	msg := Message{
		To: []string{"jrodriguez@sanservices.hn"},
		From: "mkunz@sanservices.hn",
		Content: "Keep it secreat, keep it safe",
	}

	// creating the basic object fo rthe error message struct
	// errMsg := FailedMessage{
	// 	ErrorMessage: "Message intercepter by black rider",
	// 	OriginalMessage: Message{},
	// }

	msgCh <- msg

	select {
		case receiveMsg := <- msgCh:
			fmt.Println(receiveMsg)
		case receiveError := <- errCh:
			fmt.Println(receiveError)
	}


}