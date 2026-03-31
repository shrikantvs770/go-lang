package greet

// as this is inside internal folder, this can't be used in some other project

import (
	"fmt"
	"strings"
)

var Quote string = "To be or not to be that's the Q." // this is ued in app.go
// var qt = "Hey there" This can't be exported and used in app.go

func Greet(name string){ // important to Capatalize, so it can be exported
	fmt.Println("Hello"+" "+normalizeIt(name))
}

func normalizeIt(str string) string { // cannot export this function brother
	return strings.ToUpper(str)
}