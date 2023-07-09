package transformation

import (
	"math"
	"render/tuple"
	"testing"
)

func TestTranslate(t *testing.T) {
	translate := Translation(5, -3, 2)
	inverse := InverseTranslation(5, -3, 2)
	p1 := tuple.NewPnt3(-3, 4, 5)

	r1 := tuple.NewPnt3(2, 1, 7)
	if !tuple.Equal(r1, translate(p1)) || !tuple.Equal(inverse(translate(p1)), p1) {
		t.Errorf("Translation failed")
	}
}

func TestScaling(t *testing.T) {
	scale := Scaling(2, 3, 4)
	inverse := InverseScaling(2, 3, 4)
	p1 := tuple.NewPnt3(-4, 6, 8)

	r1 := tuple.NewPnt3(-8, 18, 32)
	if !tuple.Equal(r1, scale(p1)) || !tuple.Equal(inverse(scale(p1)), p1) {
		t.Errorf("Scaling failed")
	}
}

func TestRotationX(t *testing.T) {
	p1 := tuple.NewPnt3(0, 1, 0)
	rotate2 := RotationX(math.Pi / 2)
	inverse := InverseRotationX(math.Pi / 2)

	r0 := tuple.NewPnt3(0, 0, 1)
	if !tuple.Equal(r0, rotate2(p1)) || !tuple.Equal(inverse(rotate2(p1)), p1) {
		t.Errorf("Rotate X failed")
	}
}

func TestRotationZ(t *testing.T) {
	p1 := tuple.NewPnt3(0, 1, 0)
	rotate2 := RotationZ(math.Pi / 2)
	inverse := InverseRotationZ(math.Pi / 2)

	r0 := tuple.NewPnt3(-1, 0, 0)
	if !tuple.Equal(r0, rotate2(p1)) || !tuple.Equal(inverse(rotate2(p1)), p1) {
		t.Errorf("Rotate Y failed %v", rotate2(p1))
	}
}

func TestRotationY(t *testing.T) {
	p1 := tuple.NewPnt3(0, 0, 1)
	rotate2 := RotationY(math.Pi / 2)
	inverse := InverseRotationY(math.Pi / 2)

	r0 := tuple.NewPnt3(1, 0, 0)
	if !tuple.Equal(r0, rotate2(p1)) || !tuple.Equal(inverse(rotate2(p1)), p1) {
		t.Errorf("Rotate Y failed %v", inverse(rotate2(p1)))
	}
}

func BenchmarkScaling(b *testing.B) {
	scale := Scaling(2, 3, 4)

	for i := 0; i < 32000; i++ {
		p1 := tuple.NewPnt3(-4, 6, 8)

		scale(p1)
	}
}

func BenchmarkRotateX(b *testing.B) {
	rotate := RotationX(math.Pi / 4)

	for i := 0; i < 32000; i++ {
		p1 := tuple.NewPnt3(-4, 6, 8)

		rotate(p1)
	}
}
