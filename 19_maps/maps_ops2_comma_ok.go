package main

import "fmt"

func main() {
	var stdMap map[string]string = map[string]string{
		"u001": "Srikant V S",
		"u002": "Elias Samson",
		"u003": "WWE",
		"u004": "John Cena",
		"u005": "Shawn Michaels",
	}

	dummyWrestler := stdMap["u888"]
	fmt.Println(dummyWrestler) // Prints empty, not good

	dummyWrestler, ok := stdMap["u888"]
	if ok {
		fmt.Println("He exists", dummyWrestler)
	} else {
		fmt.Println("He doesn't exist")
	}


	// We can write the if like this too, see the 11 folder brother
	if dummyWrestler, ok:=stdMap["u001"]; ok {
		fmt.Println(dummyWrestler, " exists")
	} else {
		fmt.Println("Oops!!")
	}

}
