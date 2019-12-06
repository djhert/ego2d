package ego

import (
	"sync/atomic"

	"github.com/veandco/go-sdl2/mix"
)

var (
	Audio audio
)

const (
	playmusic uint8 = iota
	stopmusic
	pausemusic
	loadmusic
	musicvolume
	playsound
)

func init() {
	Audio = audio{}
	Audio.achan = make(chan acmd, 8)
	Audio.CanInit = true
}

type audio struct {
	Current     *mix.Music
	musicVolume int32
	soundVolume int32

	IsInit  bool
	CanInit bool

	achan chan acmd
}

type acmd struct {
	cmd uint8
	v   interface{}
}

func (a *audio) PlayMusic() {
	if a.IsInit {
		a.achan <- acmd{cmd: playmusic, v: nil}
	}
}

func (a *audio) StopMusic() {
	if a.IsInit {
		a.achan <- acmd{cmd: stopmusic, v: nil}
	}
}

func (a *audio) PauseMusic() {
	if a.IsInit {
		a.achan <- acmd{cmd: pausemusic, v: nil}
	}
}

func (a *audio) PaySound(name string) {
	if a.IsInit {
		a.achan <- acmd{cmd: playsound, v: name}
	}
}

func (a *audio) SetMusicVolume(v int32) {
	if a.IsInit {
		a.achan <- acmd{cmd: musicvolume, v: v}
	}
}

func (a *audio) GetMusicVolume() int32 {
	return atomic.LoadInt32(&a.musicVolume)
}

func (a *audio) destroy() {
	close(a.achan)
	for i := range Music {
		Music[i].Free()
	}
	for i := range Sounds {
		Sounds[i].Free()
	}
	mix.CloseAudio()
	mix.Quit()
}

func (a *audio) run() {
	defer wg.Done()
	defer a.destroy()
loop:
	for {
		//Start SDL_mixer
		if !a.IsInit {
			var err error
			if err = mix.Init(0); err != nil {
				Log.Logf(2, "Unable to start Audio System: %s", err.Error())
			}
			if err = mix.OpenAudio(44100, uint16(mix.DEFAULT_FORMAT), 2, 4096); err != nil {
				Log.Logf(2, "Unable to open Audio Channels: %s", err.Error())
			}
			if err != nil {
				a.IsInit = false
				a.CanInit = false
				break loop
			} else {
				a.IsInit = true
			}
		}

		select {
		case i := <-a.achan:
			a.handleCmd(i)
		case <-closeChan:
			break loop
		}
	}
}

func (a *audio) handleCmd(cmd acmd) {
	switch cmd.cmd {
	case playmusic:
		a.Current.Play(0)
	case stopmusic:
		mix.HaltMusic()
	case pausemusic:
		mix.PauseMusic()
	case musicvolume:
		atomic.StoreInt32(&a.musicVolume, cmd.v.(int32))
		mix.VolumeMusic(int(atomic.LoadInt32(&a.musicVolume)))
	case playsound:
		snd, ok := Sounds[cmd.v.(string)]
		if ok {
			snd.Play(-1, 0)
		}
	}
}
