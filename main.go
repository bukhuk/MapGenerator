package main

import (
	"MapGenerator/MatrixTools"
	perlin "MapGenerator/PerlinNoise"
)

func main() {
	n, m := 512, 512
	matrix := make([][]float64, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = 0.
		}
	}
	scale := 256.
	k := 1.0
	for scale > 1. {
		matrix = MatrixTools.AddMatrix(matrix, MatrixTools.MultScalar(perlin.NoiseMatrix(n, m, scale, 0), k))
		k /= 1.5
		scale /= 2.
	}
	
	MatrixTools.ToIMG(matrix, "perlin.png")
}
