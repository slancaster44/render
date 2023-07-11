package tuple_test

import (
	"math"
	"render/tuple"
	"render/utility"
	"testing"
)

func TestMath(t *testing.T) {
	assert := func(t1, t2 *tuple.Tuple3) {
		if !tuple.Equal(t1, t2) {
			t.Errorf("Tuple math failure %v %v", t1, t2)
		}
	}

	assert_flt := func(f1, f2 float64) {
		if !utility.FltCmp(f1, f2) {
			t.Errorf("Float math failure %v %v", f1, f2)
		}
	}

	p1 := tuple.NewPnt3(3, -2, 5)
	v1 := tuple.NewVec3(-2, 3, 1)
	r1 := tuple.NewPnt3(1, 1, 6)
	r2 := tuple.NewPnt3(5, -5, 4)
	r3 := tuple.NewVec3(2, -3, -1)

	assert(tuple.Add(p1, v1), r1)
	assert(tuple.Subtract(p1, v1), r2)
	assert(tuple.Negate(v1), r3)

	v2 := tuple.NewVec3(2, 2, 2)
	r4 := tuple.NewVec3(4, 4, 4)
	assert(tuple.ScalarMultiply(v2, 2), r4)
	assert(tuple.ScalarDivide(r4, 2), v2)

	v3 := tuple.NewVec3(1, 2, 3)
	assert_flt(tuple.Magnitude(v3), float64(math.Sqrt(14)))

	assert_flt(tuple.Magnitude(tuple.Normalize(v2)), 1.0)

	assert_flt(tuple.Distance(v1, v2), 4.242641)

	/*WARNING: Better test needed for direction*/
	r5 := tuple.NewVec3(0.9428090415820635, -0.23570226039551587, 0.9428090415820635)
	assert(tuple.Direction(v1, v2), r5)

	assert_flt(tuple.DotProduct(v1, v2), 4)

	r6 := tuple.NewVec3(4, 6, -10)
	assert(r6, tuple.CrossProduct(v1, v2))
}

func TestReflection(t *testing.T) {
	v := tuple.NewVec3(1, -1, 0)
	n := tuple.NewVec3(0, 1, 0)

	if !tuple.Equal(tuple.Reflect(v, n), tuple.NewVec3(1, 1, 0)) {
		t.Errorf("Reflection failed")
	}

	v = tuple.NewVec3(0, -1, 0)
	n = tuple.NewVec3(0.707106, 0.707106, 0)

	if !tuple.Equal(tuple.Reflect(v, n), tuple.NewVec3(1, 0, 0)) {
		t.Errorf("Reflection failed %v", tuple.Reflect(v, n))
	}
}
