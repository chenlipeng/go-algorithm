package leetcode

/*
 * 面试题 17.23. 最大黑方阵
 * 给定一个方阵，其中每个单元(像素)非黑即白。设计一个算法，找出 4 条边皆为黑色像素的最大子方阵。
 *
 * 返回一个数组 [r, c, size] ，其中 r, c 分别代表子方阵左上角的行号和列号，size 是子方阵的边长。若有多个满足条件的子方阵，返回 r 最小的，若 r 相同，返回 c 最小的子方阵。若无满足条件的子方阵，返回空数组。
 */
/*
func findSquare(matrix [][]int) []int {

	r, c, size := 0, 0, 0
	for rowIdx, row := range matrix {
		r1, c1, size1 := 0, 0, 0
		for colIdx, col := range row {
			if col == 1 {
				continue
			}

			r1, c1, size1 := rowIdx, colIdx, 1
			for r1+size1 < len(matrix) && c2+size1 < len(row) {
				ok := true
				for idx := c1; idx < c1+size1; idx++ {
				}
			}
		}
	}

	return ret
}
*/
