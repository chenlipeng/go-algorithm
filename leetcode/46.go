package leetcode

/*
 * 46. 全排列
 * 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 *
 *
 * 示例 1：
 * 		输入：nums = [1,2,3]
 * 		输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 * 示例 2：
 * 		输入：nums = [0,1]
 * 		输出：[[0,1],[1,0]]
 *
 * 示例 3：
 * 		输入：nums = [1]
 * 		输出：[[1]]
 *
 *
 *	解题思路:
 *		递归实现
 *		空间优化，直接使用切片作为参数传递；使需要处理的数据放在连续的位置
 *
 */

func Permute(nums []int) [][]int {
	return permute_1(nums)
	return permute(nums)
}

//递归实现 —— 第一个append处需要优化
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
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

//改进
func permute_1(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}

	result := [][]int{}
	for k, v := range nums {
		nums[0], nums[k] = nums[k], nums[0]
		list := permute_1(nums[1:])
		nums[k], nums[0] = nums[0], nums[k]
		for _, item := range list {
			result = append(result, append(item[:len(item):len(item)], v))
		}
	}

	return result
}
