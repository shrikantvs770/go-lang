package main

import (
	"log"
	"os"
)

func main() {
	// https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link#33230115a6628050a19dec58369db2c8
	defer log.Println("A")
	err := os.WriteFile("dummy.txt", []byte("Hi there brother!!"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("dummy.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	defer log.Printf("B")
	defer log.Println("C")

}
