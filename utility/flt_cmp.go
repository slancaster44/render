package utility

var FLTCMP_EPSILON float64 = 0.001

func FltCmp(a, b float64) bool {
	i := a - b
	return i < FLTCMP_EPSILON && i > -FLTCMP_EPSILON
}
