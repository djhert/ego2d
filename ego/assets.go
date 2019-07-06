package ego

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	//SDL
	"github.com/veandco/go-sdl2/sdl"

	//SDL Libs
	"github.com/veandco/go-sdl2/img"
)

type texture struct {
	Texture *sdl.Texture
	W, H    int32
}

var (
	AssetAutoload bool

	Textures   []texture
	textureKey []string
	textureExt *regexp.Regexp
)

func init() {
	Textures = make([]texture, 0)
	textureKey = make([]string, 0)
	AssetAutoload = true

	textureExt = regexp.MustCompile(".(jpg|gif|png)$")
}

func LoadTexture(name, path string) error {
	surface, err := img.Load(path)
	if err != nil {
		return err
	}
	save := texture{Texture: nil, W: surface.W, H: surface.H}
	sdl.Do(func() {
		save.Texture, err = Renderer.CreateTextureFromSurface(surface)
	})
	if err != nil {
		return err
	}
	Textures = append(Textures, save)
	textureKey = append(textureKey, name)

	return nil
}

func GetTextureIndex(name string) int32 {
	for i := range textureKey {
		if textureKey[i] == name {
			return int32(i)
		}
	}
	return -1
}

func autoload() {
	aloadTextures()
}

func aloadTextures() error {
	var e error
	aDir := "assets/textures/"
	e = filepath.Walk(aDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if textureExt.MatchString(info.Name()) {
				return LoadTexture(strings.TrimSuffix(strings.TrimPrefix(path, aDir), filepath.Ext(path)), path)
			}
		}
		return nil
	})
	return e
}
