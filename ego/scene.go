package ego

import "fmt"

type scene interface {
	Start() bool
	Update()
	Destroy()

	Add(int, object) error

	getBackground() (uint8, uint8, uint8, uint8)
	draw()
}

type Scene struct {
	Background Color
	scene

	numLayers int
	layers    []*layer
}

func (s *Scene) Init(n int) {
	s.setupLayers(n)
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
