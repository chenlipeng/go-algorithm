package leetcode

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}

func searchMatrix(matrix [][]int, target int) bool {
	result := false

	r, c := len(matrix), len(matrix[0])

	low, high := 0, r*c-1
	for low <= high {
		mid := (low + high) / 2

		ridx := mid / c
		cidx := mid % c

		if matrix[ridx][cidx] == target {
			result = true
			break
		} else if matrix[ridx][cidx] > target {
			high = mid - 1
		} else if matrix[ridx][cidx] < target {
			low = mid + 1
		}
	}

	return result
}
