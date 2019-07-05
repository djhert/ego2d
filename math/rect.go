package math

import (
	"sync/atomic"

	"github.com/veandco/go-sdl2/sdl"
)

type Rect struct {
	Pos2
	w int32
	h int32
}

func NewRect(x, y, w, h int32) Rect {
	return Rect{Pos2: Pos2{x: x, y: y}, w: w, h: h}
}

func (p *Rect) SetW(w int32) {
	atomic.StoreInt32(&p.w, w)
}

func (p *Rect) SetH(h int32) {
	atomic.StoreInt32(&p.h, h)
}

func (p *Rect) SetWH(w, h int32) {
	atomic.StoreInt32(&p.w, w)
	atomic.StoreInt32(&p.h, h)
}

func (p *Rect) SetRect(s Rect) {
	atomic.StoreInt32(&p.w, s.GetW())
	atomic.StoreInt32(&p.h, s.GetH())
}

func (p *Rect) GetW() int32 {
	return atomic.LoadInt32(&p.w)
}

func (p *Rect) GetH() int32 {
	return atomic.LoadInt32(&p.h)
}

func (p *Rect) GetWH() (int32, int32) {
	return atomic.LoadInt32(&p.w), atomic.LoadInt32(&p.h)
}

func (p *Rect) AddW(w int32) {
	atomic.AddInt32(&p.w, w)
}

func (p *Rect) AddH(h int32) {
	atomic.AddInt32(&p.h, h)
}

func (p *Rect) AddWH(w, h int32) {
	atomic.AddInt32(&p.w, w)
	atomic.AddInt32(&p.h, h)
}

func (p *Rect) AddRect(s Rect) {
	atomic.AddInt32(&p.w, s.GetW())
	atomic.AddInt32(&p.h, s.GetH())
}

func (p *Rect) Equals(s Rect) bool {
	if p.GetX() == s.GetX() {
		if p.GetY() == s.GetY() {
			if p.GetW() == s.GetW() {
				if p.GetH() == s.GetH() {
					return true
				}
			}
		}
	}
	return false
}

func (p *Rect) SDLRect() *sdl.Rect {
	return &sdl.Rect{X: p.GetX(), Y: p.GetY(), W: p.GetW(), H: p.GetH()}
}
