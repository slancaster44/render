package matrix

import (
	"render/tuple"
	"testing"
)

func TestCreateMatrix(t *testing.T) {

	i := 11.3

	m1 := NewMatrix(4, 4)
	m1.PutRow(0, []float64{i, i - 1, i - 2, i + 11})
	m1.PutRow(1, []float64{i, i - 1, i - 2, i + 11})
	m1.PutRow(2, []float64{i, i - 1, i - 2, i + 11})
	m1.PutRow(3, []float64{i, i - 1, i - 2, i + 11})

}

func TestMultiply44(t *testing.T) {
	m1 := &Matrix44{
		Contents: [4][4]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 8, 7, 6},
			{5, 4, 3, 2},
		},
	}

	m2 := &Matrix44{
		Contents: [4][4]float64{
			{-2, 1, 2, 3},
			{3, 2, 1, -1},
			{4, 3, 6, 5},
			{1, 2, 7, 8},
		},
	}

	r1 := &Matrix44{
		Contents: [4][4]float64{
			{20, 22, 50, 48},
			{44, 54, 114, 108},
			{40, 58, 110, 102},
			{16, 26, 46, 42},
		},
	}

	if !Equal(Multiply44(m1, m2), r1) {
		t.Errorf("Multiply failed %v %v %v", m1, m2, Multiply44(m1, m2))
	}
}

func TestSubmatrix44(t *testing.T) {
	m1 := &Matrix44{Contents: [4][4]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	}}

	r1 := &Matrix33{Contents: [3][3]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	}}

	if !Equal(r1, Submatrix44(m1, 2, 1)) {
		t.Errorf("Submatrix33 fail %v %v", m1, Submatrix44(m1, 2, 1))
	}

}

func TestDeterminant44(t *testing.T) {
	m1 := &Matrix44{Contents: [4][4]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	}}

	if Determinant44(m1) != -4071 {
		t.Errorf("Determinant44 failed %v %v", m1, Determinant44(m1))
	}
}

func TestInverse44(t *testing.T) {
	m1 := &Matrix44{Contents: [4][4]float64{
		{-5, 2, 6, 2},
		{1, -5, 1, 1},
		{7, 7, -6, 1},
		{4, 5, -6, 1},
	}}

	r1 := &Matrix44{Contents: [4][4]float64{
		{-7 / 223.0, 18 / 223.0, 62 / 223.0, -66 / 223.0},
		{21 / 446.0, -27 / 223.0, 37 / 446.0, -25.0 / 446},
		{12 / 223.0, 1 / 223.0, 53 / 223.0, -78.0 / 223},
		{95 / 446.0, 69 / 223.0, -45 / 446.0, 163.0 / 446},
	}}

	if !Equal(Inverse44(m1), r1) {
		t.Errorf("Inverse33 failed %v\n%v\n%v", m1, r1, Inverse44(m1))
	}
}

func TestSlowMultiply(t *testing.T) {
	m1 := &Matrix44{
		Contents: [4][4]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 8, 7, 6},
			{5, 4, 3, 2},
		},
	}

	m2 := &Matrix44{
		Contents: [4][4]float64{
			{-2, 1, 2, 3},
			{3, 2, 1, -1},
			{4, 3, 6, 5},
			{1, 2, 7, 8},
		},
	}

	r1 := &Matrix44{
		Contents: [4][4]float64{
			{20, 22, 50, 48},
			{44, 54, 114, 108},
			{40, 58, 110, 102},
			{16, 26, 46, 42},
		},
	}

	if !Equal(SlowMultiply(m1, m2), r1) {
		t.Errorf("Multiply failed %v %v %v", m1, m2, Multiply44(m1, m2))
	}
}

func TestDeterminant22(t *testing.T) {
	m1 := &Matrix22{
		Contents: [2][2]float64{
			{1, 5},
			{-3, 2},
		},
	}

	if Determinant22(m1) != 17 {
		t.Errorf("Determinant test failed")
	}
}

func TestMultiply33(t *testing.T) {
	m1 := &Matrix33{
		Contents: [3][3]float64{
			{1, 2, 3},
			{5, 6, 7},
			{9, 8, 7},
		},
	}

	m2 := &Matrix33{
		Contents: [3][3]float64{
			{-2, 1, 2},
			{3, 2, 1},
			{4, 3, 6},
		},
	}

	r1 := &Matrix33{
		Contents: [3][3]float64{
			{16, 14, 22},
			{36, 38, 58},
			{34, 46, 68},
		},
	}

	if !Equal(Multiply33(m1, m2), r1) {
		t.Errorf("Multiply failed %v %v %v", m1, m2, Multiply33(m1, m2))
	}
}

func TestSlowTranspose(t *testing.T) {
	m1 := Matrix33{}
	m1.Contents[0] = [3]float64{0, 9, 3}
	m1.Contents[1] = [3]float64{9, 8, 0}
	m1.Contents[2] = [3]float64{1, 8, 5}

	r1 := Matrix33{}
	r1.Contents[0] = [3]float64{0, 9, 1}
	r1.Contents[1] = [3]float64{9, 8, 8}
	r1.Contents[2] = [3]float64{3, 0, 5}

	if !Equal(SlowTranspose(&m1), &r1) {
		t.Errorf("SlowTransform fail %v %v", m1, SlowTranspose(&m1))
	}
}

func TestSubmatrix33(t *testing.T) {
	m1 := &Matrix33{Contents: [3][3]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	}}

	r1 := &Matrix22{Contents: [2][2]float64{
		{-3, 2},
		{0, 6},
	}}

	if !Equal(r1, Submatrix33(m1, 0, 2)) {
		t.Errorf("Submatrix33 fail %v %v", m1, Submatrix33(m1, 0, 2))
	}
}

func TestTranspose33(t *testing.T) {
	m1 := Matrix33{}
	m1.Contents[0] = [3]float64{0, 9, 3}
	m1.Contents[1] = [3]float64{9, 8, 0}
	m1.Contents[2] = [3]float64{1, 8, 5}

	r1 := Matrix33{}
	r1.Contents[0] = [3]float64{0, 9, 1}
	r1.Contents[1] = [3]float64{9, 8, 8}
	r1.Contents[2] = [3]float64{3, 0, 5}

	if !Equal(Transpose33(&m1), &r1) {
		t.Errorf("SlowTransform fail %v %v", m1, SlowTranspose(&m1))
	}
}

func TestMultiply33T(t *testing.T) {
	m1 := &Matrix33{
		Contents: [3][3]float64{
			{1, 2, 3},
			{2, 4, 4},
			{8, 6, 4},
		},
	}

	t1 := tuple.NewPnt3(1, 2, 3)
	r0 := tuple.NewPnt3(14, 22, 32)

	if !tuple.Equal(r0, Multiply33T(m1, t1)) {
		t.Errorf("Matrix tuple multiplication failed %v %v %v", m1, t1, Multiply33T(m1, t1))
	}
}

func TestDeterminant33(t *testing.T) {
	m1 := &Matrix33{Contents: [3][3]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}}

	if Determinant33(m1) != -196.0 {
		t.Errorf("Determinant33 failed %v %v", m1, Determinant33(m1))
	}
}

func TestMinor33(t *testing.T) {
	m1 := &Matrix33{
		Contents: [3][3]float64{
			{3, 5, 0},
			{2, -1, -7},
			{6, -1, 5},
		},
	}

	if Minor33(m1, 1, 0) != 25.0 {
		t.Errorf("Minor33 fail %v %v", m1, Minor33(m1, 1, 0))
	}
}

func TestCofactor33(t *testing.T) {
	m1 := &Matrix33{
		Contents: [3][3]float64{
			{3, 5, 0},
			{2, -1, -7},
			{6, -1, 5},
		},
	}

	if Cofactor33(m1, 2, 2) > 0 {
		t.Errorf("Minor33 fail %v %v", m1, Cofactor33(m1, 2, 1))
	}
}

func TestInverse33(t *testing.T) {
	m1 := &Matrix33{Contents: [3][3]float64{
		{-5, 2, 6},
		{1, -5, 1},
		{7, 7, -6},
	}}

	r1 := &Matrix33{Contents: [3][3]float64{
		{23.0 / 163.0, 54.0 / 163, 32.0 / 163.0},
		{13.0 / 163.0, -12.0 / 163, 11.0 / 163.0},
		{42.0 / 163.0, 49.0 / 163, 23.0 / 163.0},
	}}

	if !Equal(Inverse33(m1), r1) {
		t.Errorf("Inverse33 failed %v\n%v\n%v", m1, r1, Inverse33(m1))
	}
}

func BenchmarkMatrix(b *testing.B) {
	for i := float64(0); i < 8192; i++ {
		m1 := Matrix44{}
		m1.Contents[0] = [4]float64{i, i - 1, i - 2, i + 11}
		m1.Contents[1] = [4]float64{i, i - 1, i - 2, i + 11}
		m1.Contents[2] = [4]float64{i, i - 1, i - 2, i + 11}
		m1.Contents[3] = [4]float64{i, i - 1, i - 2, i + 11}

		m2 := m1.Copy()

		Equal(&m1, m2)
	}
}

func Benchmark0_Matrix(b *testing.B) {
	for i := float64(0); i < 8192; i++ {
		m1 := NewMatrix(4, 4)
		m1.PutRow(0, []float64{i, i - 1, i - 2, i + 11})
		m1.PutRow(1, []float64{i, i - 1, i - 2, i + 11})
		m1.PutRow(2, []float64{i, i - 1, i - 2, i + 11})
		m1.PutRow(3, []float64{i, i - 1, i - 2, i + 11})

		m2 := m1.Copy()

		Equal(m1, m2)
	}
}

func BenchmarkMultiply44(b *testing.B) {
	for i := 0; i < 16345; i++ {
		m1 := &Matrix44{
			Contents: [4][4]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
		}

		m2 := &Matrix44{
			Contents: [4][4]float64{
				{-2, 1, 2, 3},
				{3, 2, 1, -1},
				{4, 3, 6, 5},
				{1, 2, 7, 8},
			},
		}

		Multiply44(m1, m2)
	}
}

func BenchmarkMultiply33(b *testing.B) {
	for i := 0; i < 16345; i++ {
		m1 := &Matrix33{
			Contents: [3][3]float64{
				{1, 2, 3},
				{5, 6, 7},
				{9, 8, 7},
			},
		}

		m2 := &Matrix33{
			Contents: [3][3]float64{
				{1, 2, 3},
				{5, 6, 7},
				{9, 8, 7},
			},
		}

		Multiply33(m1, m2)
	}
}

func BenchmarkSlowMultiply(b *testing.B) {
	for i := 0; i < 16345; i++ {
		m1 := &Matrix44{
			Contents: [4][4]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
		}

		m2 := &Matrix44{
			Contents: [4][4]float64{
				{-2, 1, 2, 3},
				{3, 2, 1, -1},
				{4, 3, 6, 5},
				{1, 2, 7, 8},
			},
		}

		SlowMultiply(m1, m2)
	}
}

func BenchmarkSlowTranspose(b *testing.B) {
	for i := 0; i < 16000; i++ {
		m1 := Matrix33{}
		m1.Contents[0] = [3]float64{0, 9, 3}
		m1.Contents[1] = [3]float64{9, 8, 0}
		m1.Contents[2] = [3]float64{1, 8, 5}

		SlowTranspose(&m1)
	}
}

func BenchmarkTranspose33(b *testing.B) {
	for i := 0; i < 16000; i++ {
		m1 := Matrix33{}
		m1.Contents[0] = [3]float64{0, 9, 3}
		m1.Contents[1] = [3]float64{9, 8, 0}
		m1.Contents[2] = [3]float64{1, 8, 5}

		Transpose33(&m1)
	}
}

func BenchmarkDeterminant33(b *testing.B) {
	for i := 0; i < 16000; i++ {
		m1 := &Matrix33{Contents: [3][3]float64{
			{1, 2, 6},
			{-5, 8, -4},
			{2, 6, 4},
		}}

		Determinant33(m1)
	}
}

func BenchmarkDeterminant44(b *testing.B) {
	for i := 0; i < 16000; i++ {
		m1 := &Matrix44{Contents: [4][4]float64{
			{1, 2, 6, 1},
			{-5, 8, -4, 1},
			{2, 6, 4, 1},
			{1, 2, 3, 4},
		}}

		Determinant44(m1)
	}
}

func BenchmarkInverse33(b *testing.B) {
	for i := 0; i < 16000; i++ {
		m1 := &Matrix33{Contents: [3][3]float64{
			{-5, 2, 6},
			{1, -5, 1},
			{7, 7, -6},
		}}

		Inverse33(m1)
	}
}

func BenchmarkInverse44(b *testing.B) {
	for i := 0; i < 16000; i++ {
		m1 := &Matrix44{Contents: [4][4]float64{
			{-5, 2, 6, 2},
			{1, -5, 1, 3},
			{7, 7, -6, 4},
			{4, 5, 6, 1},
		}}

		Inverse44(m1)
	}
}

func BenchmarkMultiply33T(b *testing.B) {
	for i := 0; i < 32000; i++ {
		m1 := &Matrix33{Contents: [3][3]float64{
			{-5, 2, 6},
			{1, -5, 1},
			{7, 7, -6},
		}}

		Multiply33T(m1, tuple.NewPnt3(1, 2, 3))
	}
}
