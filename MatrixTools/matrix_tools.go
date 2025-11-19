package MatrixTools

import "math"

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
