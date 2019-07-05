package ego

import (
	// hlfstr
	//	"github.com/hlfstr/configurit"
	"github.com/hlfstr/logit"

	//SDL
	"github.com/veandco/go-sdl2/sdl"

	//SDL Libs
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
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

	ActiveScene scene

	DeltaTime float32
	FPS       int32

	// counter for FPS
	frames    int32
	closeChan chan bool
)

func init() {
	Log, _ = logit.Start(logit.TermLog())
	Log.Log(0, Version())

	wg = &sync.WaitGroup{}
	closeChan = make(chan bool, 1)
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

	//Start SDL_mixer
	if err = mix.Init(0); err != nil {
		return err
	}
	if err = mix.OpenAudio(44100, uint16(mix.DEFAULT_FORMAT), 2, 1024); err != nil {
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

	defer shutdown()

	isRunning = true

	ActiveScene.Start()

	wg.Add(3)
	go fpscounter()
	go graphics()
	go logic()

	events()

	wg.Wait()

	return nil
}

func fpscounter() {
	tick := time.NewTicker(time.Second)
counter:
	for {
		select {
		case <-tick.C:
			FPS = frames
			atomic.StoreInt32(&frames, 0)
		case <-closeChan:
			break counter
		}
	}
	wg.Done()
}

func graphics() {
	for isRunning {
		sdl.Do(func() {
			Renderer.SetDrawColor(ActiveScene.GetBackground())
			Renderer.Clear()
			Renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
		})
		ActiveScene.Draw()
		sdl.Do(func() {
			Renderer.Present()
			//sdl.Delay(1000 / 60)
		})
		atomic.AddInt32(&frames, 1)
	}
	wg.Done()
}

func logic() {
	//deltaTime calc
	var last uint64
	var now uint64

	sdl.Do(func() {
		last = sdl.GetPerformanceCounter()
	})

	for isRunning {
		sdl.Do(func() {
			now = sdl.GetPerformanceCounter()
		})
		DeltaTime = float32((now-last)/(sdl.GetPerformanceFrequency()/1000.0)) / 1000.0
		last = now

		ActiveScene.Update()
		sdl.Delay(1000 / 60)
	}

	wg.Done()
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
		}
	}
}

func Quit() {
	isRunning = false
	closeChan <- true
}

func shutdown() {
	Log.Logf(0, "Shutting down: %s", Name)
	close(closeChan)
	sdl.Do(func() {
		ActiveScene.Destroy()
		window.Destroy()
		Renderer.Destroy()
		img.Quit()
		mix.Quit()
		ttf.Quit()
		sdl.Quit()
		Log.Quit()
	})
}
