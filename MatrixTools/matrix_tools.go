package MatrixTools

import "math"

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

func (cur *Matrix) AddMatrix(other Matrix) {
	for i := 0; i < cur.N; i++ {
		for j := 0; j < cur.M; j++ {
			cur.A[i][j] += other.A[i][j]
		}
	}
}

func (cur *Matrix) MultScalar(l float64) {
	for i := 0; i < cur.N; i++ {
		for j := 0; j < cur.M; j++ {
			cur.A[i][j] *= l
		}
	}
}

func (cur *Matrix) Pow(b float64) {
	for i := 0; i < cur.N; i++ {
		for j := 0; j < cur.M; j++ {
			cur.A[i][j] = math.Pow(cur.A[i][j], b)
		}
	}
}

func (cur *Matrix) Norm() {
	mn := 1e9
	mx := -1e9
	for i := 0; i < cur.N; i++ {
		for j := 0; j < cur.M; j++ {
			mn = min(mn, cur.A[i][j])
			mx = max(mx, cur.A[i][j])
		}
	}
	for i := 0; i < cur.N; i++ {
		for j := 0; j < cur.M; j++ {
			cur.A[i][j] = (cur.A[i][j] - mn) / (mx - mn)
		}
	}
}
