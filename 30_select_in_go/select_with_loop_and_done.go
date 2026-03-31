// select_with_loop_and_done.go
// Demonstrates select in a for loop with a done channel for graceful goroutine completion in Go.
// Shows how to handle multiple select cases and signal completion.
package main

import (
	"fmt"
	"strings"
)

func m1(mychan chan<- string, mychun <-chan string, done chan<- bool) {
	mychan <- "Hey There"

	data := <-mychun
	fmt.Printf("data: received :  %v\n", data)

	done <- true // Signal completion
}

func normalize(msg string) string {
	return strings.ToUpper(msg)
}

func main() {
	mychan1 := make(chan string)
	mychan2 := make(chan string)
	done := make(chan bool)

	go m1(mychan1, mychan2, done)

	for {
		select {
		case message := <-mychan1:
			message = normalize(message)
			fmt.Printf("message: %v\n", message)

		case mychan2 <- "WWE":
			fmt.Println("sent to mychan2")

		case <-done:
			fmt.Println("Goroutine completed")
			return // Exit cleanly
		}
	}
}
