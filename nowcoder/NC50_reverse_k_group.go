package nowcoder

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	if nil == head {
		return head
	}

	ktail := head
	for idx := 0; idx < k; idx++ {
		if nil == ktail.Next {
			return head
		}

		ktail = ktail.Next
	}

	stop := false
	for nil != head.Next && false == stop {
		for idx := 0; idx < k; idx++ {
		}
	}
}
