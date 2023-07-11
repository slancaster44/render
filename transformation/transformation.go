package transformation

import (
	"render/tuple"
	"render/utility"
)

type Transformation func(t1 *tuple.Tuple3) *tuple.Tuple3

func Compose(ts ...Transformation) Transformation {
	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		var out *tuple.Tuple3
		for _, tr := range ts {
			out = tr(t1)
		}

		return out
	}
}

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
	cosR := utility.Cos(radians)
	sinR := utility.Sin(radians)

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
	cosR := utility.Cos(-radians)
	sinR := utility.Sin(-radians)

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
	cosR := utility.Cos(radians)
	sinR := utility.Sin(radians)

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
	cosR := utility.Cos(-radians)
	sinR := utility.Sin(-radians)

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
	cosR := utility.Cos(radians)
	sinR := utility.Sin(radians)

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
	cosR := utility.Cos(-radians)
	sinR := utility.Sin(-radians)

	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    cosR*t1.X - sinR*t1.Y,
			Y:    sinR*t1.X + cosR*t1.Y,
			Z:    t1.Z,
			Type: t1.Type,
		}
	}
}

//For variable nm, the value indicates how much input t1.n effects the output.m
//For example, the greater the yx, the more the value of t1.y will efect the output.x
func Shear(yx, zx, xy, zy, xz, yz float64) Transformation {
	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    yx*t1.X + zx*t1.X,
			Y:    xy*t1.Y + zy*t1.Y,
			Z:    yz*t1.Z + xz*t1.Z,
			Type: t1.Type,
		}
	}
}

func InverseShear(yx, zx, xy, zy, xz, yz float64) Transformation {
	return func(t1 *tuple.Tuple3) *tuple.Tuple3 {
		return &tuple.Tuple3{
			X:    t1.X / (yx + zx),
			Y:    t1.Y / (xy + zy),
			Z:    t1.Z / (xz + yz),
			Type: t1.Type,
		}
	}
}
