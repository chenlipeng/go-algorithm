package leetcode

func ContainsNearbyDuplicate(nums []int, k int) bool {
	return containsNearbyDuplicate(nums, k)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		idx, ok := mp[nums[i]]
		if ok && i-idx <= k {
			return true
		}

		mp[nums[i]] = i
	}
	return false
}
