package pointer

import (
	"github.com/hlfstr/ego2d/component/sprite"
	"github.com/veandco/go-sdl2/sdl"
)

type Pointer struct {
	*sprite.Sprite
}

func New() *Pointer {
	return &Pointer{
		sprite.New(),
	}
}

func (p *Pointer) Start() {
	p.Focus(true)
	p.SetWH(16, 16)
}

func (p *Pointer) Update() {
	x, y, _ := sdl.GetMouseState()
	p.SetXY(x, y)
}

func (p *Pointer) Focus(b bool) {
	sdl.SetRelativeMouseMode(b)
}
