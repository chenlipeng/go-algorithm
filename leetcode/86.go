package leetcode

/**
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var lowHead, lowTail *ListNode
	var highHead, highTail *ListNode

	for nil != head {
		if head.Val < x {
			if lowHead == nil {
				lowHead = head
				lowTail = head
			} else {
				lowTail.Next = head
				lowTail = head
			}
			head = head.Next
			lowTail.Next = nil
		} else {
			if highHead == nil {
				highHead = head
				highTail = head
			} else {
				highTail.Next = head
				highTail = head
			}
			head = head.Next
			highTail.Next = nil
		}
	}

	if lowTail != nil {
		lowTail.Next = highHead
		return lowHead
	}
	return highHead
}
