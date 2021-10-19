package leetcode

func HIndex1(citations []int) int {
	return hIndex1(citations)
}

func hIndex1(citations []int) int {
	result := 0
	low, high := 0, len(citations)-1
	for low <= high {
		mid := (low + high) / 2
		if citations[mid] >= len(citations)-mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}
