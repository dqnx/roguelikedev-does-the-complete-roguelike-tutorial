package main

import (
	"fmt"
	blt "github.com/dqnx/bearlibterminal"
	"github.com/faiface/mainthread"
)

func run() {
	blt.Open()
	defer blt.Close();
	
	blt.Print(1, 1, "/r/roguelikedev Tutorial!")
	blt.Refresh()
	
	for blt.Read() != blt.TK_CLOSE {
	}  
}

func main() {
	mainthread.Run(run)
}
