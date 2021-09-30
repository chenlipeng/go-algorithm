package leetcode

func FindRepeatNumber(nums []int) int {
	return find(nums, 0)
}

func find(nums []int, pos int) int {
	if pos == len(nums) {
		return 0
	}

	if nums[pos] != pos && nums[pos] == nums[nums[pos]] {
		return nums[pos]
	}
	nums[pos], nums[nums[pos]] = nums[nums[pos]], nums[pos]
	return find(nums, pos+1)
}
