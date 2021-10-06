package leetcode

func Merge(intervals [][]int) [][]int {
	return merge(intervals)
}

//暴力破解
func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][1] < intervals[j][0] {
				continue
			}

			if intervals[i][0] > intervals[j][1] {
				continue
			}

			//i是j的子集
			if intervals[i][0] >= intervals[j][0] && intervals[i][1] <= intervals[j][1] {
				//Nothing todo
			} else if intervals[j][0] >= intervals[i][0] && intervals[j][1] <= intervals[i][1] {
				//j是i的子集
				intervals[j][0], intervals[j][1] = intervals[i][0], intervals[i][1]
			} else if intervals[i][1] >= intervals[j][0] && intervals[i][1] <= intervals[j][1] {
				//i的右侧在j中 j的左侧在i中
				intervals[j][0] = intervals[i][0]
			} else if intervals[j][1] >= intervals[i][0] && intervals[j][1] <= intervals[i][1] {
				//j的右侧在i中 i的左侧在j中
				intervals[j][1] = intervals[i][1]
			}
			intervals[i][0], intervals[i][1] = -1, -1
			break
		}
	}

	//result := [][]int{}
	//for _, v := range intervals {
	//	if v[0] == -1 {
	//		continue
	//	}
	//	result = append(result, v)
	//}

	idx := 0
	for k, v := range intervals {
		if v[0] == -1 {
			continue
		}
		if k != idx {
			intervals[idx][0], intervals[idx][1], intervals[k][0], intervals[k][1] = intervals[k][0], intervals[k][1], intervals[idx][0], intervals[idx][1]
		}
		idx++
	}
	return intervals[:idx]
}
