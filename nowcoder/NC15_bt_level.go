package nowcoder

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 */
func levelOrder(root *TreeNode) [][]int {
	// write code here
	if nil == root {
		return [][]int{}
	}

	ret := [][]int{}
	nodes := []*TreeNode{root}
	start := 0
	for len(nodes[start:]) > 0 {
		list := []int{}
		for _, node := range nodes[start:] {
			list = append(list, node.Val)

			if nil != node.Left {
				nodes = append(nodes, node.Left)
			}
			if nil != node.Right {
				nodes = append(nodes, node.Right)
			}
			start++
		}

		ret = append(ret, list)
	}

	return ret
}
