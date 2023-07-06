package tuple

//#cgo amd64 CFLAGS: -O3 -mavx
//#include "simd_x86.h"
import "C"

func Add_x86_Simd(t1, t2 [4]float64) *Tuple3 {
	res := C.Add((*C.double)(&t1[0]), (*C.double)(&t2[0]))

	return &Tuple3{
		float64(res[0]),
		float64(res[1]),
		float64(res[2]),
		float64(res[3]),
	}
}
