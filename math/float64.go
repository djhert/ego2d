package math

import (
	"math"
	"sync/atomic"
)

type AtomicFloat struct {
	v uint64
}

func NewAtomicFloat(f float64) AtomicFloat {
	return AtomicFloat{math.Float64bits(f)}
}

func (f *AtomicFloat) Set(n float64) {
	atomic.StoreUint64(&f.v, math.Float64bits(n))
}

func (f *AtomicFloat) Get() float64 {
	return math.Float64frombits(atomic.LoadUint64(&f.v))
}

func (f *AtomicFloat) Add(n float64) float64 {
	o := f.Get() + n
	f.Set(o)
	return o
}

