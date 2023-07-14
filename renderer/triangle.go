package renderer

import (
	"image/color"
	"render/tuple"
	"render/utility"
)

type Triangle struct {
	P1, P2, P3       *tuple.Tuple3
	p2d1, p2d2, p2d3 *tuple.Tuple3
	TransformationFn func(world_coord *tuple.Tuple3) (coord_2d *tuple.Tuple3)
	pixels           []*tuple.Tuple3
	color            color.RGBA
}

func NewIsometricTriangle(p1, p2, p3 *tuple.Tuple3, c color.RGBA) *Triangle {
	return &Triangle{
		P1: p1, P2: p2, P3: p3,
		TransformationFn: Isometric,
		color:            c,
	}
}

func (t *Triangle) Pixels() []*tuple.Tuple3 {
	return t.pixels
}

func (t *Triangle) Color() color.RGBA {
	return t.color
}

func (t *Triangle) Bake() {
	t.p2d1 = t.TransformationFn(t.P1)
	t.p2d2 = t.TransformationFn(t.P2)
	t.p2d3 = t.TransformationFn(t.P3)

	boundingBox := func() (float64, float64, float64, float64) {
		return utility.Min(t.p2d1.X, t.p2d2.X, t.p2d3.X), utility.Min(t.p2d1.Y, t.p2d2.Y, t.p2d3.Y),
			utility.Max(t.p2d1.X, t.p2d2.X, t.p2d3.X), utility.Max(t.p2d1.Y, t.p2d2.Y, t.p2d3.Y)
	}
	startX, startY, endX, endY := boundingBox()

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			determinate := ((t.p2d1.X - t.p2d2.X) * (y - t.p2d2.Y)) - ((t.p2d1.Y - t.p2d2.Y) * (x - t.p2d2.X))
			if determinate < 0.0 {
				continue
			}

			determinate = ((t.p2d2.X - t.p2d3.X) * (y - t.p2d3.Y)) - ((t.p2d2.Y - t.p2d3.Y) * (x - t.p2d3.X))
			if determinate < 0.0 {
				continue
			}

			determinate = ((t.p2d3.X - t.p2d1.X) * (y - t.p2d1.Y)) - ((t.p2d3.Y - t.p2d1.Y) * (x - t.p2d1.X))
			if determinate > 0.0 {
				t.pixels = append(t.pixels, tuple.NewPnt3(x, y, 0))
			}
		}
	}
}
