package renderer

import (
	"image/color"
	"render/tuple"
)

type Dot struct {
	P1               *tuple.Tuple3
	p2d1             *tuple.Tuple3
	color            color.RGBA
	TransformationFn func(world_coord *tuple.Tuple3) (coord_2d *tuple.Tuple3)
}

func NewIsometricDot(p1 *tuple.Tuple3, c color.RGBA) *Dot {
	return &Dot{
		P1:               p1,
		color:            c,
		TransformationFn: Isometric,
	}
}

func (d *Dot) Bake() {
	d.p2d1 = d.TransformationFn(d.P1)

}

func (d *Dot) Pixels() []*tuple.Tuple3 {
	return []*tuple.Tuple3{d.p2d1}
}

func (d *Dot) Color() color.RGBA {
	return d.color
}
