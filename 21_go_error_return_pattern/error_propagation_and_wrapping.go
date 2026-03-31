// Demonstrates error propagation and wrapping through multiple function layers.
// Shows how to add context to errors as they bubble up the call stack.
package main

import (
	"errors"
	"fmt"
)

func main() {
	res := handler(0)
	if res != nil {
		fmt.Println(res)
	}
}
func getUserFromDB(id int) (string, error) {
	if id == 0 {
		return "", errors.New("invalid user id")
	}
	return "Srikant", nil
}

func getUserName(id int) (string, error) {
	name, err := getUserFromDB(id)
	if err != nil {
		return "", fmt.Errorf("service: fetching user %d: %w", id, err)
	}
	return name, nil
}

func handler(id int) error {
	name, err := getUserName(id)
	if err != nil {
		return fmt.Errorf("handler: failed to get user: %w", err)
	}
	fmt.Println(name)
	return nil
}
