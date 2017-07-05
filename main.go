package main

import (
	"fmt"
	"strconv"
	"time"

	blt "github.com/dqnx/bearlibterminal"
	v2 "github.com/dqnx/vector2"
	"github.com/faiface/mainthread"
)

func run() {
	str := strconv.Itoa

	// Setup terminal.
	size := v2.Vector{80, 25}
	config := "window: size=" + str(size.X) + "x" + str(size.Y) + ", cellsize=auto, title='roguelike'; font: default;"
	blt.Set(config)
	blt.Color(blt.ColorFromARGB(100, 24, 17, 22))
	//blt.BkColor(blt.ColorFromARGB("black"))
	blt.Composition(blt.TK_OFF)
	// Open terminal.
	blt.Open()
	defer blt.Close()
	blt.Print(1, 1, "/r/roguelikedev Tutorial!")
	blt.Refresh()

	// Time between draws, in nanoseconds.
	const framesPerSecond = 60
	const frametime = time.Nanosecond * time.Duration(1000000000/framesPerSecond)
	fmt.Println("Frame time target:", frametime)

	// Initialize game map and player.
	actors := make([]actor, 1)

	var player *actor
	// Ref to player as index 0.
	player = &actors[0]

	player.Name = "Player"
	player.Code = 0x40
	player.Position.X = 10
	player.Position.Y = 10

GameLoop:
	for {
		// Start loop execution timer.
		var start time.Time
		var finish time.Duration
		start = time.Now()

		// Handle input.
		exit := false
		if blt.HasInput() {
			key := blt.Read()
			switch key {
			case blt.TK_CLOSE:
				exit = true
			case blt.TK_ENTER:
				fmt.Println("entered")
			case blt.TK_LEFT:
				player.move(v2.Vector{-1, 0})
			case blt.TK_RIGHT:
				player.move(v2.Vector{1, 0})
			case blt.TK_UP:
				player.move(v2.Vector{0, -1})
			case blt.TK_DOWN:
				player.move(v2.Vector{0, 1})
			}
		}

		// Update game logic.
		// Nothing to do right now..

		// Draw calls.
		blt.Clear()
		for _, a := range actors {
			color, code := a.draw()
			blt.Color(blt.ColorFromName(color))
			blt.Put(a.Position.X, a.Position.Y, code)
    	}
		blt.Print(1, 1, "I have been drawn")

		// Render the buffer.
		renderStart := time.Now()
		blt.Refresh()
		renderFinish := time.Since(renderStart)

		finish = time.Since(start)
		renderPercent := (renderFinish.Seconds() / finish.Seconds()) * 100.0
		fmt.Println("Frame time:", finish, ", Render time:", renderFinish, ", Rendering:", renderPercent, "%")

		// Exit the game loop if the called by user.
		if exit {
			break GameLoop
		}
	}
}

func main() {
	// Enables use of graphics calls on main os thread and goroutines together.
	mainthread.Run(run)
}
