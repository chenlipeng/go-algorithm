package leetcode

import "math"

func Reverse(x int) int {
	return reverse(x)
}

func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt64/10 || rev > math.MaxInt64/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}
