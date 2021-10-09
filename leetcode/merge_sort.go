package leetcode

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	MergeSort(left)
	MergeSort(right)

	result := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	for k, _ := range arr {
		arr[k] = result[k]
	}
	return
}
