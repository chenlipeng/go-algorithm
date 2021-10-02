package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	var head, tail *ListNode
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		node := &ListNode{Val: (v1 + v2 + carry) % 10}
		carry = (v1 + v2 + carry) / 10

		if nil == head {
			head = node
			tail = head
		} else {
			tail.Next = node
			tail = tail.Next
		}
	}

	if carry > 0 {
		node := &ListNode{Val: carry}
		tail.Next = node
	}

	return head
}

func reverseLinkList(l *ListNode) *ListNode {
	var list *ListNode
	for l != nil {
		tmp := l
		l = l.Next
		tmp.Next = list
		list = tmp
	}
	return list
}
