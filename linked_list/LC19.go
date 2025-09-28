package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head

	count := 0
	sum := 1
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		count++
	}
	if fast == nil {
		sum = count*2 - 1 + sum
	} else if fast.Next == nil {
		sum = count*2 + sum
	}

	cur := &ListNode{}
	cur.Next = head
	for i := 0; i < sum; i++ {
		if i == sum-n {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return head

}
