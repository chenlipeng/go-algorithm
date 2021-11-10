package leetcode

func SeekKthNumber(arr1, arr2 []int, k int) int {
	return seekKthNumber(arr1, arr2, k)
}

/*
 *
 * 查找正序数组arr1和arr2的第k个元素
 *
 * 1. 取长度较短的数组arr1的第pos1 = min(k/2, l1) 个元素x;（从0开始）
 * 2. 取另一个数组arr2的第pos2 = k-pos1 个元素y;
 * 3. 比较arr1[pos1-1] 和 arr2[pos2-1] 元素的大小
 *		arr1[pos1-1] == arr2[pos2-1]：
 *			直接返回arr1[pos1-1]
 *		arr1[pos1-1] > arr2[pos2-1]：
 *			继续查找arr1和arr2[pos2]的第k-pos2个元素
 *		arr1[pos1-1] < arr2[pos2-1]：
 *			继续查找arr1[pos1]和arr2的第k-pos1个元素
 *
 * 特殊情况处理：
 *		其中一个数组为空时：直接返回另一个数组中第k-1个数字
 *		k = 1时： 返回两个数组中第一个元素里最小的那一个
 *
 */
func seekKthNumber(arr1, arr2 []int, k int) int {
	if len(arr1) == 0 {
		return arr2[k-1]
	} else if len(arr2) == 0 {
		return arr1[k-1]
	}

	min := func(l, r int) int {
		if l < r {
			return l
		}
		return r
	}

	if k == 1 {
		return min(arr1[0], arr2[0])
	}

	var pos1, pos2 int
	//pos要注意
	if len(arr1) > len(arr2) {
		pos2 = min(k/2, len(arr2))
		pos1 = min(k-pos2, len(arr1))
	} else {
		pos1 = min(k/2, len(arr1))
		pos2 = min(k-pos1, len(arr2))
	}
	if arr1[pos1-1] == arr2[pos2-1] {
		return arr1[pos1-1]
	}

	if arr1[pos1-1] > arr2[pos2-1] {
		return seekKthNumber(arr1, arr2[pos2:], k-pos2)
	}

	return seekKthNumber(arr1[pos1:], arr2, k-pos1)
}
