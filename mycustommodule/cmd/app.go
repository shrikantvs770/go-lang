package main

import (
	"fmt"

	"github.com/srikantvs26/my-module/internal/greet"
)

func main() {
	greet.Greet("srikant v s")
	fmt.Println(greet.Quote)
	fmt.Println(greet.qt)
}
