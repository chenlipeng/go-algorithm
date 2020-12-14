package usort

/*
 * @desc 选择排序
 */
func SelectionSort(arr []int) {
	l := len(arr)

	for i := 0; i < l-1; i++ {
		minValueIdx := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[minValueIdx] {
				minValueIdx = j
			}
		}
		arr[i], arr[minValueIdx] = arr[minValueIdx], arr[i]
	}
}
