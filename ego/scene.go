package ego

type scene interface {
	GetBackground() (uint8, uint8, uint8, uint8)
	Start()
	Update()
	Draw()
	Destroy()
}

type Scene struct {
	Color
	scene
}

func (s *Scene) GetBackground() (uint8, uint8, uint8, uint8) {
	return s.R, s.G, s.B, s.A
}
