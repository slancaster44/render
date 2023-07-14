package renderer

import (
	"image/color"
	"math"
	"render/tuple"
	"render/utility"
)

type Line struct {
	P1, P2           *tuple.Tuple3
	p2d1, p2d2       *tuple.Tuple3
	pixels           []*tuple.Tuple3
	color            color.RGBA
	TransformationFn func(world_coord *tuple.Tuple3) (coord_2d *tuple.Tuple3)
}

func NewIsometricLine(p1, p2 *tuple.Tuple3, c color.RGBA) *Line {
	return &Line{
		P1: p1, P2: p2,
		TransformationFn: Isometric,
		color:            c,
	}
}

func (l *Line) Bake() {
	l.p2d1 = l.TransformationFn(l.P1)
	l.p2d2 = l.TransformationFn(l.P2)

	boundingBox := func() (float64, float64, float64, float64) {
		return math.Min(l.p2d1.X, l.p2d2.X), math.Min(l.p2d1.Y, l.p2d2.Y),
			math.Max(l.p2d1.X, l.p2d2.X), math.Max(l.p2d1.Y, l.p2d2.Y)
	}

	edge := tuple.Subtract(l.p2d1, l.p2d2)

	startX, startY, endX, endY := boundingBox()
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			determinate := ((edge.X - edge.X) * (y - edge.Y)) - ((edge.Y - edge.Y) * (x - edge.X))

			if utility.FltCmp(determinate, 0.0) {
				l.pixels = append(l.pixels, tuple.NewPnt3(x, y, 0))
			}
		}
	}
}

func (l *Line) Pixels() []*tuple.Tuple3 {
	return l.pixels
}

func (l *Line) Color() color.RGBA {
	return l.color
}
