package nowcoder

/*
 * 描述
 * 给定一个长度为 n 的数组，请你编写一个函数，返回该数组排序后的结果。
 *
 * 数据范围： ，数组中每个元素都满足
 * 要求：空间复杂度O(n) ，时间复杂度O(nlogn)
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 将给定数组排序
 * @param arr int整型一维数组 待排序的数组
 * @return int整型一维数组
 */
func MySort(arr []int) []int {
	// write code here
	MergeSort(arr)
	return arr
}

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := int(len(arr) / 2)
	MergeSort(arr[0:mid])
	MergeSort(arr[mid:])

	temp := []int{}
	i, j := 0, mid
	for i < mid && j < len(arr) {
		if arr[i] < arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}

	temp = append(temp, arr[i:mid]...)
	temp = append(temp, arr[j:]...)

	for i := 0; i < len(arr); i++ {
		arr[i] = temp[i]
	}
}
