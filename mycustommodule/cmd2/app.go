package main

// import "module name"/internal/greet
import "github.com/srikantvs26/my-module/internal/greet"

/*
can have multiple main packages but keep them in different folder
*/
func main() {
	greet.Greet("srikant v s")
}
