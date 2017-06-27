package main

import (
	"fmt"
	blt "github.com/dqnx/bearlibterminal"
	"github.com/faiface/mainthread"
)

func run() {
	// Setup terminal
	size
	config := "window: size=80x25, cellsize=auto, title='Omni: menu'; font: default;"
	blt.Set(config)
	blt.Color("white")
	blt.Bkcolor("black")
	
	blt.Open()
	defer blt.Close();
	
	blt.Print(1, 1, "/r/roguelikedev Tutorial!")
	blt.Refresh()
	
	for blt.Read() != blt.TK_CLOSE {
	}  
}

func main() {
	// Enables use of graphics calls on main os thread and goroutines together
	mainthread.Run(run)
}
