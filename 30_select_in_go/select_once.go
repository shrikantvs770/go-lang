// select_once.go
// Demonstrates a single execution of select statement with two channels in Go.
// Shows how select chooses a ready case and why WWE is never printed.
package main

import (
	"fmt"
	"strings"
)

func m1(mychan chan<- string, mychun <-chan string) {

	mychan <- "Hey There" // its blocking and waiting for a receiver, which we defined in select

	data := <-mychun
	fmt.Printf("data: received :  %v\n", data)

}

func normalize(msg string) string {
	return strings.ToUpper(msg)
}

func main() {
	mychan1 := make(chan string)
	mychan2 := make(chan string)

	go m1(mychan1, mychan2) // go routine working with a channel

	select {
	case message := <-mychan1: // this is receiver
		message = normalize(message)
		fmt.Printf("message: %v\n", message)

	case mychan2 <- "WWE":
		fmt.Print("sent to mychan2")
		// default: having it makes this non blocking, and this executes immediately which I don't want as of now.
		// 	fmt.Println("no data in channel brother!")
	}

}

// We will never see WWE in output as select excecutes only once, we have to wrap in for loop with a return stmt.
