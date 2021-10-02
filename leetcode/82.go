package leetcode

func DeleteDuplicates(head *ListNode) *ListNode {
	return deleteDuplicates(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var newHead, newTail, tail *ListNode

	tail = head
	cnt := 0

	for tail != nil {
		tail = tail.Next
		cnt++
		if tail != nil && head.Val == tail.Val {
			continue
		}

		if cnt == 1 {
			//保留
			if newHead == nil {
				newHead = head
				newTail = newHead
			} else {
				newTail.Next = head
				newTail = newTail.Next
			}
			head = tail
			newTail.Next = nil
		} else {
			//跳过
			head = tail
		}
		cnt = 0
	}

	return newHead
}
