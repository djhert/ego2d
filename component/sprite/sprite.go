package sprite

import (
	"github.com/hlfstr/ego2d/ego"
	"github.com/hlfstr/ego2d/math"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/veandco/go-sdl2/img"
)

type Sprite struct {
	math.Rect

	Position math.Transform

	texture *sdl.Texture
}

func LoadPNG(name string) (*Sprite, error) {
	temp, err := img.Load(name)
	if err != nil {
		return nil, err
	}
	out := &Sprite{math.NewRect(0, 0, temp.W, temp.H), math.NewTransform(0, 0, temp.W, temp.H, 0, math.NewPos2(temp.W/2, temp.H/2)), nil}

	sdl.Do(func() {
		out.texture, err = ego.Renderer.CreateTextureFromSurface(temp)
	})
	if err != nil {
		return nil, err
	}
	return out, err
}

func (s *Sprite) Destroy() {
	s.texture.Destroy()
}

func (s *Sprite) Draw() {
	sdl.Do(func() {
		ego.Renderer.CopyEx(s.texture, s.SDLRect(), s.Position.SDLRect(), s.Position.Rotation.Get(), s.Position.Center.SDLPoint(), sdl.FLIP_NONE)
	})
}
