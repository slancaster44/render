package tuple

const (
	VEC3 float64 = 0.0
	PNT3 float64 = 1.0
)

type Tuple3 struct {
	X, Y, Z float64
	Type    float64
}

func NewVec3(X, Y, Z float64) *Tuple3 {
	return &Tuple3{X, Y, Z, VEC3}
}

func NewPnt3(X, Y, Z float64) *Tuple3 {
	return &Tuple3{X, Y, Z, PNT3}
}
