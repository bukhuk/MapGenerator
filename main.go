package main

import (
	"MapGenerator/MatrixTools"
	perlin "MapGenerator/PerlinNoise"
	"math"
)

func main() {
	n, m := 512, 512
	matrix := make([][]float64, n)
	steps := 16
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

	matrix = MatrixTools.Norm(matrix)

	matrix = MatrixTools.Pow(matrix, 1.3)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j] = math.Round(matrix[i][j]*float64(steps)) / float64(steps)
		}
	}

	matrix = MatrixTools.ThermalErosion(matrix, 100, 0.05, 0.1)

	//matrix = MatrixTools.BoxBlur(matrix, 2)

	MatrixTools.ToIMG(matrix, "perlin.png")
}
