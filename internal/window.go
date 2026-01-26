package internal

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewWin(w int32, h int32, name string, fps int32) (*Window, error) {
	if w == 0 && h == 0 {
		rl.SetConfigFlags(rl.FlagFullscreenMode)
	} else if w <= 0 || h <= 0 {
		return nil, errors.New("bad window size")
	}

	rl.InitWindow(w, h, name)
	rl.SetTargetFPS(fps)

	return &Window{
		w: int32(rl.GetScreenWidth()),
		h: int32(rl.GetScreenHeight()),
	}, nil
}

type Window struct {
	w int32
	h int32
}

func (win *Window) Close() {
	rl.CloseWindow()
}

func (win *Window) StartLoop(context *Context) {
	for !rl.WindowShouldClose() {
		// calculate game logic
		context.Update()

		rl.BeginDrawing()

		// draw game logic
		context.Draw()

		rl.EndDrawing()
	}
}
