package linked_list

// AddTwoNumbers LC2 两数相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := 0
	next := 0

	node1 := l1
	node2 := l2

	node3 := &ListNode{}
	dummy := node3
	for !(node1 == nil && node2 == nil) {
		var sum int
		if node1 == nil {
			sum = node2.Val
			node2 = node2.Next
		} else if node2 == nil {
			sum = node1.Val
			node1 = node1.Next
		} else {
			sum = node1.Val + node2.Val
			node1 = node1.Next
			node2 = node2.Next
		}
		cur = (sum + next) % 10
		next = (sum + next) / 10
		node3.Next = &ListNode{Val: cur}
		node3 = node3.Next
	}
	if next > 0 {
		node3.Next = &ListNode{Val: next}
	}

	return dummy.Next
}
