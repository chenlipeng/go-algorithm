package leetcode

func MaxArea(height []int) int {
	return maxArea_1(height)
}

func maxArea_1(height []int) int {
	maxArea := 0

	low, high := 0, len(height)-1
	for low <= high {
		area := (high - low) * min(height[low], height[high])
		if area > maxArea {
			maxArea = area
		}

		if height[low] < height[high] {
			low++
		} else {
			high--
		}
	}

	return maxArea
}

func maxArea(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min(height[i], height[j])
			if maxArea < area {
				maxArea = area
			}
		}
	}

	return maxArea
}
