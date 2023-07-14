package renderer

import (
	"image/color"
	"render/tuple"
)

type Shape interface {
	Bake()
	Pixels() []*tuple.Tuple3
	Color() color.RGBA
}
