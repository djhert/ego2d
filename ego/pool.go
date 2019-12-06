package ego

import "sync"

type objectPool struct {
	available []object
	used      []object
	mx        *sync.Mutex
	create    func() object
	Count     int
}

func NewPool(c int, create func() object) *objectPool {
	return &objectPool{
		available: make([]object, 0),
		used:      make([]object, 0),
		mx:        new(sync.Mutex),
		create:    create,
		Count:     0,
	}
}

func (p *objectPool) Fill(n int) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.Count = n
	for i := 0; i < p.Count; i++ {
		p.available = append(p.available, p.create())
	}
}

func (p *objectPool) Get() object {
	p.mx.Lock()
	defer p.mx.Unlock()
	if len(p.available) > 0 {
		obj := p.available[0]
		p.available = p.available[1:]
		p.used = append(p.used, obj)
		return obj
	}
	obj := p.create()
	p.used = append(p.used, obj)
	p.Count++
	return obj
}

func (p *objectPool) Return(obj object) {
	p.mx.Lock()
	defer p.mx.Unlock()
	obj.Reset()
	i := p.findUsedIndex(obj)
	if i == -1 {
		obj.Destroy()
		return
	}
	p.used = append(p.used[:i], p.used[i+1:]...)
	p.available = append(p.available, obj)
}

func (p *objectPool) findUsedIndex(obj object) int {
	for i := range p.used {
		if p.used[i] == obj {
			return i
		}
	}
	return -1
}
