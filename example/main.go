package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	//"github.com/hlfstr/ego2d/component/pointer"
	"github.com/hlfstr/ego2d/ego"
)

/*
var (
	ptr *pointer.Pointer
)*/

func SDL_main() {
	ego.Create(1280, 720, ego.WINDOWED, "example")
	rand.Seed(time.Now().UnixNano())
	ego.AddScene("main", &exampleScene{})
	ego.AddScene("init", &initScene{})

	//	ptr = pointer.New()
	if b := ego.SetInitScene("init"); !b {
		fmt.Println("Oh shit")
		os.Exit(1)
	}

	err := ego.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Success!")
}

type exampleScene struct {
	ego.Scene
	state int
	ticks int

	objs []*object

	numObjects int
}

func (e *exampleScene) Update() {
	e.ticks++
	if e.ticks >= 60 {
		ego.Log.Logf(0, "FPS: %d | Delta: %f", ego.FPS, ego.DeltaTime)
		e.ticks = 0
	}
	//	ptr.Update()
	for i := range e.objs {
		e.objs[i].Update()
	}
}

func (e *exampleScene) Destroy() {
	for i := range e.objs {
		e.objs[i].Destroy()
	}
}

func (e *exampleScene) Start() bool {
	var err error
	e.Init(5)
	e.numObjects = 2000
	e.Background.B = 255
	e.objs = make([]*object, e.numObjects)
	for i := 0; i < e.numObjects; i++ {
		e.objs[i] = NewObject()
		e.objs[i].Start()
		if i%5 == 0 {
			e.objs[i].SetTexture("deepfryface")
			e.Add(0, e.objs[i])
		} else if i%5 == 1 {
			e.objs[i].SetTexture("beegyoshi")
			e.Add(1, e.objs[i])
		} else if i%5 == 2 {
			e.objs[i].SetTexture("burger")
			e.Add(2, e.objs[i])
		} else if i%5 == 3 {
			e.objs[i].SetTexture("pikachu")
			e.Add(3, e.objs[i])
		} else if i%5 == 4 {
			e.objs[i].SetTexture("poop")
			e.Add(4, e.objs[i])
		} else {
			e.objs[i].SetTexture("middlefinger")
			e.Add(1, e.objs[i])
		}
	}
	if err != nil {
		ego.Log.LogError(1, err)
		os.Exit(1)
	}

	return true
}

func main() {
	SDL_main()
}
