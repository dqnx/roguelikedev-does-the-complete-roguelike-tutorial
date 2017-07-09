package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	blt "github.com/dqnx/bearlibterminal"
	v2 "github.com/dqnx/vector2"
	"github.com/faiface/mainthread"
)

func run() {
	str := strconv.Itoa
	argb := blt.ColorFromARGB

	// Open terminal.
	blt.Open()
	defer blt.Close()

	// Setup terminal.
	size := v2.Vector{X: 80, Y: 40}
	config := "window: size=" + str(size.X) + "x" + str(size.Y) + ", title='roguelike'; font: ./fonts/FSEX300.ttf, size=20x20"
	blt.Set(config)
	blt.Composition(blt.TK_OFF)

	blt.Print(1, 1, "/r/roguelikedev Tutorial!")
	blt.Refresh()

	// Create draw call buffer channel. Used to create a queue of cells to be drawn.
	drawQueue := make(chan Cell, 1000)

	// Flag to notify the render routine that the rest of the loop is finished.
	loopFinished := false
	bufferDrawn := false

	// Time between draws, in nanoseconds.
	const framesPerSecond = 60
	const frametime = time.Nanosecond * time.Duration(1000000000/framesPerSecond)
	fmt.Println("Frame time target:", frametime)

	// Seed random number generator.
	rgen := rand.New(rand.NewSource(1001))

	// Initialize game map.
	roomMaxSize := 10
	roomMinSize := 6
	maxRooms := 30
	var rooms []rect

	mapSize := v2.Vector{X: 80, Y: 40}
	worldMap := createMap(mapSize)
	var playerStartPos v2.Vector
	var npcStartPos v2.Vector

	for numRooms, i := 0, 0; i < maxRooms; i++ {
		w := roomMinSize + rgen.Intn(roomMaxSize-roomMinSize)
		h := roomMinSize + rgen.Intn(roomMaxSize-roomMinSize)

		x := rgen.Intn(mapSize.X - w - 1)
		y := rgen.Intn(mapSize.Y - h - 1)

		newRoom := createRect(x, y, w, h)
		overlap := false

	OverlapCheck:
		for _, otherRoom := range rooms {
			if newRoom.intersect(otherRoom) {
				overlap = true
				break OverlapCheck
			}
		}

		if !overlap {
			worldMap.createRoom(newRoom)
			roomCenter := newRoom.center()

			if numRooms == 0 {
				// If its the first room, place the player.
				playerStartPos = roomCenter
			} else {
				if numRooms == 1 {
					// If second room, place npc.
					npcStartPos = roomCenter
				}

				// After the first room, tunnels must be carved.
				// Get the previous room's center.
				prevRoom := rooms[numRooms-1].center()

				// Randomly move vertically or horizontally to tunnel.
				if rgen.Intn(1) == 1 {
					// Move horizontally then vertically.
					worldMap.tunnelHori(prevRoom.X, roomCenter.X, prevRoom.Y)
					worldMap.tunnelVerti(prevRoom.Y, roomCenter.Y, roomCenter.X)
				} else {
					// Move vertically then horizontally.
					worldMap.tunnelVerti(prevRoom.Y, roomCenter.Y, prevRoom.X)
					worldMap.tunnelHori(prevRoom.X, roomCenter.X, roomCenter.Y)
				}
			}

			rooms = append(rooms, newRoom)
			numRooms++
		}

	}

	// Initialize entities.
	actors := make([]Actor, 1)

	// Initialize player.
	var player Actor
	player.Name = "Player"
	player.Code = 0x40
	player.Color = argb(255, 255, 255, 255)
	player.Position = playerStartPos

	actors[0] = player

	// Initialize and NPC
	npc := Actor{Name: "NPC"}
	npc.Code = 0x40
	npc.Color = argb(255, 150, 20, 70)
	npc.Position = npcStartPos

	actors = append(actors, npc)

	// Mutex for checking loop state.
	var mutex = &sync.Mutex{}

	// Logic and Render lop synchronization.
	renderDone := make(chan bool)

	// Start parallel rendering.
	go func() {
		for {
			select {
			// More drawing to do...
			case drawInfo := <-drawQueue:
				blt.Color(drawInfo.Color)
				blt.Put(drawInfo.Position.X, drawInfo.Position.Y, drawInfo.Code)

				mutex.Lock()
				bufferDrawn = true
				mutex.Unlock()
			// Drawing done.
			default:
				mutex.Lock()
				if loopFinished {
					if bufferDrawn {
						blt.Refresh()
						// Reset buffer flag.
						bufferDrawn = false
						fmt.Println("Screen drawn.")
					}
					// Reset loop flag.
					loopFinished = false
					renderDone <- true
				} // else, keep waiting
				mutex.Unlock()
			}
		}
	}()

	// Initial sceen draw.
	worldMap.draw(drawQueue)
	for _, a := range actors {
		drawQueue <- a.draw()
	}

GameLoop:
	for {
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
					drawQueue <- actors[0].draw()
				}
			case blt.TK_RIGHT:
				d := v2.Vector{X: 1, Y: 0}
				if !worldMap.collision(&actors[0], d) {
					actors[0].move(d)
					drawQueue <- actors[0].draw()
				}
			case blt.TK_UP:
				d := v2.Vector{X: 0, Y: -1}
				if !worldMap.collision(&actors[0], d) {
					actors[0].move(d)
					drawQueue <- actors[0].draw()
				}
			case blt.TK_DOWN:
				d := v2.Vector{X: 0, Y: 1}
				if !worldMap.collision(&actors[0], d) {
					actors[0].move(d)
					drawQueue <- actors[0].draw()
				}
			}
		}

		// Update game logic.
		// Nothing to do right now..

		mutex.Lock()
		loopFinished = true
		mutex.Unlock()

		// Wait for rendering to finish.
		<-renderDone

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
