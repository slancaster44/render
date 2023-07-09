package tuple

const (
	VEC3 byte = 0
	PNT3 byte = 1
)

type Tuple3 struct {
	X, Y, Z float64
	Type    byte
}

func NewVec3(X, Y, Z float64) *Tuple3 {
	return &Tuple3{X, Y, Z, VEC3}
}

func NewPnt3(X, Y, Z float64) *Tuple3 {
	return &Tuple3{X, Y, Z, PNT3}
}
