package MatrixTools

import (
	"fmt"
	"os"
)

func ToOBJ(matrix [][]int, path string) {
	f, _ := os.Create(path)
	defer f.Close()

	map_points := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Fprintln(f, "v", i, j, matrix[i][j])
			map_points++
		}
	}

	for i := 0; i < len(matrix); i++ {
		fmt.Fprintln(f, "v", i, 0, 0)
	}

	for i := 0; i < len(matrix); i++ {
		fmt.Fprintln(f, "v", i, len(matrix[i])-1, 0)
	}

	for i := 0; i < len(matrix[0]); i++ {
		fmt.Fprintln(f, "v", 0, i, 0)
	}

	for i := 0; i < len(matrix[len(matrix)-1]); i++ {
		fmt.Fprintln(f, "v", len(matrix)-1, i, 0)
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

	for i := 0; i < len(matrix)-1; i++ {
		v00 := map_points + i + 1
		v01 := map_points + i + 2
		v10 := i*len(matrix[0]) + 1
		v11 := (i+1)*len(matrix[0]) + 1
		fmt.Fprintln(f, "f", v00, v01, v11)
		fmt.Fprintln(f, "f", v00, v10, v11)
	}

	for i := 0; i < len(matrix)-1; i++ {
		v00 := map_points + len(matrix) + i + 1
		v01 := map_points + len(matrix) + i + 2
		v10 := (i + 1) * len(matrix[0])
		v11 := (i + 2) * len(matrix[0])
		fmt.Fprintln(f, "f", v00, v01, v11)
		fmt.Fprintln(f, "f", v00, v10, v11)
	}

	for i := 0; i < len(matrix[0])-1; i++ {
		v00 := map_points + 2*len(matrix) + i + 1
		v01 := map_points + 2*len(matrix) + i + 2
		v10 := i + 1
		v11 := i + 2
		fmt.Fprintln(f, "f", v00, v01, v11)
		fmt.Fprintln(f, "f", v00, v10, v11)
	}

	for i := 0; i < len(matrix[0])-1; i++ {
		v00 := map_points + 2*len(matrix) + len(matrix[0]) + i + 1
		v01 := map_points + 2*len(matrix) + len(matrix[0]) + i + 2
		v10 := (len(matrix)-1)*len(matrix[0]) + i + 1
		v11 := (len(matrix)-1)*len(matrix[0]) + i + 2
		fmt.Fprintln(f, "f", v00, v01, v11)
		fmt.Fprintln(f, "f", v00, v10, v11)
	}
	fmt.Fprintln(f, "f", map_points+1, map_points+len(matrix), map_points+2*len(matrix)+2*len(matrix[0]))
	fmt.Fprintln(f, "f", map_points+1, map_points+2*len(matrix)+len(matrix[0]), map_points+2*len(matrix)+2*len(matrix[0]))
}
