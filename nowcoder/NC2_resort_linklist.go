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
 * @return void
 */
func reorderList(head *ListNode) {
	// write code here

	// 01. 假设链表长度为X，将一个链表拆为两个链表，第一个链表为[0, X/2+1], 利用快慢指针找出第二个链表的头节点
	// 02. 将第二个链表逆置
	// 03. 将链表1、2交替链接起来

}
