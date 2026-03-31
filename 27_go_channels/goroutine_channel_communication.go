// Demonstrates safe channel communication using goroutines.
// A goroutine receives from one channel and sends to another, avoiding deadlock.
package main

import "fmt"

func main() {
	ping := make(chan string)
	pong := make(chan string)

	// both sender and receiver are executing at the same time, so no deadlock, eventhough there is no buffer allocated.

	go func() {
		msg := <-ping
		pong <- msg + " world"
	}()

	ping <- "Hello"
	greetings := <-pong

	fmt.Println(greetings)
}
