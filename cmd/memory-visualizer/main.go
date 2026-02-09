package main

import (
	"memory-visualizer/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetTraceLogLevel(rl.LogError)

	win, err := internal.NewWin(internal.WindowW, internal.WindowH, "memoryVisualizer", 0)
	if err != nil {
		panic(err)
	}

	graphics, err := internal.NewGraphicsHandler("assets")
	if err != nil {
		panic(err)
	}

	context := internal.NewContext(win, graphics)

	win.StartLoop(context)
}
