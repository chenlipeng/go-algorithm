package leetcode

func lengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring_1(s)
}

func lengthOfLongestSubstring_1(s string) int {
	maxLen := 0
	mp := make(map[rune]int)
	beg := -1
	for k, v := range s {
		pos, ok := mp[v]
		if ok && pos > beg {
			beg = pos
		}
		mp[v] = k

		if k-beg > maxLen {
			maxLen = k - beg
		}
	}
	return maxLen
}

func lengthOfLongestSubstring_2(s string) int {
	return 0
}
