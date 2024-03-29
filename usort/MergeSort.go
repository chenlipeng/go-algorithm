package usort

/*
 * 归并排序（Merge Sort）
 * 		归并排序是建立在归并操作上的一种有效的排序算法。
 *		该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
 * 		将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。
 *
 * 算法描述
 * 		把长度为n的输入序列分成两个长度为n/2的子序列；
 * 		对这两个子序列分别采用归并排序；
 * 		将两个排序好的子序列合并成一个最终的排序序列。
 */

/*
 * @desc 归并排序
 */
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
