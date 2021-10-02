package leetcode

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	return threeSum_1(nums)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for k := j + 1; k < len(nums); k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}

				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}

func threeSum_1(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		k := len(nums) - 1
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for j < k {
				val := nums[i] + nums[j] + nums[k]
				if val == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					k--
					break
				} else if val > 0 {
					k--
				} else {
					break
				}
			}
		}
	}
	return result
}
