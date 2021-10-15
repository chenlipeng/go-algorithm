package leetcode

func ReverseVowels(s string) string {
	return reverseVowels(s)
}

func reverseVowels(s string) string {
	low, high := 0, len(s)-1
	result := []byte(s)

	for low < high {
		for low < high && !isVowel(s[low]) {
			low++
		}

		for low < high && !isVowel(s[high]) {
			high--
		}

		if low == high {
			break
		}

		result[low], result[high] = s[high], s[low]
		low++
		high--
	}
	return string(result)
}

func isVowel(letter byte) bool {
	mp := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	return mp[letter]
}
