package matrix

var IdentityMatrix33 Matrix33 = Matrix33{
	Contents: [3][3]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	},
}

var IdentityMatrix44 Matrix44 = Matrix44{
	Contents: [4][4]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	},
}
