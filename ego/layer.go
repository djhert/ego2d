package ego

import (
	"sync"
)

type layer struct {
	index   int
	objects []object

	wg *sync.WaitGroup
}

func newLayer(i int) *layer {
	var err error
	o := &layer{index: i, objects: make([]object, 0), wg: &sync.WaitGroup{}}
	if err != nil {
		Log.LogError(1, err)
		return nil
	}
	return o
}

func (l *layer) draw() {
	if len(l.objects) > 0 {
		l.wg.Add(len(l.objects))
		for i := range l.objects {
			go draw(l.objects[i])
		}
		l.wg.Wait()
	}
}
