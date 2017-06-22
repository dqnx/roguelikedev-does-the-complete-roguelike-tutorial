package main

import (
	"fmt"
	"github.com/
	"gitlab.com/rauko1753/gorl"
	"github.com/faiface/mainthread"
	blt "bitbucket.org/cfyzium/bearlibterminal"
)

func run() {
	fmt.Println("/r/roguelikedev Tutorial!")
}

func main() {
	// Allows rendering to be fixed to the first os thread. Needed for Windows
	mainthread.Run(run)
}
