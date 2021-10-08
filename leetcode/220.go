package leetcode

func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	return containsNearbyAlmostDuplicate(nums, k, t)
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	abs := func(i int) int {
		if i < 0 {
			return -1 * i
		}
		return i
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < i+k+1 && j < len(nums); j++ {
			diff := abs(nums[i] - nums[j])
			if diff <= t {
				return true
			}
		}
	}

	return false
}
