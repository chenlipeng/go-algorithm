package leetcode

import "fmt"

//var mp map[string]int = make(map[string]int)
var mp map[string]int

func UniquePaths(m int, n int) int {
	mp = make(map[string]int)
	return uniquePaths(m, n)
}

func uniquePaths(m int, n int) int {
	return runUniquePaths(m, n)
}

func runUniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	max := func(l, r int) int {
		if l > r {
			return l
		}
		return r
	}

	min := func(l, r int) int {
		if l < r {
			return l
		}
		return r
	}

	if v, ok := mp[fmt.Sprintf("%d_%d", max(m, n), min(m, n))]; ok {
		return v
	}

	l := uniquePaths(m-1, n)
	mp[fmt.Sprintf("%d_%d", max(m-1, n), min(m-1, n))] = l
	r := uniquePaths(m, n-1)
	mp[fmt.Sprintf("%d_%d", max(m, n-1), min(m, n-1))] = r
	return l + r
}
