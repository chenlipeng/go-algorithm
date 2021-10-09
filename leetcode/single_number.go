package leetcode

func SingleNumber(nums []int) int {
	return singleNumber(nums)
}

func singleNumber(nums []int) int {
	and := 0
	for _, v := range nums {
		and = (and | v) & ^(and & v)
	}

	return and
}
