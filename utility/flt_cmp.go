package utility

import (
	"math"
)

var FLTCMP_EPSILON float64 = 0.001

func FltCmp(a, b float64) bool {
	return math.Abs(a-b) < FLTCMP_EPSILON
}
