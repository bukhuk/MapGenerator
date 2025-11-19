package MatrixTools

type Matrix struct {
	N, M int
	A    [][]float64
}

func Make(n, m int, value float64) Matrix {
	values := make([][]float64, n)
	for i := 0; i < n; i++ {
		values[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			values[i][j] = value
		}
	}
	return Matrix{N: n, M: m, A: values}
}
