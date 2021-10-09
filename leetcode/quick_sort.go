package leetcode

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	low, high, flag := 0, len(arr)-1, arr[0]
	for low < high {
		for arr[high] > flag && low < high {
			high--
		}
		arr[low] = arr[high]

		for arr[low] <= flag && low < high {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = flag

	QuickSort(arr[:low])
	QuickSort(arr[low+1:])
}
