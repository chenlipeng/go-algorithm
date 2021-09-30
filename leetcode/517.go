package leetcode

/*
 * 假设有 n 台超级洗衣机放在同一排上。开始的时候，每台洗衣机内可能有一定量的衣服，也可能是空的。
 * 在每一步操作中，你可以选择任意 m (1 <= m <= n) 台洗衣机，与此同时将每台洗衣机的一件衣服送到相邻的一台洗衣机。
 * 给定一个整数数组 machines 代表从左至右每台洗衣机中的衣物数量，请给出能让所有洗衣机中剩下的衣物的数量相等的 最少的操作步数。如果不能使每台洗衣机中衣物的数量相等，则返回 -1。
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/super-washing-machines
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/*
 * 示例 1：
 * 	输入：machines = [1,0,5]
 * 	输出：3
 * 	解释：
 * 	第一步:    1     0 <-- 5    =>    1     1     4
 * 	第二步:    1 <-- 1 <-- 4    =>    2     1     3
 * 	第三步:    2     1 <-- 3    =>    2     2     2
 *
 * 示例 2：
 * 	输入：machines = [0,3,0]
 * 	输出：2
 * 	解释：
 * 	第一步:    0 <-- 3     0    =>    1     2     0
 * 	第二步:    1     2 --> 0    =>    1     1     1
 *
 * 示例 3：
 * 	输入：machines = [0,2,0]
 * 	输出：-1
 * 	解释：
 * 	不可能让所有三个洗衣机同时剩下相同数量的衣物。
 */

/*
 * A B C ———— 三台洗衣机
 * 操作开始前，洗衣机B里没衣服，就算从相邻的洗衣机C里获得了衣服 也不能送到A
 * 操作中，每台洗衣机只能有一次送出
 */
func FindMinMoves(machines []int) int {
	return findMinMoves(machines)
}

func findMinMoves(machines []int) int {
	if len(machines) <= 0 {
		return -1
	}

	//01. 判断是否可以让每台洗衣机中衣物的数量相等 所有衣物数 / 洗衣机数量
	clothesCnt := 0
	for _, cnt := range machines {
		clothesCnt += cnt
	}

	//02. 最终每台洗衣机中的衣服数量
	avgCnt := clothesCnt / len(machines)
	if avgCnt*len(machines) != clothesCnt {
		return -1
	}

	//03.
	totalSteps := 0
	//for _, cnt := range machines {
	//	if cnt >= avgCnt {
	//		continue
	//	}
	//	totalSteps += avgCnt - cnt
	//}

	return totalSteps
}

func findMinMoves_1(machines []int) (ans int) {
	tot := 0
	for _, v := range machines {
		tot += v
	}
	n := len(machines)
	if tot%n > 0 {
		return -1
	}
	avg := tot / n
	sum := 0
	for _, num := range machines {
		num -= avg
		sum += num
		ans = max(ans, max(abs(sum), num))
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
