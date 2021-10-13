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
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
func threeOrders(root *TreeNode) [][]int {
	// write code here
	list := [][]int{}
	list = append(list, preOrder(root))
	list = append(list, midOrder(root))
	list = append(list, postOrder(root))
	return list
}

func preOrder_1(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	list := []int{}
	stack := []*TreeNode{root}
	list = append(list, root.Val)

	for len(stack) > 0 {
		current := stack[len(stack)-1].Left
		for nil != current {
			stack = append(stack, current)
			list = append(list, current.Val)
			current = current.Left
		}

		for len(stack) > 0 {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if nil != current.Right {
				stack = append(stack, current.Right)
				list = append(list, current.Right.Val)
				current = current.Right
				break
			}
		}
	}
	return list
}

func midOrder_1(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	list := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1].Left
		for nil != current {
			stack = append(stack, current)
			current = current.Left
		}

		for len(stack) > 0 {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			list = append(list, current.Val)
			if nil != current.Right {
				stack = append(stack, current.Right)
				current = current.Right
				break
			}
		}
	}
	return list
}

func postOrder_1(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	list := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		current := stack[len(stack)-1].Left
		for nil != current {
			stack = append(stack, current)
			current = current.Left
		}

		tempList = []int{}
		for len(stack) > 0 {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if nil != current.Right {
				stack = append(stack, current.Right)
				current = current.Right
				break
			}
		}
		list = append(list, current.Val)
	}
	return list
}

//递归实现先

func preOrder(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	list := []int{}
	list = append(list, root.Val)
	list = append(list, preOrder(root.Left)...)
	list = append(list, preOrder(root.Right)...)
	return list
}

func midOrder(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	list := []int{}
	list = append(list, midOrder(root.Left)...)
	list = append(list, root.Val)
	list = append(list, midOrder(root.Right)...)
	return list
}

func postOrder(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}

	list := []int{}
	list = append(list, postOrder(root.Left)...)
	list = append(list, postOrder(root.Right)...)
	list = append(list, root.Val)
	return list
}
