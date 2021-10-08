package leetcode

func MaxSlidingWindow(nums []int, k int) []int {
	return maxSlidingWindow(nums, k)
}

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}

	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	result := make([]int, len(nums)-k+1)
	result[0] = nums[q[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		if q[0] < i-k+1 {
			q = q[1:]
		}
		result[i-k+1] = nums[q[0]]
	}
	return result
}
