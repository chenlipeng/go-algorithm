package leetcode

func Subsets(nums []int) [][]int {
	return subsets_1(nums)
}

func subsets(nums []int) [][]int {
	if len(nums) <= 0 {
		return [][]int{{}}
	}

	result1 := subsets(nums[1:])
	for k, _ := range result1 {
		result1[k] = append(result1[k], nums[0])
	}

	result2 := subsets(nums[1:])
	return append(result1, result2...)
}

func subsets_1(nums []int) [][]int {
	result := [][]int{{}}
	for _, v := range nums {
		for k, _ := range result {
			result = append(result, append(result[k][0:len(result[k]):len(result[k])], v))
		}
	}
	return result
}

/*
 * 错误示范
 * 		slice在append中的陷阱
 * 		append的第一个参数有足够空间时，不会新分配空间, 这会导致后面的元素被覆盖!!!!!
 */
func subsets_2(nums []int) [][]int {
	result := [][]int{{}}
	for _, v := range nums {
		for k, _ := range result {
			result = append(result, append(result[k], v))
		}
	}
	return result
}
