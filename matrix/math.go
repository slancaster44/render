package matrix

import (
	"render/tuple"
)

func Equal(m1, m2 Matrix) bool {
	if (m1.RowCount() != m2.RowCount()) || (m1.ColumnCount() != m2.ColumnCount()) {
		return false
	}

	for i := 0; i < m1.RowCount(); i++ {
		for j := 0; j < m2.RowCount(); j++ {
			if m1.Get(i, j) != m2.Get(i, j) {
				return false
			}
		}
	}

	return true
}

func Determinant22(m1 *Matrix22) float64 {
	return m1.Contents[0][0]*m1.Contents[1][1] - m1.Contents[0][1]*m1.Contents[1][0]
}

func SlowMultiply(m1, m2 Matrix) Matrix {
	retVal := NewMatrix(m2.ColumnCount(), m1.RowCount())

	for row := 0; row < retVal.RowCount(); row++ {
		for column := 0; column < retVal.ColumnCount(); column++ {
			res := 0.0
			for j := 0; j < retVal.RowCount(); j++ {
				res += m1.Get(row, j) * m2.Get(j, column)
			}
			retVal.Put(row, column, res)
		}
	}

	return retVal
}

func SlowTranspose(m1 Matrix) Matrix {
	var out Matrix

	switch m1_as_primative := m1.(type) {
	case *Matrix33:
		out = Transpose33(m1_as_primative)
	case *Matrix44:
		out = SlowTranspose(m1_as_primative)
	default:
		out = NewMatrix(m1.RowCount(), m1.ColumnCount())

		for column := 0; column < m1.ColumnCount(); column++ {
			out.PutRow(column, m1.GetColumn(column))
		}
	}

	return out
}

func SlowDeterminant(m1 Matrix) float64 {
	switch m_typed := m1.(type) {
	case (*Matrix22):
		return Determinant22(m_typed)
	case (*Matrix33):
		return Determinant33(m_typed)
	case (*Matrix44):
		return Determinant44(m_typed)
	default:
		panic("Determinant not supported for this type of matrix")
	}
}

func SlowInverse(m1 Matrix) Matrix {
	switch m_typed := m1.(type) {
	case (*Matrix33):
		return Inverse33(m_typed)
	case (*Matrix44):
		return Inverse44(m_typed)
	default:
		panic("Determinant not supported for this type of matrix")
	}
}

func Multiply44(m1, m2 *Matrix44) *Matrix44 {
	return &Matrix44{
		Contents: [4][4]float64{
			{
				m1.Contents[0][0]*m2.Contents[0][0] + m1.Contents[0][1]*m2.Contents[1][0] + m1.Contents[0][2]*m2.Contents[2][0] + m1.Contents[0][3]*m2.Contents[3][0],
				m1.Contents[0][0]*m2.Contents[0][1] + m1.Contents[0][1]*m2.Contents[1][1] + m1.Contents[0][2]*m2.Contents[2][1] + m1.Contents[0][3]*m2.Contents[3][1],
				m1.Contents[0][0]*m2.Contents[0][2] + m1.Contents[0][1]*m2.Contents[1][2] + m1.Contents[0][2]*m2.Contents[2][2] + m1.Contents[0][3]*m2.Contents[3][2],
				m1.Contents[0][0]*m2.Contents[0][3] + m1.Contents[0][1]*m2.Contents[1][3] + m1.Contents[0][2]*m2.Contents[2][3] + m1.Contents[0][3]*m2.Contents[3][3],
			},
			{
				m1.Contents[1][0]*m2.Contents[0][0] + m1.Contents[1][1]*m2.Contents[1][0] + m1.Contents[1][2]*m2.Contents[2][0] + m1.Contents[1][3]*m2.Contents[3][0],
				m1.Contents[1][0]*m2.Contents[0][1] + m1.Contents[1][1]*m2.Contents[1][1] + m1.Contents[1][2]*m2.Contents[2][1] + m1.Contents[1][3]*m2.Contents[3][1],
				m1.Contents[1][0]*m2.Contents[0][2] + m1.Contents[1][1]*m2.Contents[1][2] + m1.Contents[1][2]*m2.Contents[2][2] + m1.Contents[1][3]*m2.Contents[3][2],
				m1.Contents[1][0]*m2.Contents[0][3] + m1.Contents[1][1]*m2.Contents[1][3] + m1.Contents[1][2]*m2.Contents[2][3] + m1.Contents[1][3]*m2.Contents[3][3],
			},
			{
				m1.Contents[2][0]*m2.Contents[0][0] + m1.Contents[2][1]*m2.Contents[1][0] + m1.Contents[2][2]*m2.Contents[2][0] + m1.Contents[2][3]*m2.Contents[3][0],
				m1.Contents[2][0]*m2.Contents[0][1] + m1.Contents[2][1]*m2.Contents[1][1] + m1.Contents[2][2]*m2.Contents[2][1] + m1.Contents[2][3]*m2.Contents[3][1],
				m1.Contents[2][0]*m2.Contents[0][2] + m1.Contents[2][1]*m2.Contents[1][2] + m1.Contents[2][2]*m2.Contents[2][2] + m1.Contents[2][3]*m2.Contents[3][2],
				m1.Contents[2][0]*m2.Contents[0][3] + m1.Contents[2][1]*m2.Contents[1][3] + m1.Contents[2][2]*m2.Contents[2][3] + m1.Contents[2][3]*m2.Contents[3][3],
			},
			{
				m1.Contents[3][0]*m2.Contents[0][0] + m1.Contents[3][1]*m2.Contents[1][0] + m1.Contents[3][2]*m2.Contents[2][0] + m1.Contents[3][3]*m2.Contents[3][0],
				m1.Contents[3][0]*m2.Contents[0][1] + m1.Contents[3][1]*m2.Contents[1][1] + m1.Contents[3][2]*m2.Contents[2][1] + m1.Contents[3][3]*m2.Contents[3][1],
				m1.Contents[3][0]*m2.Contents[0][2] + m1.Contents[3][1]*m2.Contents[1][2] + m1.Contents[3][2]*m2.Contents[2][2] + m1.Contents[3][3]*m2.Contents[3][2],
				m1.Contents[3][0]*m2.Contents[0][3] + m1.Contents[3][1]*m2.Contents[1][3] + m1.Contents[3][2]*m2.Contents[2][3] + m1.Contents[3][3]*m2.Contents[3][3],
			},
		},
	}
}

func Multiply44T(m1 *Matrix44, t1 *tuple.Tuple3) *tuple.Tuple3 {
	return &tuple.Tuple3{
		X:    m1.Contents[0][0]*t1.X + m1.Contents[0][1]*t1.Y + m1.Contents[0][2]*t1.Z + m1.Contents[0][3]*float64(t1.Type),
		Y:    m1.Contents[1][0]*t1.X + m1.Contents[1][1]*t1.Y + m1.Contents[1][2]*t1.Z + m1.Contents[1][3]*float64(t1.Type),
		Z:    m1.Contents[2][0]*t1.X + m1.Contents[2][1]*t1.Y + m1.Contents[2][2]*t1.Z + m1.Contents[2][3]*float64(t1.Type),
		Type: t1.Type,
	}
}

func Transpose44(m1 *Matrix44) *Matrix44 {
	return &Matrix44{
		Contents: [4][4]float64{
			{m1.Contents[0][0], m1.Contents[1][0], m1.Contents[2][0], m1.Contents[3][0]},
			{m1.Contents[0][1], m1.Contents[1][1], m1.Contents[2][1], m1.Contents[3][1]},
			{m1.Contents[0][2], m1.Contents[1][2], m1.Contents[2][2], m1.Contents[3][2]},
			{m1.Contents[0][3], m1.Contents[1][3], m1.Contents[2][3], m1.Contents[3][3]},
		},
	}
}

func Submatrix44(m1 *Matrix44, rowToRemove, columnToRemove int) *Matrix33 {
	res := [3][3]float64{}

	row := 0
	for i := 0; i < 4; i++ {
		if i != rowToRemove {
			switch columnToRemove {
			case 0:
				res[row] = [3]float64{m1.Contents[i][1], m1.Contents[i][2], m1.Contents[i][3]}
			case 1:
				res[row] = [3]float64{m1.Contents[i][0], m1.Contents[i][2], m1.Contents[i][3]}
			case 2:
				res[row] = [3]float64{m1.Contents[i][0], m1.Contents[i][1], m1.Contents[i][3]}
			case 3:
				res[row] = [3]float64{m1.Contents[i][0], m1.Contents[i][1], m1.Contents[i][2]}
			default:
				panic("Illegal column to remove")
			}
			row++
		}
	}

	return &Matrix33{Contents: res}
}

func Minor44(m1 *Matrix44, rowToRemove, columnToRemove int) float64 {
	return Determinant33(Submatrix44(m1, rowToRemove, columnToRemove))
}

func Cofactor44(m1 *Matrix44, rowToRemove, columnToRemove int) float64 {
	min := Minor44(m1, rowToRemove, columnToRemove)
	if (rowToRemove+columnToRemove)%2 == 1 {
		return -min
	}

	return min
}

func Determinant44(m1 *Matrix44) float64 {
	return m1.Contents[0][0]*Cofactor44(m1, 0, 0) +
		m1.Contents[0][1]*Cofactor44(m1, 0, 1) +
		m1.Contents[0][2]*Cofactor44(m1, 0, 2) +
		m1.Contents[0][3]*Cofactor44(m1, 0, 3)
}

func Inverse44(m1 *Matrix44) *Matrix44 {
	d := Determinant44(m1)
	return &Matrix44{Contents: [4][4]float64{
		{Cofactor44(m1, 0, 0) / d, Cofactor44(m1, 1, 0) / d, Cofactor44(m1, 2, 0) / d, Cofactor44(m1, 3, 0) / d},
		{Cofactor44(m1, 0, 1) / d, Cofactor44(m1, 1, 1) / d, Cofactor44(m1, 2, 1) / d, Cofactor44(m1, 3, 1) / d},
		{Cofactor44(m1, 0, 2) / d, Cofactor44(m1, 1, 2) / d, Cofactor44(m1, 2, 2) / d, Cofactor44(m1, 3, 2) / d},
		{Cofactor44(m1, 0, 3) / d, Cofactor44(m1, 1, 3) / d, Cofactor44(m1, 2, 3) / d, Cofactor44(m1, 3, 3) / d},
	}}
}

func Multiply33(m1, m2 *Matrix33) *Matrix33 {
	return &Matrix33{
		Contents: [3][3]float64{
			{
				m1.Contents[0][0]*m2.Contents[0][0] + m1.Contents[0][1]*m2.Contents[1][0] + m1.Contents[0][2]*m2.Contents[2][0],
				m1.Contents[0][0]*m2.Contents[0][1] + m1.Contents[0][1]*m2.Contents[1][1] + m1.Contents[0][2]*m2.Contents[2][1],
				m1.Contents[0][0]*m2.Contents[0][2] + m1.Contents[0][1]*m2.Contents[1][2] + m1.Contents[0][2]*m2.Contents[2][2],
			},
			{
				m1.Contents[1][0]*m2.Contents[0][0] + m1.Contents[1][1]*m2.Contents[1][0] + m1.Contents[1][2]*m2.Contents[2][0],
				m1.Contents[1][0]*m2.Contents[0][1] + m1.Contents[1][1]*m2.Contents[1][1] + m1.Contents[1][2]*m2.Contents[2][1],
				m1.Contents[1][0]*m2.Contents[0][2] + m1.Contents[1][1]*m2.Contents[1][2] + m1.Contents[1][2]*m2.Contents[2][2],
			},
			{
				m1.Contents[2][0]*m2.Contents[0][0] + m1.Contents[2][1]*m2.Contents[1][0] + m1.Contents[2][2]*m2.Contents[2][0],
				m1.Contents[2][0]*m2.Contents[0][1] + m1.Contents[2][1]*m2.Contents[1][1] + m1.Contents[2][2]*m2.Contents[2][1],
				m1.Contents[2][0]*m2.Contents[0][2] + m1.Contents[2][1]*m2.Contents[1][2] + m1.Contents[2][2]*m2.Contents[2][2],
			},
		},
	}
}

func Multiply33T(m1 *Matrix33, t1 *tuple.Tuple3) *tuple.Tuple3 {
	return &tuple.Tuple3{
		X:    m1.Contents[0][0]*t1.X + m1.Contents[0][1]*t1.Y + m1.Contents[0][2]*t1.Z,
		Y:    m1.Contents[1][0]*t1.X + m1.Contents[1][1]*t1.Y + m1.Contents[1][2]*t1.Z,
		Z:    m1.Contents[2][0]*t1.X + m1.Contents[2][1]*t1.Y + m1.Contents[2][2]*t1.Z,
		Type: t1.Type,
	}
}

func Transpose33(m1 *Matrix33) *Matrix33 {
	return &Matrix33{
		Contents: [3][3]float64{
			{m1.Contents[0][0], m1.Contents[1][0], m1.Contents[2][0]},
			{m1.Contents[0][1], m1.Contents[1][1], m1.Contents[2][1]},
			{m1.Contents[0][2], m1.Contents[1][2], m1.Contents[2][2]},
		},
	}
}

func Submatrix33(m1 *Matrix33, rowToRemove, columnToRemove int) *Matrix22 {
	res := [2][2]float64{}

	row := 0
	for i := 0; i < 3; i++ {
		if i != rowToRemove {
			switch columnToRemove {
			case 0:
				res[row] = [2]float64{m1.Contents[i][1], m1.Contents[i][2]}
			case 1:
				res[row] = [2]float64{m1.Contents[i][0], m1.Contents[i][2]}
			case 2:
				res[row] = [2]float64{m1.Contents[i][0], m1.Contents[i][1]}
			default:
				panic("Illegal column to remove")
			}
			row++
		}
	}

	return &Matrix22{Contents: res}
}

func Minor33(m1 *Matrix33, rowToRemove, columnToRemove int) float64 {
	return Determinant22(Submatrix33(m1, rowToRemove, columnToRemove))
}

func Cofactor33(m1 *Matrix33, rowToRemove, columnToRemove int) float64 {
	min := Minor33(m1, rowToRemove, columnToRemove)
	if (rowToRemove+columnToRemove)%2 == 1 {
		return -min
	}

	return min
}

func Determinant33(m1 *Matrix33) float64 {
	return m1.Contents[0][0]*Cofactor33(m1, 0, 0) + m1.Contents[0][1]*Cofactor33(m1, 0, 1) + m1.Contents[0][2]*Cofactor33(m1, 0, 2)
}

func Inverse33(m1 *Matrix33) *Matrix33 {
	d := Determinant33(m1)
	return &Matrix33{Contents: [3][3]float64{
		{Cofactor33(m1, 0, 0) / d, Cofactor33(m1, 1, 0) / d, Cofactor33(m1, 2, 0) / d},
		{Cofactor33(m1, 0, 1) / d, Cofactor33(m1, 1, 1) / d, Cofactor33(m1, 2, 1) / d},
		{Cofactor33(m1, 0, 2) / d, Cofactor33(m1, 1, 2) / d, Cofactor33(m1, 2, 2) / d},
	}}
}
