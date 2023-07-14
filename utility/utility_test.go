package utility

import (
	"math"
	"testing"
)

func TestTrig(t *testing.T) {
	if !FltCmp(math.Sin(math.Pi/4), Sin(math.Pi/4)) {
		t.Errorf("Sin failed %v %v", Sin(math.Pi/4), math.Sin(math.Pi/4))
	}

	if !FltCmp(math.Sin(5*math.Pi/4), Sin(5*math.Pi/4)) {
		t.Errorf("Sin failed %v %v", Sin(5*math.Pi/4), math.Sin(5*math.Pi/4))
	}

	if !FltCmp(math.Sin(3*math.Pi/4), Sin(3*math.Pi/4)) {
		t.Errorf("Sin failed %v %v", Sin(3*math.Pi/4), math.Sin(3*math.Pi/4))
	}

	if !FltCmp(math.Sin(7*math.Pi/4), Sin(7*math.Pi/4)) {
		t.Errorf("Sin failed %v %v", Sin(7*math.Pi/4), math.Sin(7*math.Pi/4))
	}

	if Sin(-(3 * math.Pi / 4)) != Sin(5*math.Pi/4) {
		t.Errorf("Sin failed %v", Sin(-(3*math.Pi/4)) == Sin(5*math.Pi/4))
	}

	if !FltCmp(math.Tan(math.Pi/4), Tan(math.Pi/4)) {
		t.Errorf("Tan failed %v", Tan(math.Pi/4))
	}

	if !FltCmp(math.Tan(3*math.Pi/4), Tan(3*math.Pi/4)) {
		t.Errorf("Tan failed")
	}
}

func TestDoJobs(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	job_fn := func(x int) int {
		return x + 1
	}

	output := DoJobs(list, 4, job_fn)
	if len(output) != len(list) || !Contains(output, 10) || Contains(output, 1) {
		t.Errorf("Do jobs test failed %v", output)
	}
}

func BenchmarkMathSin(b *testing.B) {
	for i := 0; i < 16000; i++ {
		math.Sin(math.Pi / 4)
	}
}

func BenchmarkSin(b *testing.B) {
	for i := 0; i < 16000; i++ {
		Sin(math.Pi / 4)
	}
}
