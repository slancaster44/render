package ray_tracer

import (
	"image/color"
	"render/tuple"
)

type Light struct {
	Location  *tuple.Tuple3
	Color     color.RGBA
	Intensity float64
}

func NewLight(l *tuple.Tuple3, c color.RGBA) *Light {
	return &Light{l, c, 1.0}
}
