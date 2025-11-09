package MatrixTools

import (
	"fmt"
	"os"
)

func ToOBJ(matrix [][]int, path string) {
	f, _ := os.Create(path)
	defer f.Close()

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Fprintln(f, "v", i, j, matrix[i][j])
		}
	}

	fmt.Fprintln(f)
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			v00 := i*len(matrix) + j + 1
			v01 := i*len(matrix) + j + 2
			v10 := (i+1)*len(matrix) + j + 1
			v11 := (i+1)*len(matrix) + j + 2
			fmt.Fprintln(f, "f", v00, v01, v11)
			fmt.Fprintln(f, "f", v00, v10, v11)
		}
	}
}
