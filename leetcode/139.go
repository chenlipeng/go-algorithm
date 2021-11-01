package leetcode

func WordBreak(s string, wordDict []string) bool {
	return wordBreak(s, wordDict)
}

func wordBreak(s string, wordDict []string) bool {
	var statusMap = make(map[int]bool)
	return runWordBreak(s, wordDict, 0, statusMap)
}

func runWordBreak(s string, wordDict []string, idx int, statusMap map[int]bool) bool {
	if len(s) == 0 {
		return true
	}

	if _, ok := statusMap[idx]; ok {
		return false
	}

	for _, v := range wordDict {
		l := len(v)
		if l > len(s) || s[0:l] != v {
			continue
		}

		result := runWordBreak(s[l:], wordDict, idx+l, statusMap)
		if result == true {
			return true
		}
	}

	statusMap[idx] = false
	return false
}
