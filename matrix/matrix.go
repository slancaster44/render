package matrix

type Matrix interface {
	Get(row, col int) float64
	GetColumn(col int) []float64
	Put(row, col int, val float64)
	PutRow(rowNumber int, rowContents []float64)
	RowCount() int
	ColumnCount() int
	Copy() Matrix
}

func NewMatrix(rows, columns int) Matrix {
	if rows == 4 && columns == 4 {
		return &Matrix44{}
	} else if rows == 2 && columns == 2 {
		return &Matrix22{}
	} else if rows == 3 && columns == 3 {
		return &Matrix33{}
	}

	panic("Unsupported matrix dimensions")
}

type Matrix44 struct {
	Contents [4][4]float64
}

func (m *Matrix44) Get(row, col int) float64 {
	return m.Contents[row][col]
}

func (m *Matrix44) Put(row, col int, val float64) {
	m.Contents[row][col] = val
}

func (m *Matrix44) PutRow(rowNumber int, rowContents []float64) {
	for i := 0; i < 4; i++ {
		m.Contents[rowNumber][i] = rowContents[i]
	}
}

func (m *Matrix44) GetColumn(column int) []float64 {
	return []float64{
		m.Contents[0][column],
		m.Contents[1][column],
		m.Contents[2][column],
		m.Contents[3][column],
	}
}

func (m *Matrix44) Copy() Matrix {
	return &Matrix44{
		Contents: [4][4]float64{
			m.Contents[0],
			m.Contents[1],
			m.Contents[2],
			m.Contents[3],
		},
	}
}

func (m *Matrix44) RowCount() int {
	return 4
}

func (m *Matrix44) ColumnCount() int {
	return 4
}

type Matrix33 struct {
	Contents [3][3]float64
}

func (m *Matrix33) Get(row, col int) float64 {
	return m.Contents[row][col]
}

func (m *Matrix33) Put(row, col int, val float64) {
	m.Contents[row][col] = val
}

func (m *Matrix33) PutRow(rowNumber int, rowContents []float64) {
	for i := 0; i < 3; i++ {
		m.Contents[rowNumber][i] = rowContents[i]
	}
}

func (m *Matrix33) GetColumn(column int) []float64 {
	return []float64{
		m.Contents[0][column],
		m.Contents[1][column],
		m.Contents[2][column],
	}
}

func (m *Matrix33) Copy() Matrix {
	return &Matrix33{
		Contents: [3][3]float64{
			m.Contents[0],
			m.Contents[1],
			m.Contents[2],
		},
	}
}

func (m *Matrix33) RowCount() int {
	return 3
}

func (m *Matrix33) ColumnCount() int {
	return 3
}

type Matrix22 struct {
	Contents [2][2]float64
}

func (m *Matrix22) Get(row, col int) float64 {
	return m.Contents[row][col]
}

func (m *Matrix22) Put(row, col int, val float64) {
	m.Contents[row][col] = val
}

func (m *Matrix22) PutRow(rowNumber int, rowContents []float64) {
	for i := 0; i < 2; i++ {
		m.Contents[rowNumber][i] = rowContents[i]
	}
}

func (m *Matrix22) GetColumn(column int) []float64 {
	return []float64{
		m.Contents[0][column],
		m.Contents[1][column],
	}
}

func (m *Matrix22) Copy() Matrix {
	return &Matrix22{
		Contents: [2][2]float64{
			m.Contents[0],
			m.Contents[1],
		},
	}
}

func (m *Matrix22) RowCount() int {
	return 2
}

func (m *Matrix22) ColumnCount() int {
	return 22
}
