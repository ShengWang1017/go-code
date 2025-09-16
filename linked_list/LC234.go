package linked_list

func IsPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	halfHead := slow.Next
	slow.Next = nil
	reverseHead := ReverseList(halfHead)
	for reverseHead != nil {
		if reverseHead.Val != head.Val {
			return false
		}
		reverseHead = reverseHead.Next
		head = head.Next
	}
	return true
}
