package main

import (
	"math/rand"

	"github.com/hlfstr/ego2d/component/sprite"
	"github.com/hlfstr/ego2d/ego"
)

type object struct {
	Sprite *sprite.Sprite

	speedX   int32
	speedY   int32
	speedRot int32
}

func NewObject() object {
	obj := sprite.New("player/deepfryface")
	obj.SetXY(rand.Int31n(ego.Width), rand.Int31n(ego.Height))
	obj.Rotation.Set(float64(rand.Int31n(360)))

	return object{obj, rand.Int31n(9) + 1, rand.Int31n(9) + 1, rand.Int31n(5) + 1}
}

func (o *object) Update() {
	if o.Sprite.GetX() >= (ego.Width - o.Sprite.GetW()) {
		o.speedX *= -1
		o.Sprite.SetX(ego.Width - o.Sprite.GetW())

	} else if o.Sprite.GetX() <= 0 {
		o.speedX *= -1
		o.Sprite.SetX(0)
	}
	if o.Sprite.GetY() >= (ego.Height - o.Sprite.GetH()) {
		o.speedY *= -1
		o.Sprite.SetY(ego.Height - o.Sprite.GetH())
	} else if o.Sprite.GetY() <= 0 {
		o.speedY *= -1
		o.Sprite.SetY(0)
	}

	o.Sprite.AddX(o.speedX)
	o.Sprite.AddY(o.speedY)
	o.Sprite.Rotation.Add(float64(o.speedRot))
}

func (o *object) Draw() {
	o.Sprite.Draw()
}
