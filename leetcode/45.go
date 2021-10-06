package leetcode

func Trap(height []int) int {
	return trap(height)
}

func trap(height []int) int {
	low, high := 0, len(height)-1
	lowMax, highMax, result := 0, 0, 0

	//max函数
	max := func(l, r int) int {
		if l > r {
			return l
		}
		return r
	}

	for low < high {
		if height[low] > height[high] {
			result += max(highMax-height[high], 0)
			highMax = max(highMax, height[high])
			high--
		} else {
			result += max(lowMax-height[low], 0)
			lowMax = max(lowMax, height[high])
			low++
		}
	}

	return result
}
