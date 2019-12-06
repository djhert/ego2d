package math

import (
	"sync"
)

type AtomicFloat struct {
	v  float64
	mx *sync.RWMutex
}

func NewAtomicFloat(f float64) AtomicFloat {
	return AtomicFloat{
		v:  f,
		mx: new(sync.RWMutex),
	}
}

func (f *AtomicFloat) Set(n float64) {
	f.mx.Lock()
	defer f.mx.Unlock()
	f.v = n
}

func (f *AtomicFloat) Get() float64 {
	f.mx.RLock()
	defer f.mx.RUnlock()
	return f.v
}

func (f *AtomicFloat) Add(n float64) {
	f.mx.Lock()
	defer f.mx.Unlock()
	f.v += n
}
