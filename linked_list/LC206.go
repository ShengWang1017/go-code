package linked_list

// ReverseList
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func ReverseList(head *ListNode) *ListNode {
	// 1. 循环
	/*	var pre *ListNode = nil
		cur := head

		for cur != nil {
			temp := cur.Next
			cur.Next = pre
			pre = cur
			cur = temp
		}
		return pre*/

	// 2. 递归
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	newHead := ReverseList(next)

	next.Next = head
	head.Next = nil
	return newHead
}
