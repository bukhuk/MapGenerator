package MatrixTools

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func ToIMG(matrix [][]float64, path string) {
	matrix = Norm(matrix)
	img := image.NewGray(image.Rect(0, 0, len(matrix), len(matrix[0])))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			c := uint8(matrix[i][j] * 255)
			img.SetGray(i, j, color.Gray{Y: c})
		}
	}
	f, _ := os.Create(path)
	png.Encode(f, img)
}
