package leetcode

func FindSubsequences(nums []int) [][]int {
	return findSubsequences(nums)
}

/*
 * 使用一个数组保存中间结果
 * 记录中间结果的数组大小L
 * 遍历数组中的元素，如果对该元素进行追加Nums[i]后仍递增，则将结果作为新元素插入数组
 * 如果当前元素i和其前一个元素i-1相同，则从L开始遍历
 * 如果Nums[i]未插入过中间结果，则插入
 */
func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	finalResult := [][]int{}
	mp := make(map[int]int)

	for idx, _ := range nums {
		nstartPos := len(result) - 1
		for k, v := range result {
			if nums[idx] < v[len(v)-1] {
				continue
			}

			if pos, ok := mp[nums[idx]]; ok && k <= pos {
				continue
			}

			newItem := append(v[:len(v):len(v)], nums[idx])
			result = append(result, newItem)

			if len(newItem) >= 2 {
				finalResult = append(finalResult, newItem)
			}
		}

		if _, ok := mp[nums[idx]]; !ok {
			result = append(result, []int{nums[idx]})
		}
		mp[nums[idx]] = nstartPos
	}

	return finalResult
}

/*
 * 使用一个数组保存中间结果
 * 记录中间结果的数组大小L
 * 遍历数组中的元素，如果对该元素进行追加Nums[i]后仍递增，则将结果作为新元素插入数组
 * 如果当前元素i和其前一个元素i-1相同，则从L开始遍历
 * 如果Nums[i]未插入过中间结果，则插入
 */
func findSubsequences3(nums []int) [][]int {
	result := [][]int{}
	finalResult := [][]int{}
	mp := make(map[int]int)

	var dfs func(idx, pIdx int)
	dfs = func(idx, pIdx int) {
		if len(nums) == idx {
			return
		}
		nstartPos := len(result) - 1
		for k, v := range result {
			if nums[idx] < v[len(v)-1] {
				continue
			}

			if pos, ok := mp[nums[idx]]; ok && k <= pos {
				continue
			}

			newItem := append(v[:len(v):len(v)], nums[idx])
			result = append(result, newItem)

			if len(newItem) >= 2 {
				finalResult = append(finalResult, newItem)
			}
		}

		if _, ok := mp[nums[idx]]; !ok {
			result = append(result, []int{nums[idx]})
		}
		mp[nums[idx]] = nstartPos
		dfs(idx+1, idx)
	}

	dfs(0, -1)
	return finalResult
}

/*
 * 使用一个数组保存中间结果
 * 记录中间结果的数组大小L
 * 遍历数组中的元素，如果对该元素进行追加Nums[i]后仍递增，则将结果作为新元素插入数组
 * 如果当前元素i和其前一个元素i-1相同，则从L开始遍历
 * 如果Nums[i]未插入过中间结果，则插入
 */
func findSubsequences2(nums []int) [][]int {
	result := [][]int{}
	finalResult := [][]int{}
	mp := make(map[int]int)

	var dfs func(idx, pIdx int)
	dfs = func(idx, pIdx int) {
		if len(nums) == idx {
			return
		}

		startPos := -1
		if pos, ok := mp[nums[idx]]; ok {
			startPos = pos
		}
		nstartPos := len(result) - 1

		flag := (pIdx != -1 && nums[idx] == nums[pIdx])
		for k, v := range result {
			if nums[idx] < v[len(v)-1] {
				continue
			}

			if flag && k <= startPos {
				continue
			}

			newItem := append(v[:len(v):len(v)], nums[idx])
			result = append(result, newItem)
			if len(newItem) >= 2 {
				finalResult = append(finalResult, newItem)
			}
		}

		if _, ok := mp[nums[idx]]; !ok {
			result = append(result, []int{nums[idx]})
		}
		mp[nums[idx]] = nstartPos
		dfs(idx+1, idx)
	}

	dfs(0, -1)
	return finalResult
}

func findSubsequences1(nums []int) [][]int {
	result := [][]int{}
	finalResult := [][]int{}
	startPos := -1
	mp := make(map[int]bool)

	var dfs func(idx, pIdx int)
	dfs = func(idx, pIdx int) {
		if len(nums) == idx {
			return
		}

		nstartPos := len(result) - 1
		flag := (pIdx != -1 && nums[idx] == nums[pIdx])
		for k, v := range result {
			if nums[idx] < v[len(v)-1] {
				continue
			}

			if flag && k <= startPos {
				continue
			}

			newItem := append(v[:len(v):len(v)], nums[idx])
			result = append(result, newItem)

			if len(newItem) >= 2 {
				finalResult = append(finalResult, newItem)
			}
		}

		if _, ok := mp[nums[idx]]; !ok {
			result = append(result, []int{nums[idx]})
			mp[nums[idx]] = true
		}
		startPos = nstartPos
		dfs(idx+1, idx)
	}

	dfs(0, -1)
	return finalResult
}
