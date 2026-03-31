// Illustrates the use of directional channels and goroutines.
// User input is sent to a goroutine, which processes and returns the result via another channel.
package main

import (
	"fmt"
	"strings"
)

// ping <-chan          someone should send into ping
// pong chan<-           someone should receive from pong
func shout(ping <-chan string, pong chan<- string) {

	data := <-ping                // this is blocking, Its waiting for someone to send
	pong <- strings.ToUpper(data) // this is blocking, its waiting for someone to receive, this will affect response := <-pong

}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	var userInput string
	_, _ = fmt.Scanln(&userInput)

	ping <- userInput  // this is that someone who is sending it will affect data := <-ping
	response := <-pong // this is that someone who is receiving

	fmt.Println("Response ", response)

	// no need to close as there no loops as such
	// https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link#33430115a66280e79243f7611299a15e

}
