package leetcode

func WordPattern(pattern string, s string) bool {
	return wordPattern(pattern, s)
}

func wordPattern(pattern string, s string) bool {
	mp := make(map[byte][]int, len(pattern))

	i, j := 0, 0
	for ; i < len(pattern); i++ {
		target, ok := mp[pattern[i]]

		low := j
		for j < len(s) && s[j] != ' ' {
			if ok && (target[1]-target[0] < j-low || s[target[0]+j-low] != s[j]) {
				return false
			}
			j++
		}

		if ok && j-low != target[1]-target[0] {
			return false
		}

		mp[pattern[i]] = []int{low, j}
		j++
	}

	if i != len(pattern) || j != len(s) {
		return false
	}

	return true
}
