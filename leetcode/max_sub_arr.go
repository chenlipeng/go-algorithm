package leetcode

import "fmt"

func MaxSubArr1(arr []int) (mx int) {

	mx = arr[0]
	preSum := 0
	for idx, val := range arr {
		if idx == 0 {
			if val > 0 {
				preSum = val
			}
			continue
		}

		preSum += val
		if preSum > 0 {
			if preSum > mx {
				mx = preSum
			}
		} else {
			preSum = 0
		}
	}

	return mx
}

func MaxSubArr(arr []int) (mx int) {
	fmt.Println(arr)

	total := 0

	max := arr[0]
	curIdx := 0
	space := make([]int, 0)
	for idx, _ := range arr {
		total++
		if arr[idx] > max {
			max = arr[idx]
		}

		if idx == 0 {
			space = append(space, arr[idx])
			continue
		}

		if arr[idx] > 0 && space[curIdx] > 0 || arr[idx] < 0 && space[curIdx] < 0 {
			space[curIdx] += arr[idx]
		} else if arr[idx] == 0 {
			continue
		} else {
			curIdx++
			space = append(space, arr[idx])
		}
	}
	fmt.Println(space)

	for idx, _ := range space {
		if space[idx] > max {
			max = space[idx]
		}

		total++
		if idx == 0 || idx+1 == len(space) || space[idx] > 0 {
			continue
		}

		if !(space[idx]+space[idx-1] > 0 && space[idx]+space[idx+1] > 0) {
			continue
		}

		curIdx = idx - 1
		space[curIdx] += space[idx] + space[idx+1]
		total--
		for {
			total++
			if curIdx < 2 || !(space[curIdx-1]+space[curIdx] > 0 && space[curIdx-1]+space[curIdx-2] > 0) {
				break
			}

			space[curIdx-2] += space[curIdx-1] + space[curIdx]
			curIdx -= 2
		}
	}
	fmt.Println(space[:curIdx+1])

	for _, val := range space[:curIdx+1] {
		total++
		if val > max {
			max = val
		}
	}

	fmt.Println(len(arr), total)
	return max
}
