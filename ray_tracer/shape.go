package ray_tracer

import (
	"image/color"
	"render/tuple"
	"render/utility"
)

type Triangle struct {
	P1, P2, P3 *tuple.Tuple3
	Material   *Material
}

func NewTriangle(p1, p2, p3 *tuple.Tuple3) *Triangle {
	return &Triangle{
		p1, p2, p3, &Material{color.RGBA{128, 255, 255, 255}, 0.1, 0.9, 0.9, 200},
	}
}

//Stolen: https://www.scratchapixel.com/lessons/3d-basic-rendering/ray-tracing-rendering-a-triangle/barycentric-coordinates.html
//A point on a triangle will divide that triangle into 3 sub-triangles
//If the area of all those sub-triangles is equal to the area of the full triangle,
//then the point lays on the triangle, otherwise it does not lay on the full triangle
func (t *Triangle) Intersect(r *Ray) (*Intersection, *tuple.Tuple3) {

	/* Test if triangle and ray are parallel */
	n := t.Normal()
	n_dot_raydir := tuple.DotProduct(n, r.Direction)
	if utility.FltCmp(n_dot_raydir, 0.0) {
		return nil, n
	}

	/* Calculate plane-ray intersection */
	ray_t := tuple.DotProduct(tuple.Subtract(t.P1, r.Origin), n) / tuple.DotProduct(n, r.Direction)
	intersecting_point := r.PositionAt(ray_t)

	/* Calculate Areas */

	//The normal of a triangle is proportonal to its area
	area := func(p1, p2, p3 *tuple.Tuple3) float64 {
		BA := tuple.Subtract(p1, p2)
		CA := tuple.Subtract(p3, p2)
		return tuple.Magnitude(tuple.CrossProduct(BA, CA)) / 2 //triangle normal / 2.0
	}

	bigTriangleArea := tuple.Magnitude(n) / 2
	uArea := area(t.P1, t.P2, intersecting_point)
	vArea := area(t.P1, t.P3, intersecting_point)
	wArea := area(t.P2, t.P3, intersecting_point)

	if utility.FltCmp(bigTriangleArea, uArea+vArea+wArea) {
		return &Intersection{t, ray_t}, n
	}

	return nil, n
}

//Stolen: https://web.ma.utexas.edu/users/m408m/Display12-5-4.shtml
func (t *Triangle) Normal() *tuple.Tuple3 {
	QR := tuple.Subtract(t.P1, t.P2)
	QS := tuple.Subtract(t.P3, t.P2)
	return tuple.CrossProduct(QR, QS)
}
