package ray_tracer

import "render/tuple"

type Ray struct {
	Origin    *tuple.Tuple3
	Direction *tuple.Tuple3
}

func NewRay(ox, oy, oz, dirx, diry, dirz float64) *Ray {
	return &Ray{
		Origin:    tuple.NewPnt3(ox, oy, oz),
		Direction: tuple.NewVec3(dirx, diry, dirz),
	}
}

func (r *Ray) PositionAt(t float64) *tuple.Tuple3 {
	return tuple.Add(r.Origin, tuple.ScalarMultiply(r.Direction, t))
}
