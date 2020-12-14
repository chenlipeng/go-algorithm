package usort

/*
 * @desc 插入排序
 */
func InsertionSort(arr []int) {
	l := len(arr)
	for i := 1; i < l; i++ {
		currentValue := arr[i]
		pos := i - 1
		for ; pos >= 0; pos-- {
			if arr[pos] < currentValue {
				arr[pos+1] = arr[pos]
			} else { //Tips. break
				break
			}
		}
		arr[pos+1] = currentValue
	}
}
