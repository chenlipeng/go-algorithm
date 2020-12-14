package usort

/*
 * 快速排序
 */
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	low, high := 0, len(arr)-1
	pivotValue := arr[0]
	for low < high {
		for low < high && arr[high] > pivotValue {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] <= pivotValue {
			low++
		}
		arr[high] = arr[low]
	}

	arr[low] = pivotValue

	QuickSort(arr[0:low])
	QuickSort(arr[low+1:])
}
