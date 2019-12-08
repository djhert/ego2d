package text

import (
	"github.com/hlfstr/ego2d/ego"
	"github.com/hlfstr/ego2d/math"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	*ego.Object

	textureRegion math.Rect
	texture       *sdl.Texture
	font          *ttf.Font
	color         sdl.Color
	message       string
}

func New(msg string, ft *ttf.Font, col sdl.Color) *Text {
	out := &Text{
		Object:        ego.NewObject(),
		textureRegion: math.NewRect(0, 0, 0, 0),
		texture:       nil,
		message:       msg,
		font:          ft,
		color:         col,
	}
	out.SetMessage(msg)
	return out
}

func (s *Text) Draw() {
	if s.texture != nil {
		sdl.Do(func() {
			ego.Renderer.CopyEx(s.texture, s.textureRegion.SDLRect(), s.SDLRect(), s.Rotation.Get(), s.Center.SDLPoint(), sdl.FLIP_NONE)
		})
	}
}

func (s *Text) SetMessage(msg string) error {
	s.message = msg
	surface, err := s.font.RenderUTF8Solid(s.message, s.color)
	if err != nil {
		ego.Log.LogError(1, err)
		return err
	}
	sdl.Do(func() {
		s.texture, err = ego.Renderer.CreateTextureFromSurface(surface)
	})
	if err != nil {
		ego.Log.LogError(1, err)
		return err
	}
	s.SetTextureRegion(0, 0, surface.W, surface.H)
	s.SetWH(surface.W, surface.H)
	surface.Free()
	return nil
}

func (s *Text) SetTextureRegion(x, y, w, h int32) {
	s.textureRegion.SetXY(x, y)
	s.textureRegion.SetWH(w, h)
	s.Center.SetXY(w/2, h/2)
}

func (s *Text) Reset() {
	s.SetTextureRegion(0, 0, 0, 0)
	s.texture = nil
}

func (s *Text) Destroy() {
	s.texture.Destroy()
}
