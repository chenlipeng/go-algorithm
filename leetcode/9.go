package leetcode

func IsPalindrome(x int) bool {
	return isPalindrome1(x)
	return isPalindrome(x)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0
	for i := x; i != 0; i = i / 10 {
		reverse = (i % 10) + reverse*10
	}

	return reverse == x
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var arr []byte
	for i := x; i != 0; i = i / 10 {
		arr = append(arr, byte(i%10))
	}

	for low, high := 0, len(arr)-1; low < high; {
		if arr[low] != arr[high] {
			return false
		}

		low++
		high--
	}

	return true
}
