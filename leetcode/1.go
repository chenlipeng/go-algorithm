package leetcode

func TwoSum(nums []int, target int) []int {
	return twoSum_2(nums, target)
}

//方法一、暴力搜索时间O(n方), 空间O(1)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

//方法二、使用map 时间O(n), 空间O(n)
func twoSum_2(nums []int, target int) []int {
	mp := map[int]byte{}
	for _, v := range nums {
		if mp[v] >= 2 {
			continue
		}
		mp[v]++
	}

	for k, v1 := range nums {
		v2 := target - v1
		if (v1 == v2 && mp[v2] == 2) || (v1 != v2 && mp[v2] >= 1) {
			for ik, v := range nums {
				if k != ik && v == v2 {
					return []int{k, ik}
				}
			}
			break
		}
	}
	return []int{}
}

// 方法三、排序 + 二分查找 时间O(nlogn), 空间O(n) —— 不是最优方案
