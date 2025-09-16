package linked_list

// GetIntersectionNode LC160 相交链表
// 如果两个链表相交，返回相交的节点
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	// 1. 利用map存储一条链表
	/*	set := map[*ListNode]struct{}{}

		for headA != nil {
			set[headA] = struct{}{}
			headA = headA.Next
		}

		for headB != nil {
			if _, exist := set[headB]; exist {
				return headB
			}
			headB = headB.Next
		}
		return nil*/

	// 2. 遍历完调到另一个链表上
	curA := headA
	curB := headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}

		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}

	}
	return curB
}
