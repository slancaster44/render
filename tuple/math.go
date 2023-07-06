package tuple

import (
	"math"
	"render/utility"
)

/* TODO: SIMD? */

func Equal(t1, t2 *Tuple3) bool {
	return utility.FltCmp(t1.X, t2.X) && utility.FltCmp(t1.Y, t2.Y) && utility.FltCmp(t1.Z, t2.Z) && t1.Type == t2.Type
}

func Add(t1, t2 *Tuple3) *Tuple3 {
	return &Tuple3{t1.X + t2.X, t1.Y + t2.Y, t1.Z + t2.Z, t1.Type + t2.Type}
}

func Subtract(t1, t2 *Tuple3) *Tuple3 {
	return &Tuple3{t1.X - t2.X, t1.Y - t2.Y, t1.Z - t2.Z, t1.Type - t2.Type}
}

func DotProduct(t1, t2 *Tuple3) float64 {
	return t1.X*t2.X + t1.Y*t2.Y + t1.Z*t2.Z
}

func CrossProduct(t1, t2 *Tuple3) *Tuple3 {
	return NewVec3(t1.Y*t2.Z-t1.Z*t2.Y, t1.Z*t2.X-t1.X*t2.Z, t1.X*t2.Y-t1.Y*t2.X)
}

func ScalarMultiply(t1 *Tuple3, factor float64) *Tuple3 {
	return &Tuple3{t1.X * factor, t1.Y * factor, t1.Z * factor, t1.Type}
}

func ScalarDivide(t1 *Tuple3, divisor float64) *Tuple3 {
	return &Tuple3{t1.X / divisor, t1.Y / divisor, t1.Z / divisor, t1.Type}
}

func Negate(t1 *Tuple3) *Tuple3 {
	return &Tuple3{-t1.X, -t1.Y, -t1.Z, t1.Type}
}

func Magnitude(t1 *Tuple3) float64 {
	return math.Sqrt(t1.X*t1.X + t1.Y*t1.Y + t1.Z*t1.Z)
}

func Normalize(t1 *Tuple3) *Tuple3 {
	mag := Magnitude(t1)
	return &Tuple3{t1.X / mag, t1.Y / mag, t1.X / mag, t1.Type}
}

func Distance(t1, t2 *Tuple3) float64 {
	return Magnitude(Subtract(t1, t2))
}

func Direction(from, to *Tuple3) *Tuple3 {
	return Normalize(Subtract(to, from))
}
