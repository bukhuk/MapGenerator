package MatrixTools

func (cur *Matrix) Map(low, high float64) {
	cur.Norm()
	for i := 0; i < cur.N; i++ {
		for j := 0; j < cur.M; j++ {
			cur.A[i][j] = low + cur.A[i][j]*(high-low)
		}
	}
}
