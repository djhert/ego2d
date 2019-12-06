package ego

import (
	"github.com/hlfstr/ego2d/math"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	displays      []display
	activeDisplay int
	WindowMode    WinMode
)

func init() {
	displays = make([]display, 0)
	activeDisplay = 0
}

type WinMode int32

const (
	FULLSCREEN WinMode = iota
	WINDOWED
	BORDERLESS_FULLSCREEN
)

func (w WinMode) String() string {
	switch w {
	case FULLSCREEN:
		return "FULLSCREEN"
	case WINDOWED:
		return "WINDOWED"
	case BORDERLESS_FULLSCREEN:
		return "BORDERLESS_FULLSCREEN"
	default:
		return "UNKNOWN"
	}
}

func (w WinMode) mode() uint32 {
	if w == FULLSCREEN {
		return sdl.WINDOW_FULLSCREEN
	}
	return 0
}

type display struct {
	math.Rect
	name string
}

func getDisplays() error {
	d, err := sdl.GetNumVideoDisplays()
	if err != nil {
		Log.LogError(2, err)
		return err
	}
	for i := 0; i < d; i++ {
		s, err := sdl.GetDisplayName(i)
		if err != nil {
			Log.LogError(1, err)
			return err
		}
		r, err := sdl.GetDisplayBounds(i)
		if err != nil {
			Log.LogError(1, err)
			return err
		}
		outd := display{
			Rect: math.NewRect(r.X, r.Y, r.W, r.H),
			name: s,
		}
		displays = append(displays, outd)
	}

	return nil
}

func fullscreen() error {
	err := window.SetFullscreen(WindowMode.mode())
	if err != nil {
		Log.LogError(2, err)
		return err
	}

	return nil
}

//borderlessFullscreen forces the game to still be a window
//but takes up the entire usable screen
//forces resolution to be the size of the display
func borderlessFullscreen() error {
	err := window.SetFullscreen(WindowMode.mode())
	if err != nil {
		Log.LogError(2, err)
		return err
	}
	window.SetBordered(false)
	window.SetPosition(displays[activeDisplay].GetXY())
	window.Raise()

	return nil
}

func windowed() error {
	err := window.SetFullscreen(WindowMode.mode())
	if err != nil {
		Log.LogError(2, err)
		return err
	}
	window.SetPosition(displays[activeDisplay].GetX()+(displays[activeDisplay].GetW()-Width)/2, displays[activeDisplay].GetY()+(displays[activeDisplay].GetH()-Height)/2)
	window.SetBordered(true)
	window.Raise()

	return nil
}

func SetWindowMode(m WinMode) error {
	var err error
	WindowMode = m

	switch WindowMode {
	case FULLSCREEN:
		sdl.Do(func() {
			err = fullscreen()
		})
	case WINDOWED:
		sdl.Do(func() {
			err = windowed()
		})
	case BORDERLESS_FULLSCREEN:
		SetWindowSize(displays[activeDisplay].GetWH())
		sdl.Do(func() {
			err = borderlessFullscreen()
		})
	default:
		Log.Logf(1, "No WindowMode with value: %d | Defaulting to Windowed WindowMode", m)
		WindowMode = WINDOWED
		sdl.Do(func() {
			err = windowed()
		})
	}

	return err
}

func SetWindowSize(w, h int32) {
	Width = w
	Height = h
	sdl.Do(func() {
		window.SetSize(Width, Height)
	})
}
