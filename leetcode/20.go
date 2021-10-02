package leetcode

func IsValid(s string) bool {
	return isValid(s)
}

func isValid(s string) bool {
	stack := []rune{}

	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}

		top := stack[len(stack)-1]
		if top == '(' && v == ')' || top == '{' && v == '}' || top == '[' && v == ']' {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, v)
	}

	return len(stack) == 0
}
