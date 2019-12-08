package ego

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type scene interface {
	Start() bool
	Update()
	Destroy()

	Add(int, object) error

	getBackground() (uint8, uint8, uint8, uint8)
	draw()

	MouseMotion(*sdl.MouseMotionEvent)
	MouseButton(*sdl.MouseButtonEvent)
	MouseWheel(*sdl.MouseWheelEvent)
	Keyboard(*sdl.KeyboardEvent)
	WindowGainFocus(*sdl.WindowEvent)
	WindowLoseFocus(*sdl.WindowEvent)
}

type Scene struct {
	Background Color
	scene

	numLayers int
	layers    []*layer

	input []inputs
}

func (s *Scene) Init(n int) {
	s.setupLayers(n)
	s.input = make([]inputs, 0)
}

func (s *Scene) getBackground() (uint8, uint8, uint8, uint8) {
	return s.Background.R, s.Background.G, s.Background.B, s.Background.A
}

func (s *Scene) draw() {
	for i := range s.layers {
		s.layers[i].draw()
	}
}

func (l *Scene) setupLayers(n int) {
	l.numLayers = n
	l.layers = make([]*layer, l.numLayers)
	for i := range l.layers {
		l.layers[i] = newLayer(i)
	}
}

func (l *Scene) Add(i int, obj object) error {
	if i >= l.numLayers {
		return fmt.Errorf("No layer: %d", i)
	}
	obj.setWaitGroup(l.layers[i].wg)
	l.layers[i].objects = append(l.layers[i].objects, obj)
	return nil
}

func (m *Scene) AddInput(i inputs) {
	m.input = append(m.input, i)
}

func (m *Scene) Keyboard(e *sdl.KeyboardEvent) {
	for i := range m.input {
		found := m.input[i].Keyboard(e)
		if found {
			return
		}
	}
}

func (m *Scene) MouseMotion(e *sdl.MouseMotionEvent) {
	for i := range m.input {
		found := m.input[i].MouseMotion(e)
		if found {
			return
		}
	}
}

func (m *Scene) MouseButton(e *sdl.MouseButtonEvent) {
	for i := range m.input {
		found := m.input[i].MouseButton(e)
		if found {
			return
		}
	}
}

func (m *Scene) MouseWheel(e *sdl.MouseWheelEvent) {
	for i := range m.input {
		found := m.input[i].MouseWheel(e)
		if found {
			return
		}
	}
}

func (m *Scene) WindowGainFocus(e *sdl.WindowEvent) {

}

func (m *Scene) WindowLoseFocus(e *sdl.WindowEvent) {

}
