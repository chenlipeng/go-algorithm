package leetcode

func LongestPalindrome(s string) string {
	return longestPalindrome_2(s)
}

//中心扩散法
func longestPalindrome_2(s string) string {
	maxLen := 0
	maxPalindrome := ""

	for i := 0; i < len(s); i++ {
		if maxLen < 1 {
			maxLen = 1
			maxPalindrome = s[i : i+1]
		}

		low, high := i-1, i+1
		for low >= 0 && high < len(s) {
			if s[low] != s[high] {
				break
			}
			low--
			high++
		}

		l := high - low - 2 + 1
		if l > maxLen {
			maxLen = l
			maxPalindrome = s[low+1 : high]
		}
	}

	for i := 0; i < len(s)-1; i++ {
		low, high := i, i+1
		for low >= 0 && high < len(s) {
			if s[low] != s[high] {
				break
			}
			low--
			high++
		}

		l := high - low - 2 + 1
		if l > maxLen {
			maxLen = l
			maxPalindrome = s[low+1 : high]
		}
	}

	return maxPalindrome
}

/*
 * 暴力破解基础上做了优化
 * 当前已查找出的回文长度为L
 * 当备选字符串长度小于等于L时不再进行搜索
 */
func longestPalindrome_1(s string) string {

	maxLen := 0
	maxPalindrome := ""
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i+maxLen; j-- {
			low, high, ok := i, j, true
			for low <= high {
				if s[low] != s[high] {
					ok = false
					break
				}
				low++
				high--
			}

			if ok {
				l := j - i + 1
				if l > maxLen {
					maxLen = l
					maxPalindrome = s[i : j+1]
				}
			}
		}
	}

	return maxPalindrome
}

//方法一：暴力
func longestPalindrome(s string) string {

	maxLen := 0
	maxPalindrome := ""

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			low, high, ok := i, j, true
			for low <= high {
				if s[low] != s[high] {
					ok = false
					break
				}
				low++
				high--
			}

			if ok {
				l := j - i + 1
				if l > maxLen {
					maxLen = l
					maxPalindrome = s[i : j+1]
				}
			}
		}
	}

	return maxPalindrome
}
