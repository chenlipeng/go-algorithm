package usort

/*
 * 堆排序（Heap Sort）
 * 		堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。
 *
 * 算法描述
 * 		将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
 * 		将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
 * 		由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。
 */

func HeapSort(arr []int) {
	buildHeap(arr)
	for i := 0; i < len(arr)-1; i++ {
		arr[0], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[0]
		//buildHeap(arr[:len(arr)-1-i])
		swap(arr[:len(arr)-1-i], 0)
	}
}

func buildHeap(arr []int) {
	//if len(arr) <= 1 {
	//	return
	//}

	//Tips. 从叶节点到根节点
	for i := int(len(arr)/2) - 1; i >= 0; i-- {
		swap(arr, i)
	}
}

func swap(arr []int, root int) {
	left := root<<1 + 1
	if left > len(arr)-1 {
		return
	}
	right := left + 1

	var pos int
	if right > len(arr)-1 {
		pos = left
	} else if arr[left] >= arr[right] {
		pos = left
	} else if arr[left] < arr[right] {
		pos = right
	} else {
		return
	}

	if arr[root] < arr[pos] {
		arr[root], arr[pos] = arr[pos], arr[root]
		swap(arr, pos)
	}
}
