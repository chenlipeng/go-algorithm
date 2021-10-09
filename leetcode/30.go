package leetcode

import "fmt"

/*
 * 30. 串联所有单词的子串
 * 	给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
 * 	注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。
 *
 * 	示例 1：
 * 		输入：s = "barfoothefoobarman", words = ["foo","bar"]
 * 		输出：[0,9]
 *
 * 		解释：
 * 		从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
 * 		输出的顺序不重要, [9,0] 也是有效答案。
 *
 * 	示例 2：
 * 		输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
 * 		输出：[]
 *
 * 	示例 3：
 * 		输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
 * 		输出：[6,9,12]
 *
 *
 *	解题思路
 *		重点：长度相同的单词words
 */

func FindSubstring(s string, words []string) []int {
	return findSubstring_1(s, words)
	return findSubstring(s, words)
}

func findSubstring(s string, words []string) []int {
	result := []int{}

	wordsMap := make(map[string]int)
	for _, v := range words {
		wordsMap[v]++
	}

	wordsLen := len(words) * len(words[0])
	for i := 0; i <= len(s)-wordsLen; i++ {
		mp := make(map[string]int)
		for j := i; j < i+wordsLen; j += len(words[0]) {
			mp[s[j:j+len(words[0])]]++
		}

		flag := true
		for k, v := range mp {
			if wordsMap[k] != v {
				flag = false
				break
			}
		}

		if flag {
			result = append(result, i)
		}
	}
	return result
}

func findSubstring_1(s string, words []string) []int {
	result := []int{}

	wordsMap := make(map[string]int)
	for _, v := range words {
		wordsMap[v]++
	}

	fmt.Println(wordsMap)

	wordsLen := len(words) * len(words[0])
	for i := 0; i < len(words[0]) && i < len(s); i++ {
		mp := make(map[string]int)

		for j := i; j < len(s)-wordsLen+1; j += len(words[0]) {
			if j == i {
				for k := j; k < j+wordsLen; k += len(words[0]) {
					mp[s[k:k+len(words[0])]]++
				}
			} else {
				mp[s[j+wordsLen-len(words[0]):j+wordsLen]]++
				mp[s[j-len(words[0]):j]]--
			}

			fmt.Println(j, mp)

			flag := true
			for k, v := range mp {
				if wordsMap[k] != v {
					flag = false
					break
				}
			}

			if flag {
				result = append(result, j)
			}
		}
	}
	return result
}
