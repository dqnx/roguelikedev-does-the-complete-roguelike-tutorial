package main

import (
	"fmt"
	blt "github.com/dqnx/bearlibterminal"
	v2 "github.com/dqnx/roguelikedev-does-the-complete-roguelike-tutorial/vector2"
	"github.com/faiface/mainthread"
	"time"
)

func run() {
	// Setup terminal
	size := v2.Vector{x: 80, y: 25}
	config := "window: size=" + size.x + "x" + size.y + ", cellsize=auto, title='roguelike'; font: default;"
	blt.Set(config)
	blt.Color("white")
	blt.Bkcolor("black")
	
	blt.Open()
	defer blt.Close()
	
	// Initial Screen
	blt.Print(1, 1, "/r/roguelikedev Tutorial!")
	
	const (
		frametime := time.Nanosecond * (1.0e9 / 60.0)
	)
	
	GameLoop:
		for {
			// Start loop execution timer
			start := time.Now()

			// Render buffer
			blt.Refresh()

			// Handle input

			for blt.HasInput() {
				key := blt.Read()
				switch key {
					case blt.TK_CLOSE: break GameLoop
				}
			}

			// Update

			// Wait for loop cycle to equal frametime duration
			finish := time.Since(start)
			remainder := frametime.Sub(finish)

			if (remainder > 0) {
				time.Sleep(remainder)	
			} 
			else {
				fmt.Println("W: Update loop is lagging behind by", remainder)
			}
		}
}

func main() {
	// Enables use of graphics calls on main os thread and goroutines together
	mainthread.Run(run)
}
