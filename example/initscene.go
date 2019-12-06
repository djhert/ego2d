package main

import (
	"github.com/hlfstr/ego2d/ego"
)

type initScene struct {
	ego.Scene

	ticks  int
	loaded bool
}

func (i *initScene) Start() bool {
	i.Init(1)
	ego.Log.Logf(0, "Hello init scene!  Loading assets...")
	ego.LoadTexture("middlefinger", "assets/textures/middlefinger.png")
	//ptr.SetTexture("middlefinger")
	//ptr.Start()
	ego.LoadAssets()
	i.Background.B = 129
	return true
}

func (i *initScene) Update() {
	//ptr.Update()
	if i.ticks > 100 && !i.loaded {
		if ego.AssetsLoaded {
			ego.Log.Logf(0, "Assets are loaded, Preparing main scene")
			ego.SetNextScene("main")
			i.loaded = true
		}
	}
	if i.ticks > 150 && i.loaded {
		ego.Log.Log(0, "Setting the next scene!")
		ego.NextScene()
	}
	i.ticks++
}

func (i *initScene) Destroy() {}
