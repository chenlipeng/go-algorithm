package leetcode

/*
 * 47. 全排列 II
 * 	给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 *
 * 	示例 1：
 * 		输入：nums = [1,1,2]
 * 		输出：
 * 			[[1,1,2], [1,2,1], [2,1,1]]
 *
 * 	示例 2：
 * 		输入：nums = [1,2,3]
 * 		输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *	解题思路:
 *		基于【46.全排列】进行改进
 *		使用map记录当前位置是否放过相同的值, 若放过则跳过
 */

func PermuteUnique(nums []int) (ans [][]int) {
	return permuteUnique(nums)
}

func permuteUnique(nums []int) (ans [][]int) {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}

	result := [][]int{}
	mp := map[int]bool{}
	for k, v := range nums {
		if mp[v] {
			continue
		}
		mp[v] = true

		nums[0], nums[k] = nums[k], nums[0]
		list := permuteUnique(nums[1:])
		nums[k], nums[0] = nums[0], nums[k]
		for _, item := range list {
			result = append(result, append(item[:len(item):len(item)], v))
		}
	}

	return result
}
