package leetcode

/*
 * 41. 缺失的第一个正数
 * 		给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
 * 		请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
 *
 *
 * 	示例 1：
 * 		输入：nums = [1,2,0]
 * 		输出：3
 *
 * 	示例 2：
 * 		输入：nums = [3,4,-1,1]
 * 		输出：2
 *
 * 	示例 3：
 * 		输入：nums = [7,8,9,11,12]
 * 		输出：1
 *
 *	解题思路:
 *		修改原数组
 *		遍历一遍数组，将nums[idx]放到切片idx-1的位置
 *		第一个nums[idx] != idx+1 即为所求, 否则为len(nums)+1
 */

func FirstMissingPositive(nums []int) int {
	return firstMissingPositive(nums)
}

func firstMissingPositive(nums []int) int {
	idx := 0
	for idx < len(nums) {
		if nums[idx] <= 0 || nums[idx] > len(nums) {
			nums[idx] = 0
			idx++
			continue
		}

		if nums[idx] == idx+1 {
			idx++
			continue
		}

		if nums[idx] == nums[nums[idx]-1] {
			nums[idx] = 0
			idx++
			continue
		}

		nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
	}

	for k, v := range nums {
		if v <= 0 {
			return k + 1
		}
	}

	return len(nums) + 1
}
