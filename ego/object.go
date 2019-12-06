package ego

import (
	"sync"

	"github.com/hlfstr/ego2d/math"
)

type object interface {
	//Required
	Start()
	Update()
	Draw()
	Destroy()

	//Optional
	Reset()

	//Internal only
	setWaitGroup(*sync.WaitGroup)
	doneDraw()
}

type Object struct {
	object
	math.Transform

	wg *sync.WaitGroup
}

func NewObject() *Object {
	return &Object{
		Transform: math.NewTransform(0, 0, 0, 0, 0, math.NewPoint(0, 0)),
		wg:        nil,
	}
}

func (o *Object) setWaitGroup(w *sync.WaitGroup) {
	o.wg = w
}

func (o *Object) doneDraw() {
	o.wg.Done()
}

func (o *Object) Reset() {}

func draw(o object) {
	o.Draw()
	o.doneDraw()
}
