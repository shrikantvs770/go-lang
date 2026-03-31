package main

import "fmt"

func main() {
	chan1 := make(chan bool)


	go func() {
		chan1 <- true
	}()

	select {
	case x := <-chan1:
		fmt.Printf("received x: %v\n", x)
	}
}