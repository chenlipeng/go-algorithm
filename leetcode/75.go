package leetcode

func SortColors(nums []int) {
	sortColors(nums)
}

func sortColors(nums []int) {

	slice := []int{0, 0, 0}
	for _, v := range nums {
		slice[v] += 1
	}

	idx := 0
	for k, v := range slice {
		for i := 0; i < v; i++ {
			nums[idx] = k
			idx++
		}
	}

}
