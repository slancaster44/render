package transformation

import (
	"math"
	"render/tuple"
)

type Transformation func(t1 *tuple.Tuple3) *tuple.Tuple3

func Translation(x, y, z float64) Transformation {
	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return tuple.Add(t1, tuple.NewVec3(x, y, z))
	}
}

func InverseTranslation(x, y, z float64) Transformation {
	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return tuple.Subtract(t1, tuple.NewVec3(x, y, z))
	}
}

func Scaling(x, y, z float64) Transformation {
	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    t1.X * x,
			Y:    t1.Y * y,
			Z:    t1.Z * z,
			Type: t1.Type,
		}
	}
}

func InverseScaling(x, y, z float64) Transformation {
	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    t1.X / x,
			Y:    t1.Y / y,
			Z:    t1.Z / z,
			Type: t1.Type,
		}
	}
}

func RotationX(radians float64) Transformation {
	cosR := math.Cos(radians)
	sinR := math.Sin(radians)

	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    t1.X,
			Y:    cosR*t1.Y - sinR*t1.Z,
			Z:    sinR*t1.Y + cosR*t1.Z,
			Type: t1.Type,
		}
	}
}

func InverseRotationX(radians float64) Transformation {
	cosR := math.Cos(-radians)
	sinR := math.Sin(-radians)

	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    t1.X,
			Y:    cosR*t1.Y - sinR*t1.Z,
			Z:    sinR*t1.Y + cosR*t1.Z,
			Type: t1.Type,
		}
	}
}

func RotationY(radians float64) Transformation {
	cosR := math.Cos(radians)
	sinR := math.Sin(radians)

	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    cosR*t1.X + sinR*t1.Z,
			Y:    t1.Y,
			Z:    cosR*t1.Z - sinR*t1.X,
			Type: t1.Type,
		}
	}
}

func InverseRotationY(radians float64) Transformation {
	cosR := math.Cos(-radians)
	sinR := math.Sin(-radians)

	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    cosR*t1.X + sinR*t1.Z,
			Y:    t1.Y,
			Z:    cosR*t1.Z - sinR*t1.X,
			Type: t1.Type,
		}
	}
}

func RotationZ(radians float64) Transformation {
	cosR := math.Cos(radians)
	sinR := math.Sin(radians)

	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    cosR*t1.X - sinR*t1.Y,
			Y:    sinR*t1.X + cosR*t1.Y,
			Z:    t1.Z,
			Type: t1.Type,
		}
	}
}

func InverseRotationZ(radians float64) Transformation {
	cosR := math.Cos(-radians)
	sinR := math.Sin(-radians)

	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    cosR*t1.X - sinR*t1.Y,
			Y:    sinR*t1.X + cosR*t1.Y,
			Z:    t1.Z,
			Type: t1.Type,
		}
	}
}
