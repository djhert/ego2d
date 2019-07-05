package math

type Transform struct {
	Rect
	Rotation AtomicFloat
	Center   Pos2
}

func NewTransform(x, y, w, h int32, rot float64, center Pos2) Transform {
	return Transform{NewRect(x, y, w, h), NewAtomicFloat(rot), center}
}

