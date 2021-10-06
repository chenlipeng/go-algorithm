package leetcode

func LengthOfLastWord(s string) int {
	return lengthOfLastWord(s)
}

func lengthOfLastWord(s string) int {

	length, lastLength := 0, 0
	for idx := 0; idx < len(s); idx++ {
		if s[idx] == ' ' {
			if length > 0 {
				lastLength = length
			}
			length = 0
			continue
		}

		length++
	}

	if length > 0 {
		lastLength = length
	}

	return lastLength
}
