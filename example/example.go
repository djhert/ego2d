package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/hlfstr/ego2d/ego"
)

func main() {
	ego.Create(1920, 1080, ego.FULLSCREEN, "example")
	rand.Seed(time.Now().UnixNano())
	ego.ActiveScene = &exampleScene{}
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

	objs []object

	numObjects int
}

func (e *exampleScene) Update() {
	for i := range e.objs {
		e.objs[i].Update()
	}
}

func (e *exampleScene) Draw() {
	e.ticks++
	if e.ticks >= 60 {
		ego.Log.Logf(0, "FPS: %d | Delta: %f", ego.FPS, ego.DeltaTime)
		e.ticks = 0
	}
	for i := range e.objs {
		e.objs[i].Draw()
	}
}

func (e *exampleScene) Destroy() {
	for i := range e.objs {
		e.objs[i].Sprite.Destroy()
	}
}

func (e *exampleScene) Start() {
	var err error
	e.numObjects = 1000
	e.objs = make([]object, e.numObjects)
	for i := 0; i < e.numObjects; i++ {
		e.objs[i] = NewObject()

	}
	if err != nil {
		ego.Log.LogError(1, err)
		os.Exit(1)
	}
}
