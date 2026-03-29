package greet

// as this is inside internal folder, this can't be used in some other project

import (
	"fmt"
	"strings"
)

func Greet(name string){ // important to Capatalize, so it can be exported
	fmt.Println("Hello"+" "+normalizeIt(name))
}

func normalizeIt(str string) string { // cannot export this function brother
	return strings.ToUpper(str)
}