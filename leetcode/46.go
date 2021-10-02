package leetcode

func Permute(nums []int) [][]int {
	return permute(nums)
}

//递归实现
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	result := [][]int{}
	for k, v := range nums {
		list := permute(append(nums[0:k:k], nums[k+1:len(nums):len(nums)]...))
		for _, item := range list {
			result = append(result, append(item[:len(item):len(item)], v))
		}
	}

	return result
}
