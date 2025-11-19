package perlin

import (
	"MapGenerator/MatrixTools"
	"math/rand"
)

func NoiseMatrix(n, m int, scale float64, seed int) MatrixTools.Matrix {
	if seed == 0 {
		seed = rand.Int()
	}
	matrix := MatrixTools.Make(n, m, 0)
	perm := randomPermutation(int64(seed))
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix.A[i][j] = noise(float64(i)/scale, float64(j)/scale, perm)
		}
	}
	return matrix
}
