package leetcode

func FindDuplicates(nums []int) []int {
	return findDuplicates(nums)
}

/*
 * 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
 * 找到所有出现两次的元素。
 * 你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/find-all-duplicates-in-an-array
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
/*
01. 在原数组上进行操作
02. 将nums[idx]取负值, idx与nums[idx]-1位置上的值交换
03. idx位置的元素为负值
*/
func findDuplicates(nums []int) []int {
	ret := []int{}

	for idx := 0; idx < len(nums); {
		if nums[idx] <= 0 {
			idx++
			continue
		}

		if idx+1 == nums[idx] {
			nums[idx] *= -1
			idx++
			continue
		}

		if nums[idx]+nums[nums[idx]-1] == 0 || nums[idx] == nums[nums[idx]-1] {
			ret = append(ret, nums[idx])

			nums[nums[idx]-1] = 0
			nums[idx] = 0
			idx++
			continue
		}

		nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], -1*nums[idx]
	}

	return ret
}
