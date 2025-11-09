package perlin

import (
	"fmt"
	"math/rand"
)

func NoiseMatrix(n, m, seed int) [][]float64 {
	if seed == 0 {
		seed = rand.Int()
		fmt.Println("seed:", seed)
	}
	matrix := make([][]float64, n)
	perm := randomPermutation(int64(seed))
	scale := 100.0
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = noise(float64(i)/scale, float64(j)/scale, perm)
		}
	}
	return matrix
}
