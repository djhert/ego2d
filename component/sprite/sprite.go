package sprite

import (
	"github.com/hlfstr/ego2d/ego"
	"github.com/hlfstr/ego2d/math"

	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	math.Transform

	TextureRegion math.Rect
	texture       *sdl.Texture
}

func New(name string) *Sprite {
	i := ego.GetTextureIndex(name)
	if i == -1 {
		ego.Log.Logf(0, "Unable to create Sprite from image: %s", name)
		return nil
	}
	out := &Sprite{math.NewTransform(0, 0, ego.Textures[i].W, ego.Textures[i].H, 0, math.NewPos2(ego.Textures[i].W/2, ego.Textures[i].H/2)), math.NewRect(0, 0, ego.Textures[i].W, ego.Textures[i].H), ego.Textures[i].Texture}
	return out
}

func (s *Sprite) Draw() {
	sdl.Do(func() {
		ego.Renderer.CopyEx(s.texture, s.TextureRegion.SDLRect(), s.SDLRect(), s.Rotation.Get(), s.Center.SDLPoint(), sdl.FLIP_NONE)
	})
}
