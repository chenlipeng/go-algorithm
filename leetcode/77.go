package leetcode

func Combine(n int, k int) [][]int {
	return combine(n, k)
}

func combine(n int, k int) [][]int {
	if 0 == k {
		return [][]int{{}}
	}
	if n == 0 || n < k {
		return [][]int{}
	}

	l1 := combine(n-1, k-1)
	for k, _ := range l1 {
		l1[k] = append(l1[k], n)
	}

	l2 := combine(n-1, k)
	list := append(l1, l2...)
	return list
}
