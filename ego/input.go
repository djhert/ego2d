package ego

import (
	"github.com/veandco/go-sdl2/sdl"
)

type inputs interface {
	MouseMotion(*sdl.MouseMotionEvent) bool
	MouseButton(*sdl.MouseButtonEvent) bool
	MouseWheel(*sdl.MouseWheelEvent) bool
	Keyboard(*sdl.KeyboardEvent) bool
}
