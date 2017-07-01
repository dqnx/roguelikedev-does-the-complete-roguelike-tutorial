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
	size := v2.Vector{X: 80, Y: 25}
	config := "window: size=" + str(size.X) + "x" + str(size.Y) + ", cellsize=auto, title='roguelike'; font: default;"
	blt.Set(config)
	blt.Color(blt.ColorFromName("white"))
	blt.BkColor(blt.ColorFromName("black"))

	// Open terminal.
	blt.Open()
	defer blt.Close()
	blt.Print(1, 1, "/r/roguelikedev Tutorial!")
	blt.Refresh()

	// Time between draws, in nanoseconds.
	const framesPerSecond = 60
	const frametime = time.Nanosecond * time.Duration(1000000000/framesPerSecond)
	fmt.Println("Frame time target:", frametime)

	// Flags for synchronization
	exit := make(chan bool, 1)
	inputHandled := make(chan bool, 1)
	updatedLogic := make(chan bool, 1)
	drawn := make(chan bool, 1)

GameLoop:
	for {
		var start time.Time
		var finish time.Duration

		// Start loop execution timer.
		start = time.Now()

		// Handle input.
		go func() {
			//fmt.Println("Input handle entered")
			exitBuffer := false

			if blt.HasInput() {
				//fmt.Println("Has input")
				key := blt.Read()
				switch key {
				case blt.TK_CLOSE:
					exitBuffer = true
				case blt.TK_ENTER:
					fmt.Println("entered")
				}
			}

			//fmt.Println("Input ran")
			exit <- exitBuffer
			//fmt.Println("Input send")
			inputHandled <- true
			//fmt.Println("Input handled")
		}()

		// Update game logic.
		go func() {
			//fmt.Println("Logic Update entered")
			<-inputHandled
			x := 1
			x++
			updatedLogic <- true
			//fmt.Println("Logic Updated")
		}()

		// Draw calls.
		go func() {
			//fmt.Println("Draw entered")
			<-updatedLogic
			blt.Clear()
			blt.Print(1, 1, "I have been drawn")
			drawn <- true
			//fmt.Println("Drawn")
		}()

		//fmt.Println(finish, remainder)
		/*
			If remainder is greater than zero, the loop is running behind the
			desired frame time. This will cause frames per second to not meet the
			target.
		*/
		<-drawn

		renderStart := time.Now()
		blt.Refresh()
		renderFinish := time.Since(renderStart)

		finish = time.Since(start)
		renderPercent := (renderFinish.Seconds() / finish.Seconds()) * 100.0
		fmt.Println("Frame time:", finish, ", Render time:", renderFinish, ", Rendering:", renderPercent, "%")

		if <-exit {
			break GameLoop
		}
	}
}

func main() {
	// Enables use of graphics calls on main os thread and goroutines together.
	/*
		go func() {
			fmt.Println("Start")
		}()
	*/
	mainthread.Run(run)
}
