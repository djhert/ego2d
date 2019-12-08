package sprite

import (
	"github.com/hlfstr/ego2d/ego"
	"github.com/hlfstr/ego2d/math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	NONE int = iota
	HORIZONTAL
	VERTICAL
)

type Sprite struct {
	*ego.Object

	textureRegion math.Rect
	texture       *sdl.Texture
	flip          sdl.RendererFlip
}

func New() *Sprite {
	out := &Sprite{
		Object:        ego.NewObject(),
		textureRegion: math.NewRect(0, 0, 0, 0),
		texture:       nil,
		flip:          sdl.FLIP_NONE,
	}
	return out
}

func (s *Sprite) Draw() {
	if s.texture != nil {
		sdl.Do(func() {
			ego.Renderer.CopyEx(s.texture, s.textureRegion.SDLRect(), s.SDLRect(), s.Rotation.Get(), s.Center.SDLPoint(), s.flip)
		})
	}
}

func (s *Sprite) SetTexture(name string) {
	i, ok := ego.Textures[name]
	if !ok {
		ego.Log.Logf(0, "Unable to create Sprite from image: %s", name)
		return
	}
	s.texture = i.Texture
	s.SetTextureRegion(0, 0, i.W, i.H)
	s.SetWH(i.W, i.H)
}

func (s *Sprite) SetTextureRegion(x, y, w, h int32) {
	s.textureRegion.SetXY(x, y)
	s.textureRegion.SetWH(w, h)
	s.Center.SetXY(w/2, h/2)
}

func (s *Sprite) Reset() {
	s.SetTextureRegion(0, 0, 0, 0)
	s.texture = nil
}

func (s *Sprite) Destroy() {
	s.texture.Destroy()
}

func (s *Sprite) Flip(t int) {
	if t == HORIZONTAL {
		s.flip = sdl.FLIP_HORIZONTAL
	} else if t == VERTICAL {
		s.flip = sdl.FLIP_VERTICAL
	} else {
		s.flip = sdl.FLIP_NONE
	}
}
