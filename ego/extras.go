package ego

import "sync"

var (
	exObjects     []extras
	extraUpdateWG *sync.WaitGroup
	extraDrawWG   *sync.WaitGroup

	extrasUsed bool
)

func init() {
	exObjects = make([]extras, 0)

	extraUpdateWG = new(sync.WaitGroup)
	extraDrawWG = new(sync.WaitGroup)
}

func RegisterExtra(e extras) {
	e.setUpdatewg(extraUpdateWG)
	e.setDrawwg(extraDrawWG)
	exObjects = append(exObjects, e)

	if !extrasUsed {
		extrasUsed = true
	}
}

type extras interface {
	Start()
	Update()
	Draw()

	setUpdatewg(s *sync.WaitGroup)
	setDrawwg(s *sync.WaitGroup)
	doneUdwg()
	doneDwwg()
}

type ExtraObject struct {
	extras
	updateWG *sync.WaitGroup
	drawWG   *sync.WaitGroup
}

func (e *ExtraObject) setUpdatewg(s *sync.WaitGroup) {
	e.setUpdatewg(s)
}

func (e *ExtraObject) setDrawwg(s *sync.WaitGroup) {
	e.setDrawwg(s)
}

func (e *ExtraObject) doneUdwg() {
	e.updateWG.Done()
}

func (e *ExtraObject) doneDwwg() {
	e.drawWG.Done()
}

func updateExtra(e extras) {
	e.Update()
	e.doneUdwg()
}

func drawExtra(e extras) {
	e.Draw()
	e.doneDwwg()
}

func updateExtras() {
	if extrasUsed {
		extraUpdateWG.Add(len(exObjects))
		for i := range exObjects {
			go updateExtra(exObjects[i])
		}
		extraUpdateWG.Wait()
	}
}

func drawExtras() {
	if extrasUsed {
		extraDrawWG.Add(len(exObjects))
		for i := range exObjects {
			go drawExtra(exObjects[i])
		}
		extraDrawWG.Wait()
	}
}
