package ego

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	//SDL
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"

	//SDL Libs
	"github.com/veandco/go-sdl2/img"
)

type texture struct {
	Texture *sdl.Texture
	W, H    int32
}

var (
	AssetsLoaded bool

	Textures   map[string]texture
	textureExt *regexp.Regexp

	Music    map[string]*mix.Music
	musicExt *regexp.Regexp

	Sounds   map[string]*mix.Chunk
	soundExt *regexp.Regexp
)

func init() {
	Textures = make(map[string]texture)

	Music = make(map[string]*mix.Music)

	Sounds = make(map[string]*mix.Chunk)

	textureExt = regexp.MustCompile(".(jpg|gif|png|bmp|tiff)$")
	musicExt = regexp.MustCompile(".(wav|mp3|ogg)$")
	soundExt = regexp.MustCompile(".(wav|ogg)$")
}

//Textures
func LoadTexture(name, path string) error {
	if _, ok := Textures[name]; ok {
		return fmt.Errorf("Duplicate texture: %s", name)
	}
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
	Textures[name] = save

	return nil
}

func LoadAssets() {
	go func() {
		if !AssetsLoaded {
			go aloadTextures()
			for !AssetsLoaded {
				if Audio.IsInit {
					go aloadMusic()
					go aloadSounds()
					AssetsLoaded = true
				}
				if !Audio.CanInit {
					Log.Log(2, "No Audio will be loaded")
					AssetsLoaded = true
				}
			}
		}
	}()
}

func aloadTextures() {
	aDir := "assets/textures/"
	filepath.Walk(aDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			Log.LogError(2, err)
			return nil
		}
		if !info.IsDir() {
			if textureExt.MatchString(info.Name()) {
				err = LoadTexture(strings.TrimSuffix(strings.TrimPrefix(path, aDir), filepath.Ext(path)), path)
			}
		}
		if err != nil {
			Log.LogError(2, err)
		}
		return nil
	})
	return
}

//Music
func LoadMusic(name, path string) error {
	if _, ok := Music[name]; ok {
		return fmt.Errorf("Duplicate texture: %s", name)
	}
	save, err := mix.LoadMUS(path)
	if err != nil {
		return err
	}
	Music[name] = save

	return nil
}

func aloadMusic() {
	aDir := "assets/music/"
	filepath.Walk(aDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			Log.LogError(2, err)
			return nil
		}
		if !info.IsDir() {
			if musicExt.MatchString(info.Name()) {
				err = LoadMusic(strings.TrimSuffix(strings.TrimPrefix(path, aDir), filepath.Ext(path)), path)
			}
		}
		if err != nil {
			Log.LogError(2, err)
		}
		return nil
	})
	return
}

///Sounds
func LoadSound(name, path string) error {
	if _, ok := Sounds[name]; ok {
		return fmt.Errorf("Duplicate texture: %s", name)
	}
	save, err := mix.LoadWAV(path)
	if err != nil {
		return err
	}
	Sounds[name] = save

	return nil
}

func aloadSounds() {
	aDir := "assets/sounds/"
	filepath.Walk(aDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			Log.LogError(2, err)
			return nil
		}
		if !info.IsDir() {
			if soundExt.MatchString(info.Name()) {
				err = LoadSound(strings.TrimSuffix(strings.TrimPrefix(path, aDir), filepath.Ext(path)), path)
			}
		}
		if err != nil {
			Log.LogError(2, err)
			return nil
		}
		return nil
	})
	return
}

func destroyAssets() {
	for i := range Textures {
		Textures[i].Texture.Destroy()
	}
}
