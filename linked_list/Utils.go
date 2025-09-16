package linked_list

func ArrayConvToListNode(arr []int) *ListNode {
	dummy := &ListNode{}

	cur := dummy

	for _, val := range arr {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return dummy.Next
}
