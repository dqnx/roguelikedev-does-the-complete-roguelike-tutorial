package main

import (
	"fmt"
	"gitlab.com/rauko1753/gorl"
)

func main() {
	fmt.Println("/r/roguelikedev Tutorial!")
	
	// Setup terminal view
	gorl.TermMustInit()
	defer gorl.TermDone()
	
}
