package main

import (
	"math/rand"

	"github.com/hlfstr/ego2d/component/sprite"
	"github.com/hlfstr/ego2d/ego"
)

type object struct {
	*sprite.Sprite

	speedX   int32
	speedY   int32
	speedRot int32

	index int
}

func NewObject() *object {
	obj := &object{
		Sprite: sprite.New(),
	}
	return obj
}

func (o *object) Update() {
	if o.Sprite.GetX() >= (ego.Width - o.Sprite.GetW()) {
		o.speedX *= -1
		o.Sprite.SetX((ego.Width - o.Sprite.GetW()) - 5)
		//	ego.Audio.PaySound("bump")
	} else if o.Sprite.GetX() <= 0 {
		o.speedX *= -1
		o.Sprite.SetX(5)
		//	ego.Audio.PaySound("bump")
	}

	if o.Sprite.GetY() >= (ego.Height - o.Sprite.GetH()) {
		o.speedY *= -1
		o.Sprite.SetY((ego.Height - o.Sprite.GetH()) - 5)
		//	ego.Audio.PaySound("bump")

	} else if o.Sprite.GetY() <= 0 {
		o.speedY *= -1
		o.Sprite.SetY(5)
		//	ego.Audio.PaySound("bump")
	}

	o.Sprite.AddX(o.speedX)
	o.Sprite.AddY(o.speedY)
	o.Sprite.Rotation.Add(float64(o.speedRot))
}

func (o *object) Start() {
	o.SetXY(rand.Int31n(ego.Width), rand.Int31n(ego.Height))
	o.Rotation.Set(float64(rand.Int31n(360)))
	o.speedX = rand.Int31n(9) + 1
	o.speedY = rand.Int31n(9) + 1
	o.speedRot = rand.Int31n(5) + 1
}
