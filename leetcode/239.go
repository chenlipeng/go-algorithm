package leetcode

/*
 * 239. 滑动窗口最大值
 *
 * 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
 *
 * 使用单调队列处理, 队列中存储数组索引
 *
 *
 */

func MaxSlidingWindow(nums []int, k int) []int {
	return maxSlidingWindow(nums, k)
}

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}

	push := func(i int) {
		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	result := []int{}
	result = append(result, nums[q[0]])
	for i := k; i < len(nums); i++ {
		push(i)
		if q[0] < i-k+1 {
			q = q[1:]
		}

		result = append(result, nums[q[0]])
	}
	return result
}
