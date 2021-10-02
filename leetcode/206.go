package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var nHead *ListNode

	for head != nil {
		tmp := head
		head = head.Next

		tmp.Next = nHead
		nHead = tmp
	}

	return nHead
}
