package tuple

import (
	"image/color"
	"math"
)

const (
	VEC3 byte = 0
	PNT3 byte = 1
)

type Tuple3 struct {
	X, Y, Z float64
	Type    byte
}

func NewVec3(X, Y, Z float64) *Tuple3 {
	return &Tuple3{X, Y, Z, VEC3}
}

func NewPnt3(X, Y, Z float64) *Tuple3 {
	return &Tuple3{X, Y, Z, PNT3}
}

func ColorToVec(c color.RGBA) *Tuple3 {
	return NewVec3(float64(c.R)/255, float64(c.G)/255, float64(c.B)/255)
}

func VecToColor(t *Tuple3) color.RGBA {
	return color.RGBA{
		R: uint8(math.Min(t.X*255, 255)),
		G: uint8(math.Min(t.Y*255, 255)),
		B: uint8(math.Min(t.Z*255, 255)),
		A: 255,
	}
}
