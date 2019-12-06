package math

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type Rect struct {
	rect *sdl.Rect
	mx   *sync.RWMutex
}

func NewRect(x, y, w, h int32) Rect {
	return Rect{
		rect: new(sdl.Rect),
		mx:   new(sync.RWMutex),
	}
}

func (r *Rect) SetX(x int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.X = x
}

func (r *Rect) SetY(y int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.Y = y
}

func (r *Rect) SetXY(x, y int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.X = x
	r.rect.Y = y
}

func (r *Rect) SetPos(s Rect) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.X = s.GetX()
	r.rect.Y = s.GetY()
}

func (r *Rect) GetX() int32 {
	r.mx.RLock()
	defer r.mx.RUnlock()
	return r.rect.X
}

func (r *Rect) GetY() int32 {
	r.mx.RLock()
	defer r.mx.RUnlock()
	return r.rect.Y
}

func (r *Rect) GetXY() (int32, int32) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	return r.rect.X, r.rect.Y
}

func (r *Rect) AddX(x int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.X += x
}

func (r *Rect) AddY(y int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.Y += y
}

func (r *Rect) AddXY(x, y int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.X += x
	r.rect.Y += y
}

func (r *Rect) SetW(w int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.W = w
}

func (r *Rect) SetH(h int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.H = h
}

func (r *Rect) SetWH(w, h int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.W = w
	r.rect.H = h
}

func (r *Rect) SetXYWH(x, y, w, h int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.X = x
	r.rect.Y = y
	r.rect.W = w
	r.rect.H = h
}

func (r *Rect) SetRect(s Rect) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.X, r.rect.Y = s.GetXY()
	r.rect.W, r.rect.H = s.GetWH()
}

func (r *Rect) GetW() int32 {
	r.mx.RLock()
	defer r.mx.RUnlock()
	return r.rect.W
}

func (r *Rect) GetH() int32 {
	r.mx.RLock()
	defer r.mx.RUnlock()
	return r.rect.H
}

func (r *Rect) GetWH() (int32, int32) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	return r.rect.W, r.rect.H
}

func (r *Rect) AddW(w int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.W += w
}

func (r *Rect) AddH(h int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.H += h
}

func (r *Rect) AddWH(w, h int32) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rect.W += w
	r.rect.H += h
}

func (r *Rect) Equals(s Rect) bool {
	r.mx.RLock()
	defer r.mx.RUnlock()
	if r.GetX() == s.GetX() {
		if r.GetY() == s.GetY() {
			if r.GetW() == s.GetW() {
				if r.GetH() == s.GetH() {
					return true
				}
			}
		}
	}
	return false
}

func (r *Rect) SDLRect() *sdl.Rect {
	r.mx.RLock()
	defer r.mx.RUnlock()
	return r.rect
}
