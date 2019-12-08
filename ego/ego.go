package ego

import (
	// hlfstr
	//	"github.com/hlfstr/configurit"
	"fmt"

	"github.com/hlfstr/logit"

	//SDL
	"github.com/veandco/go-sdl2/sdl"

	//SDL Libs
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/ttf"

	//stdlib
	"sync"
	"sync/atomic"
	"time"
)

var (
	Log *logit.Logger
	//	Conf *configurit.Conf

	Width, Height int32
	FullScreen    bool
	Name          string

	window   *sdl.Window
	Renderer *sdl.Renderer

	wg *sync.WaitGroup

	isRunning bool

	DeltaTime float32
	FPS       int32

	fpsTarget uint32

	// counter for FPS
	frames    int32
	closeChan chan bool
)

func init() {
	Log, _ = logit.Start(logit.TermLog())

	wg = &sync.WaitGroup{}
	closeChan = make(chan bool, 1)

	fpsTarget = 60
}

func Create(w, h int32, wm WinMode, n string) {
	Width = w
	Height = h
	Name = n
	WindowMode = wm
}

/*
func CreateFrom(filepath string) error {
	return initSDL()
}
*/

func Run() error {
	var err error
	sdl.Main(func() {
		err = run()
	})
	return err
}

func initSDL() error {
	var err error
	// Start SDL
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return err
	}

	//Create Window
	window, err = sdl.CreateWindow(Name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, Width, Height, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}

	getDisplays()
	window.SetResizable(false)

	//Create Renderer
	Renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}

	//Start SDL_imhage
	img.Init(img.INIT_PNG)

	//Start SDL_TTF
	if err = ttf.Init(); err != nil {
		return err
	}

	return nil
}

func run() error {
	var err error
	sdl.Do(func() {
		err = initSDL()
	})
	if err != nil {
		return err
	}
	SetWindowMode(WindowMode)

	wg.Add(1)
	go Audio.run()

	defer shutdown()

	isRunning = true

	if ActiveScene == nil {
		return fmt.Errorf("No active scene set")
	}

	sceneStarted = ActiveScene.Start()

	wg.Add(3)
	go fpscounter()
	go graphics()
	go logic()

	events()

	wg.Wait()

	return nil
}

func fpscounter() {
	defer wg.Done()

	tick := time.NewTicker(time.Second)
counter:
	for {
		select {
		case <-tick.C:
			FPS = frames
			atomic.StoreInt32(&frames, 0)
		case <-closeChan:
			tick.Stop()
			break counter
		}
	}
}

func graphics() {
	defer wg.Done()
	var startTicks uint32
	var endTicks uint32
	var deltaTicks uint32

	for isRunning {
		startTicks = sdl.GetTicks()

		sdl.Do(func() {
			Renderer.SetDrawColor(ActiveScene.getBackground())
			Renderer.Clear()
			Renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
		})

		if sceneStarted {
			ActiveScene.draw()
		}

		drawExtras()

		sdl.Do(func() {
			Renderer.Present()
		})

		endTicks = sdl.GetTicks()
		deltaTicks = endTicks - startTicks
		if deltaTicks < 1000/fpsTarget {
			sdl.Delay((1000 / fpsTarget) - deltaTicks)
		}
		atomic.AddInt32(&frames, 1)
	}
}

func logic() {
	defer wg.Done()

	var startTicks uint32
	var endTicks uint32
	var deltaTicks uint32

	var lastTicks uint32

	for isRunning {
		startTicks = sdl.GetTicks()

		DeltaTime = float32(startTicks-lastTicks) / 1000.0
		lastTicks = startTicks

		if sceneStarted {
			ActiveScene.Update()
		}

		updateExtras()

		endTicks = sdl.GetTicks()
		deltaTicks = endTicks - startTicks

		if deltaTicks < 1000/fpsTarget {
			sdl.Delay((1000 / fpsTarget) - deltaTicks)
		}
	}
}

func events() {
	for isRunning {
		event := sdl.WaitEvent()
		// Found very rarely the event was nil
		// if so, wait and move to next event
		if event == nil {
			sdl.Delay(500)
			continue
		}
		switch event.(type) {
		case *sdl.QuitEvent:
			Quit()
		case *sdl.MouseMotionEvent:
			ActiveScene.MouseMotion(event.(*sdl.MouseMotionEvent))
		case *sdl.MouseButtonEvent:
			ActiveScene.MouseButton(event.(*sdl.MouseButtonEvent))
		case *sdl.MouseWheelEvent:
			ActiveScene.MouseWheel(event.(*sdl.MouseWheelEvent))
		case *sdl.KeyboardEvent:
			ActiveScene.Keyboard(event.(*sdl.KeyboardEvent))
		case *sdl.WindowEvent:
			we := event.(*sdl.WindowEvent)
			switch we.Event {
			case 11:
				fallthrough
			case 13:
				ActiveScene.WindowLoseFocus(we)
			case 10:
				fallthrough
			case 12:
				ActiveScene.WindowGainFocus(we)
			}
			//fmt.Printf("%T | %d\n", event, event.(*sdl.WindowEvent).Event)
		}
	}
}

func Quit() {
	isRunning = false
	close(closeChan)
}

func shutdown() {
	Log.Logf(0, "Shutting down: %s", Name)
	sdl.Do(func() {
		ActiveScene.Destroy()
		destroyAssets()
		window.Destroy()
		Renderer.Destroy()
		img.Quit()
		ttf.Quit()
		sdl.Quit()
		Log.Quit()
	})
}
