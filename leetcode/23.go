package leetcode

/*
 * 23. 合并K个升序链表
 * 给你一个链表数组，每个链表都已经按升序排列。
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 */

func MergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	var head, tail, node *ListNode
	for nil != left && nil != right {
		if left.Val <= right.Val {
			node = left
			left = left.Next
		} else {
			node = right
			right = right.Next
		}

		if nil == head {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = tail.Next
		}
		tail.Next = nil
	}

	if nil != left {
		node = left
	}
	if nil != right {
		node = right
	}

	if nil == head {
		head = node
	} else {
		tail.Next = node
	}
	return head
}
