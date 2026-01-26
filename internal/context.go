// app core logic and state

package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WindowW = 800
const WindowH = 600
const Margins = 80
const SliderBallRadius = 15
const BaseSpeed = 100.0
const FontSize = 25

func NewContext(win *Window, graphics *Graphics) *Context {
	context := &Context{win: win, graphics: graphics}
	context.Init()
	return context
}

type Context struct {
	win      *Window
	graphics *Graphics

	backgroundColour rl.Color
	sliders          []*Slider
}

// public in case game state needs to reset
func (c *Context) Init() {
	c.backgroundColour = rl.NewColor(35, 35, 35, 0)

	type info struct {
		colour rl.Color
		speed  float64
		title  string
	}

	infoList := []info{
		{
			colour: rl.NewColor(255, 0, 0, 255),
			speed:  2345,
			title:  "L1 Cache",
		},
		{
			colour: rl.NewColor(0, 255, 0, 255),
			speed:  894,
			title:  "L2 Cache",
		},
		{
			colour: rl.NewColor(0, 0, 255, 255),
			speed:  368,
			title:  "L3 Cache",
		},
		{
			colour: rl.NewColor(255, 255, 0, 255),
			speed:  48,
			title:  "RAM",
		},
		{
			colour: rl.NewColor(255, 0, 255, 255),
			speed:  7,
			title:  "SSD",
		},
		{
			colour: rl.NewColor(0, 255, 255, 255),
			speed:  0.2,
			title:  "HDD",
		},
	}

	spacing := WindowH / (len(infoList) + 1)
	ratio := BaseSpeed / infoList[0].speed

	for i, info_ := range infoList {
		c.sliders = append(c.sliders, NewSlider(
			int32(Margins),
			int32((i+1)*spacing),
			ratio*info_.speed,
			info_.colour,
			info_.title,
		))
	}

}

func (c *Context) Update() {
	for _, slider := range c.sliders {
		slider.Update()
	}
}

func (c *Context) Draw() {
	rl.ClearBackground(c.backgroundColour)

	for _, slider := range c.sliders {
		slider.Draw(c.win, c.graphics)
	}
}
