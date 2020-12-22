package usort

func BubbleSort_1(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func SelectionSort_1(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func InsertionSort_1(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		currentValue := arr[i]
		for ; j >= 0 && arr[j] > currentValue; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = currentValue
	}
}

func InsertionSortn_1(arr []int, step int) {
	for i := step; i < len(arr); i++ {
		pos := i - step
		currentValue := arr[i]
		for ; pos >= 0 && arr[pos] > currentValue; pos -= step {
			arr[pos+step] = arr[pos]
		}
		arr[pos+step] = currentValue
	}
}

func ShellSort_1(arr []int) {
	for step := int(len(arr) / 2); step >= 1; step = int(step / 2) {
		InsertionSortn_1(arr, step)
	}
}

func MergeSort_1(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := int(len(arr) / 2)
	MergeSort_1(arr[0:mid])
	MergeSort_1(arr[mid:])

	left, right := 0, mid
	temp := []int{}
	for left < mid && right < len(arr) {
		if arr[left] < arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	temp = append(temp, arr[left:mid]...)
	temp = append(temp, arr[right:]...)

	for idx := 0; idx < len(arr); idx++ {
		arr[idx] = temp[idx]
	}
}

func QuickSort_1(arr []int) {
	for len(arr) <= 1 {
		return
	}

	pivot := arr[0]
	left, right := 0, len(arr)-1
	for left < right {
		for left < right && arr[right] > pivot {
			right--
		}
		arr[left] = arr[right]

		for left < right && arr[left] <= pivot {
			left++
		}

		arr[right] = arr[left]
	}

	arr[left] = pivot
	QuickSort_1(arr[0:left])
	QuickSort_1(arr[left+1:])
}

func HeapSort_1(arr []int) {
	buildHeap_1(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		buildHeap_1(arr[0:i])
	}
}

func buildHeap_1(arr []int) {
	for root := int(len(arr)/2) - 1; root >= 0; root-- {
		swap_1(arr, root)
	}
}

func swap_1(arr []int, root int) {
	if len(arr) <= root*2+1 {
		return
	}

	pos := root*2 + 1
	right := root*2 + 2

	if right < len(arr) && arr[right] > arr[pos] {
		pos = right
	}

	if arr[root] < arr[pos] {
		arr[pos], arr[root] = arr[root], arr[pos]
		swap_1(arr, pos)
	}
}
