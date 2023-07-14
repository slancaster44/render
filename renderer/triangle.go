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

	/*Determines if two points fall on the oppisate sides
	 * of the triangles bounding box
	 */
	areOpposites := func(p1, p2 *tuple.Tuple3) bool {
		return (p1.X == startX && p2.X == endX) ||
			(p2.X == startX && p1.X == endX) ||
			(p1.Y == startY && p2.Y == endY) ||
			(p2.Y == startY && p1.X == endY)
	}

	//edge0 := [2]*tuple.Tuple3{t.p2d2, t.p2d1}
	//edge1 := [2]*tuple.Tuple3{t.p2d3, t.p2d2}
	//edge2 := [2]*tuple.Tuple3{t.p2d1, t.p2d3}

	var edge0 [2]*tuple.Tuple3
	var edge1 [2]*tuple.Tuple3
	var edge2 [2]*tuple.Tuple3

	if areOpposites(t.p2d2, t.p2d1) {
		edge0 = [2]*tuple.Tuple3{t.p2d2, t.p2d1}
		edge1 = [2]*tuple.Tuple3{t.p2d3, t.p2d2}
		edge2 = [2]*tuple.Tuple3{t.p2d1, t.p2d3}
	} else if areOpposites(t.p2d1, t.p2d3) {
		edge0 = [2]*tuple.Tuple3{t.p2d1, t.p2d3}
		edge1 = [2]*tuple.Tuple3{t.p2d3, t.p2d2}
		edge2 = [2]*tuple.Tuple3{t.p2d2, t.p2d1}
	} else {
		edge0 = [2]*tuple.Tuple3{t.p2d3, t.p2d2}
		edge1 = [2]*tuple.Tuple3{t.p2d2, t.p2d1}
		edge2 = [2]*tuple.Tuple3{t.p2d1, t.p2d3}
	}

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			determinate := ((edge0[1].X - edge0[0].X) * (y - edge0[0].Y)) - ((edge0[1].Y - edge0[0].Y) * (x - edge0[0].X))
			if determinate < 0.0 {
				continue
			}

			determinate = ((edge1[1].X - edge1[0].X) * (y - edge1[0].Y)) - ((edge1[1].Y - edge1[0].Y) * (x - edge1[0].X))
			if determinate < 0.0 {
				continue
			}

			determinate = ((edge2[1].X - edge2[0].X) * (y - edge2[0].Y)) - ((edge2[1].Y - edge2[0].Y) * (x - edge2[0].X))
			if determinate > 0.0 {
				t.pixels = append(t.pixels, tuple.NewPnt3(x, y, 0))
			}
		}
	}
}
