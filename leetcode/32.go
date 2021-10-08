package leetcode

/*
 * 32. 最长有效括号
 * 	给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
 *
 * 示例 1：
 * 		输入：s = "(()"
 * 		输出：2
 * 		解释：最长有效括号子串是 "()"
 *
 * 示例 2：
 * 		输入：s = ")()())"
 * 		输出：4
 * 		解释：最长有效括号子串是 "()()"
 *
 * 示例 3：
 * 		输入：s = ""
 * 		输出：0
 *
 *
 * 解题思路：
 *		使用栈来实现
 *		栈中存储元素索引
 *		当栈顶元素与即将入栈元素有效匹配时（即：栈顶元素为'(', 入栈元素为')'）, 栈顶元素出栈，即将入栈元素也不再入栈
 *			不匹配时: 入栈
 *		全部遍历完成后，栈中仍有的元素为无效元素下标，通过此计算合法元素长度
 *		注意：遍历完成后，将字符串长度入栈，统一逻辑计算
 *
 *		时间复杂度：O(n)
 *		空间复杂度：O(n)
 *
 */

func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}

func longestValidParentheses(s string) int {
	stack := []int{}
	maxLen := 0

	for k, v := range s {
		if len(stack) == 0 {
			stack = append(stack, k)
			continue
		}

		if v == '(' {
			stack = append(stack, k)
		} else if v == ')' {
			stackTopElement := s[stack[len(stack)-1]]
			if stackTopElement == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, k)
			}
		}
	}

	max := func(l, r int) int {
		if l > r {
			return l
		}
		return r
	}

	stack = append(stack, len(s))
	for i := 0; i < len(stack); i++ {
		if i == 0 {
			maxLen = max(maxLen, stack[i])
		} else {
			maxLen = max(maxLen, stack[i]-stack[i-1]-1)
		}
	}
	return maxLen
}
