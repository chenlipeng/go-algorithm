package leetcode

func LargestRectangleArea(heights []int) int {
	return largestRectangleArea_1(heights)
	return largestRectangleArea(heights)
}

/*
 * 采用单调栈求解
 * left 和 right 两个
 */
func largestRectangleArea(heights []int) int {
	result := 0
	stack, left, right := []int{}, []int{}, []int{}

	for i := 0; i < len(heights); i++ {
		for j := len(stack) - 1; j >= 0; j-- {
			if heights[stack[j]] < heights[i] {
				break
			}
			stack = stack[:j]
		}

		if len(stack) == 0 {
			left = append(left, -1)
		} else {
			left = append(left, stack[len(stack)-1])
		}
		stack = append(stack, i)
	}

	stack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for j := len(stack) - 1; j >= 0; j-- {
			if heights[stack[j]] < heights[i] {
				break
			}
			stack = stack[:j]
		}

		if len(stack) == 0 {
			right = append(right, len(heights))
		} else {
			right = append(right, stack[len(stack)-1])
		}
		stack = append(stack, i)
	}

	max := func(l, r int) int {
		if l > r {
			return l
		}
		return r
	}

	for k, v := range heights {
		result = max(result, (right[len(heights)-k-1]-left[k]-1)*v)
	}
	return result
}

func largestRectangleArea_1(heights []int) int {
	result := 0
	stack, left, right := []int{}, make([]int, len(heights)), make([]int, len(heights))

	for i := 0; i < len(heights); i++ {
		for j := len(stack) - 1; j >= 0; j-- {
			if heights[stack[j]] < heights[i] {
				break
			}
			stack = stack[:j]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for j := len(stack) - 1; j >= 0; j-- {
			if heights[stack[j]] < heights[i] {
				break
			}
			stack = stack[:j]
		}

		if len(stack) == 0 {
			right[i] = len(heights)
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	max := func(l, r int) int {
		if l > r {
			return l
		}
		return r
	}

	for k, v := range heights {
		result = max(result, (right[k]-left[k]-1)*v)
	}
	return result
}
