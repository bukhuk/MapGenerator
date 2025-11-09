package main

import (
	"MapGenerator/PerlinNoise"
	"fmt"
)

func main() {
	matrix := perlin.NoiseMatrix(5, 5, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
