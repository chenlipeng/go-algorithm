package leetcode

func Multiply(num1 string, num2 string) string {
	return multiply(num1, num2)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := make([]byte, len(num1)+len(num2))
	minIdx := len(num1) + len(num2)
	for idx1 := len(num1) - 1; idx1 >= 0; idx1-- {
		extra := byte(0)
		digit1 := byte(num1[idx1] - '0')
		for idx2 := len(num2) - 1; idx2 >= 0; idx2-- {
			digit2 := byte(num2[idx2] - '0')

			curpos := result[idx1+idx2+1]
			if curpos >= '0' {
				curpos -= '0'
			}

			ret := digit1*digit2 + extra + curpos
			extra = ret / 10
			result[idx1+idx2+1] = ret%10 + '0'

			if idx1+idx2+1 < minIdx {
				minIdx = idx1 + idx2 + 1
			}
		}

		if extra > 0 {
			result[idx1] = extra + '0'
			if idx1 < minIdx {
				minIdx = idx1
			}
		}
	}

	return string(result[minIdx:])
}
