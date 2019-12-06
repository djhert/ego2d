package math

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

// Point is a thread-safe implementation of SDL_POINT
type Point struct {
	point *sdl.Point
	mx    *sync.RWMutex
}

func NewPoint(x, y int32) Point {
	return Point{
		point: new(sdl.Point),
		mx:    new(sync.RWMutex),
	}
}

func (p *Point) SetX(x int32) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.point.X = x
}

func (p *Point) SetY(y int32) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.point.Y = y
}

func (p *Point) SetXY(x, y int32) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.point.X = x
	p.point.Y = y
}

func (p *Point) SetPos(s Point) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.point.X = s.GetX()
	p.point.Y = s.GetY()
}

func (p *Point) GetX() int32 {
	p.mx.RLock()
	defer p.mx.RUnlock()
	return p.point.X
}

func (p *Point) GetY() int32 {
	p.mx.RLock()
	defer p.mx.RUnlock()
	return p.point.Y
}

func (p *Point) GetXY() (int32, int32) {
	p.mx.RLock()
	defer p.mx.RUnlock()
	return p.point.X, p.point.Y
}

func (p *Point) AddX(x int32) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.point.X += x
}

func (p *Point) AddY(y int32) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.point.Y += y
}

func (p *Point) AddXY(x, y int32) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.point.X += x
	p.point.Y += y
}

func (p *Point) AddPos(s Point) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.point.X += s.GetX()
	p.point.Y += s.GetY()
}

func (p Point) Equal(s Point) bool {
	p.mx.RLock()
	defer p.mx.RUnlock()
	if p.point.X == s.GetX() {
		if p.point.Y == s.GetY() {
			return true
		}
	}
	return false
}

func (p Point) SDLPoint() *sdl.Point {
	p.mx.RLock()
	defer p.mx.RUnlock()
	return p.point
}
