package linked_list

func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 阶段一：检测是否有环，并找到相遇点
	// 两个指针必须都从 head 开始
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// 如果相遇，说明有环
		if slow == fast {
			break
		}
	}

	// 如果 fast 走到了尽头，说明没有环
	if fast == nil || fast.Next == nil {
		return nil
	}

	for slow2 := head; slow2 != slow; {
		slow = slow.Next
		slow2 = slow2.Next

	}

	return slow

}
