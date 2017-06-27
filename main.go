package main

import (
	"fmt"
	blt "github.com/dqnx/bearlibterminal"
	v2 "github.com/dqnx/roguelikedev-does-the-complete-roguelike-tutorial/vector2"
	"github.com/faiface/mainthread"
)

func run() {
	// Setup terminal
	size := v2.Vector{x: 80, y: 25}
	config := "window: size=" + size.x + "x" + size.y + ", cellsize=auto, title='Omni: menu'; font: default;"
	blt.Set(config)
	blt.Color("white")
	blt.Bkcolor("black")
	
	blt.Open()
	defer blt.Close();
	
	blt.Print(1, 1, "/r/roguelikedev Tutorial!")
	blt.Refresh()
	
	for running := true; running == true {
		// Handle input
		key := blt.Read()
		switch key {
			case blt.TK_CLOSE: running = false
		}
	}
}

func main() {
	// Enables use of graphics calls on main os thread and goroutines together
	mainthread.Run(run)
}
