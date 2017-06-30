package main

import (
	"errors"
	"fmt"
	"time"
	
	blt "github.com/dqnx/bearlibterminal"
	v2 "github.com/dqnx/vector2"
	"github.com/faiface/mainthread"
)

func run() {
	// Setup terminal.
	size := v2.Vector{x: 80, y: 25}
	config := "window: size=" + size.x + "x" + size.y + ", cellsize=auto, title='roguelike'; font: default;"
	blt.Set(config)
	blt.Color("white")
	blt.Bkcolor("black")
	
	// Open terminal.
	blt.Open()
	defer blt.Close()
	blt.Print(1, 1, "/r/roguelikedev Tutorial!")
	
	const (
		// Time between draws, in nanoseconds.
		frametime := time.Nanosecond * (1.0e9 / 60.0)
	)
	
	GameLoop:
		for {
			// Start loop execution timer.
			start := time.Now()

			// Render the terminal buffer.
			blt.Refresh()

			// Handle input.
			for blt.HasInput() {
				key := blt.Read()
				switch key {
					case blt.TK_CLOSE: break GameLoop
				}
			}

			// Update game logic.

			
			finish := time.Since(start)
			remainder := frametime.Sub(finish)
			
			/*	
			If remainder is greater than zero, the loop is running behind the
			desired frame time. This will cause frames per second to not meet the
			target.
			*/
			if (remainder < 0) {
				fmt.Println("W: Update loop is lagging behind by", remainder)
				return
			}
			
			// Wait for loop cycle to equal frametime duration.
			time.Sleep(remainder)	
		}
}

func main() {
	// Enables use of graphics calls on main os thread and goroutines together.
	mainthread.Run(run)
}
