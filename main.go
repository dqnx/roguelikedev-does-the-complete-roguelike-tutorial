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
	argb := blt.ColorFromARGB

	// Setup terminal.
	size := v2.Vector{X: 80, Y: 60}
	config := "window: size=" + str(size.X) + "x" + str(size.Y) + ", cellsize=auto, title='roguelike'; font: default;"
	blt.Set(config)
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

	// Initialize game map.
	mapSize := v2.Vector{X: 80, Y: 60}
	worldMap := createMap(mapSize)

	// Add 2 rooms.
	worldMap.createRoom(createRect(20, 15, 10, 15))
	worldMap.createRoom(createRect(50, 15, 10, 15))

	// Connect the rooms.
	worldMap.tunnelHori(25, 55, 23)

	// Initialize entities.
	actors := make([]Actor, 1)

	// Initialize player.
	var player Actor
	player.Name = "Player"
	player.Code = 0x40
	player.Color = argb(255, 255, 255, 255)
	player.Position.X = 25
	player.Position.Y = 23

	actors[0] = player

	// Initialize and NPC
	npc := Actor{Name: "NPC"}
	npc.Code = 0x40
	npc.Color = argb(255, 150, 20, 70)
	npc.Position.X = 21
	npc.Position.Y = 16

	actors = append(actors, npc)

	var renderPercent float64
GameLoop:
	for {
		// Start loop execution timer.
		var start time.Time
		var finish time.Duration
		start = time.Now()

		// Handle input.
		exit := false
		for blt.HasInput() {
			key := blt.Read()

			switch key {
			case blt.TK_CLOSE:
				exit = true
			case blt.TK_ENTER:
				fmt.Println("entered")
			case blt.TK_LEFT:
				d := v2.Vector{X: -1, Y: 0}
				if !worldMap.collision(&actors[0], d) {
					actors[0].move(d)
				}
			case blt.TK_RIGHT:
				d := v2.Vector{X: 1, Y: 0}
				if !worldMap.collision(&actors[0], d) {
					actors[0].move(d)
				}
			case blt.TK_UP:
				d := v2.Vector{X: 0, Y: -1}
				if !worldMap.collision(&actors[0], d) {
					actors[0].move(d)
				}
			case blt.TK_DOWN:
				d := v2.Vector{X: 0, Y: 1}
				if !worldMap.collision(&actors[0], d) {
					actors[0].move(d)
				}
			}
		}

		// Update game logic.
		// Nothing to do right now..

		// Draw calls.
		blt.Clear()

		worldMap.draw()

		for _, a := range actors {
			color, code := a.draw()
			blt.Color(color)
			blt.Put(a.Position.X, a.Position.Y, code)
		}

		// Render the buffer.
		renderStart := time.Now()
		blt.Refresh()
		renderFinish := time.Since(renderStart)

		finish = time.Since(start)
		renderPercent = (renderFinish.Seconds() / finish.Seconds()) * 100.0

		//fmt.Println("Frame time:", finish, ", Render time:", renderFinish, ", Rendering:", renderPercent, "%")

		// Exit the game loop if the called by user.
		if exit {
			break GameLoop
		}
	}
	fmt.Println(renderPercent)
}

func main() {
	// Enables use of graphics calls on main os thread and goroutines together.
	mainthread.Run(run)
}
