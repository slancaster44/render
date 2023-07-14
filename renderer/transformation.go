package renderer

import "render/tuple"

func Isometric(point *tuple.Tuple3) *tuple.Tuple3 {
	return &tuple.Tuple3{(point.X - point.Z), point.Y + ((point.X + point.Z) / 2), 0, tuple.PNT3}
}
