package Visualization

import (
	"MapGenerator/MatrixTools"
	"image"
	"image/color"
	"image/png"
	"os"
)

func MakeIMG(matrix MatrixTools.Matrix, path string) {
	matrix.Norm()
	img := image.NewGray(image.Rect(0, 0, matrix.N, matrix.M))
	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			c := uint8(matrix.A[i][j] * 255)
			img.SetGray(i, j, color.Gray{Y: c})
		}
	}
	f, _ := os.Create(path)
	png.Encode(f, img)
}
