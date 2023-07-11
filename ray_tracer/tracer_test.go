package ray_tracer

import (
	"render/tuple"
	"testing"
)

func TestRayPosition(t *testing.T) {
	ray := NewRay(2, 3, 4, 1, 0, 0)

	r0 := tuple.NewPnt3(2, 3, 4)
	r1 := tuple.NewPnt3(3, 3, 4)
	r2 := tuple.NewPnt3(1, 3, 4)
	r3 := tuple.NewPnt3(4.5, 3, 4)

	if !tuple.Equal(r0, ray.PositionAt(0)) ||
		!tuple.Equal(r1, ray.PositionAt(1)) ||
		!tuple.Equal(r2, ray.PositionAt(-1)) ||
		!tuple.Equal(r3, ray.PositionAt(2.5)) {

		t.Errorf("Ray position finding failed")
	}
}

func TestGetHit(t *testing.T) {
	l := []*Intersection{
		NewIntersection(nil, -1),
		NewIntersection(nil, 0),
		NewIntersection(nil, 100),
	}

	if GetHit(l).TValue != 0 {
		t.Errorf("Expected 0 tvalue. Get Hit failed")
	}

	l2 := []*Intersection{
		NewIntersection(nil, -1),
		NewIntersection(nil, 100),
		NewIntersection(nil, 1),
	}

	if GetHit(l2).TValue != 1 {
		t.Errorf("Expected 0 tvalue. Get Hit failed")
	}
}

func TestCamera(t *testing.T) {
	c := NewCamera(201, 101)
	r := c.RayForPixel(100, 50)

	expected_origin := tuple.NewPnt3(0, 0, 1.5)
	expected_direction := tuple.NewVec3(0, 0, -1)

	if !tuple.Equal(r.Origin, expected_origin) || !tuple.Equal(r.Direction, expected_direction) {
		t.Errorf("Camera failed %v %v", r.Origin, r.Direction)
	}

	r2 := c.RayForPixel(0, 0)
	ed2 := tuple.NewVec3(0.66519, 0.33259, -0.66851)
	if !tuple.Equal(r2.Direction, ed2) {
		t.Errorf("camera failed %v", ed2)
	}
}

func TestRayTraceLocations(t *testing.T) {
	r := NewRayTracer(100, 100)
	z1 := r.screenCoordToIndex(25, 25)
	x, y := r.screenIndexToCoord(z1)

	if x != 25 || y != 25 {
		t.Errorf("Locations failed")
	}
}
