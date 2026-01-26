package internal

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Slider struct {
	X     int32
	Y     int32
	Speed float64
	Title string

	sliderPos float64
	maxPos    float64
	direction int

	lineColour   rl.Color
	sliderColour rl.Color

	counter int
}

func NewSlider(x int32, y int32, speed float64, colour rl.Color, title string) *Slider {
	slider := &Slider{
		X:            x,
		Y:            y,
		Speed:        speed,
		Title:        title,
		sliderColour: colour,
	}
	slider.init()
	return slider
}

func (s *Slider) init() {
	s.sliderPos = 0
	s.maxPos = WindowW - Margins*2
	s.direction = 1
	s.counter = 0

	s.lineColour = rl.NewColor(255, 255, 255, 255)
}

func (s *Slider) Update() {
	s.sliderPos += s.Speed * float64(s.direction)

	if s.sliderPos > s.maxPos {
		excess := s.sliderPos - s.maxPos
		s.sliderPos = s.maxPos - excess
		s.counter++
		s.direction *= -1
	} else if s.sliderPos < 0 {
		s.sliderPos *= -1
		s.direction *= -1
	}
}

func (s *Slider) Draw(win *Window, graphics *Graphics) {
	rl.DrawText(s.Title, s.X, s.Y-FontSize-SliderBallRadius/2, FontSize, s.lineColour)
	w := rl.MeasureText(strconv.Itoa(s.counter), FontSize)
	rl.DrawText(strconv.Itoa(s.counter), s.X+int32(s.maxPos)-w, s.Y-FontSize-SliderBallRadius/2, FontSize, s.lineColour)

	rl.DrawLine(s.X, s.Y, s.X+int32(s.maxPos), s.Y, s.lineColour)
	rl.DrawCircle(s.X+int32(s.sliderPos), s.Y, SliderBallRadius, s.sliderColour)
}
