package utility

import (
	"math"
)

//TODO: Increase accuracy by only storing values from sin(0) to sin(pi/2)

const TABLE_SIZE int = 8192
const TABLE_SIZE_FLT float64 = 8192.0

const STEP_SIZE float64 = math.Pi * 2.0 / TABLE_SIZE_FLT
const TWO_PI float64 = math.Pi * 2
const PI_OVER_TWO float64 = math.Pi / 2

var sinTable [TABLE_SIZE]float32

func genLookupTables() {
	for i := 0; i < TABLE_SIZE; i++ {
		sinTable[i] = float32(math.Sin(float64(i) * STEP_SIZE))
	}
}

func Sin(radians float64) float64 {
	i0 := int(radians/STEP_SIZE) % TABLE_SIZE
	if i0 < 0 {
		i0 += TABLE_SIZE
	}
	i1 := i0 + 1

	return float64((sinTable[i0] + sinTable[i1]) / 2)
}

func Cos(radians float64) float64 {
	radians -= PI_OVER_TWO

	i0 := int(radians/STEP_SIZE) % TABLE_SIZE
	if i0 < 0 {
		i0 += TABLE_SIZE
	}

	i1 := i0 + 1

	return float64((sinTable[i0] + sinTable[i1]) / 2)
}

func Tan(radians float64) float64 {
	return -(Sin(radians) / Cos(radians))
}

func Min(x, y, z float64) float64 {
	if x < y && x < z {
		return x
	} else if y < z {
		return y
	}

	return z
}

func Max(x, y, z float64) float64 {
	if x > y && x > z {
		return x
	} else if y > z {
		return y
	}

	return z
}
