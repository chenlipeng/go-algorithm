package leetcode

func IntegerBreak(n int) int {
	return integerBreak(n)
}

func integerBreak(n int) int {
	mp := make(map[int]int, n+1)
	mp[0] = 1
	mp[1] = 1

	max := func(l, r int) int {
		if l > r {
			return l
		}
		return r
	}

	for i := 2; i <= n; i++ {
		maxRet := 0
		for j := 1; j < i; j++ {
			maxRet = max(maxRet, max(j*(i-j), j*mp[i-j]))
		}
		mp[i] = maxRet
	}

	return mp[n]
}
