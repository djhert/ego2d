package pointer

import (
	"github.com/hlfstr/ego2d/component/sprite"
	"github.com/hlfstr/ego2d/ego"
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
	//	sdl.SetRelativeMouseMode(true)
	ego.ExtraDraw(p.Draw)
	p.SetWH(16, 16)
}

func (p *Pointer) Update() {
	x, y, _ := sdl.GetMouseState()
	p.SetXY(x, y)
}
