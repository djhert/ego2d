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
	obj, _ := sprite.LoadPNG("./deepfryface.png")
	obj.Position.SetXY(rand.Int31n(ego.Width), rand.Int31n(ego.Height))
	obj.Position.Rotation.Set(float64(rand.Int31n(360)))
	return object{obj, rand.Int31n(10), rand.Int31n(10), rand.Int31n(6)}
}

func (o *object) Update() {
	if o.Sprite.Position.GetX() >= ego.Width || o.Sprite.Position.GetX() <= 0 {
		o.speedX *= -1
	}
	if o.Sprite.Position.GetY() >= ego.Height || o.Sprite.Position.GetY() <= 0 {
		o.speedY *= -1
	}
	o.Sprite.Position.AddX(o.speedX)
	o.Sprite.Position.AddY(o.speedY)
	o.Sprite.Position.Rotation.Add(float64(o.speedRot))
}

func (o *object) Draw() {
	o.Sprite.Draw()
}
