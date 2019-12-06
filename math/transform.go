package math

type Transform struct {
	Rect
	Rotation AtomicFloat
	Center   Point
}

func NewTransform(x, y, w, h int32, rot float64, center Point) Transform {
	return Transform{
		Rect:     NewRect(x, y, w, h),
		Rotation: NewAtomicFloat(rot),
		Center:   center,
	}
}
