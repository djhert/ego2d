package math

import (
	"sync/atomic"

	"github.com/veandco/go-sdl2/sdl"
)

type Pos2 struct {
	x int32
	y int32
}

func NewPos2(x, y int32) Pos2 {
	return Pos2{x: x, y: y}
}

func (p *Pos2) SetX(x int32) {
	atomic.StoreInt32(&p.x, x)
}

func (p *Pos2) SetY(y int32) {
	atomic.StoreInt32(&p.y, y)
}

func (p *Pos2) SetXY(x, y int32) {
	atomic.StoreInt32(&p.x, x)
	atomic.StoreInt32(&p.y, y)
}

func (p *Pos2) SetPos(s Pos2) {
	atomic.StoreInt32(&p.x, s.GetX())
	atomic.StoreInt32(&p.y, s.GetY())
}

func (p *Pos2) GetX() int32 {
	return atomic.LoadInt32(&p.x)
}

func (p *Pos2) GetY() int32 {
	return atomic.LoadInt32(&p.y)
}

func (p *Pos2) GetXY() (int32, int32) {
	return atomic.LoadInt32(&p.x), atomic.LoadInt32(&p.y)
}

func (p *Pos2) AddX(x int32) {
	atomic.AddInt32(&p.x, x)
}

func (p *Pos2) AddY(y int32) {
	atomic.AddInt32(&p.y, y)
}

func (p *Pos2) AddXY(x, y int32) {
	atomic.AddInt32(&p.x, x)
	atomic.AddInt32(&p.y, y)
}

func (p *Pos2) AddPos(s Pos2) {
	atomic.AddInt32(&p.x, s.GetX())
	atomic.AddInt32(&p.y, s.GetY())
}

func (p Pos2) Equal(s Pos2) bool {
	if p.GetX() == s.GetX() {
		if p.GetY() == s.GetY() {
			return true
		}
	}
	return false
}

func (p Pos2) SDLPoint() *sdl.Point {
	return &sdl.Point{X: p.GetX(), Y: p.GetY()}
}
