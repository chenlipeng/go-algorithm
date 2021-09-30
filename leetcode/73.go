package leetcode

func SetZeroes(matrix [][]int) {
	setZeroes(matrix)
}

func setZeroes(matrix [][]int) {
	col0, row0 := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}

	for _, row := range matrix {
		if row[0] == 0 {
			col0 = true
			break
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row0 {
		for k, _ := range matrix[0] {
			matrix[0][k] = 0
		}
	}
	if col0 {
		for k, _ := range matrix {
			matrix[k][0] = 0
		}
	}
}
