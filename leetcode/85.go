package leetcode

func MaximalRectangle(matrix [][]byte) (ans int) {
	return maximalRectangle_1(matrix)
	return maximalRectangle(matrix)
}

func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalRectangle_1(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}

	left := make([][]int, len(matrix))
	for i, row := range matrix {
		left[i] = make([]int, len(row))
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	max := func(l, r int) int {
		if l > r {
			return l
		}
		return r
	}

	for j := 0; j < len(left[0]); j++ {
		l, r := make([]int, len(left)), make([]int, len(left))
		stack := []int{}
		for i := 0; i < len(left); i++ {
			for len(stack) > 0 && left[i][j] <= left[stack[len(stack)-1]][j] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				l[i] = -1
			} else {
				l[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}

		stack = []int{}
		for i := len(left) - 1; i >= 0; i-- {
			for len(stack) > 0 && left[i][j] <= left[stack[len(stack)-1]][j] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				r[len(left)-i-1] = len(left)
			} else {
				r[len(left)-i-1] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}

		for i := 0; i < len(left); i++ {
			ans = max(ans, left[i][j]*(r[i]-l[i]-1))
		}
	}

	return
}
