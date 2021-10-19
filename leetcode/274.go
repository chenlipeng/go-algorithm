package leetcode

import (
	"fmt"
	"sort"
)

func HIndex(citations []int) int {
	return hIndex(citations)
}

func hIndex(citations []int) int {
	sort.Ints(citations)
	fmt.Println(citations)

	result := 0
	for idx := 0; idx < len(citations); idx++ {
		if citations[len(citations)-1-idx] >= result+1 {
			result++
		} else {
			break
		}
	}
	return result
}
