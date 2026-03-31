// Demonstrates using sync.WaitGroup to coordinate multiple concurrent goroutines.
// Shows how to fetch data in parallel and wait for all tasks to complete before exiting main.
package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func dataFetcher(key string, wg *sync.WaitGroup) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/" + key)
	if err != nil {
		panic(err)
	}
	// 1. Always close the body when finished
	defer res.Body.Close()
	defer wg.Done()
	// 2. Read the body stream into bytes
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// 3. Convert bytes to string to print it
	fmt.Println(string(body))
}
func main() {
	// go run 26_go_wait_groups/main.go  > test.txt
	keywords := []string{"comments", "users"}

	// for i,val:=range keywords{
	// 	fmt.Println(i,val)
	// }

	var wg sync.WaitGroup

	for _, key := range keywords {
		wg.Add(1)
		go dataFetcher(key, &wg)
	}

	wg.Wait()

}
