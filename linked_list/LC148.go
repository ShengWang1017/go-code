package linked_list

// SortList LC148 排序链表
// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	for {
		swapped := false
		prev := dummy
		cur := dummy.Next

		for cur != nil && cur.Next != nil {
			if cur.Val > cur.Next.Val {
				// 相邻交换：prev -> cur -> nxt  变成  prev -> nxt -> cur
				nxt := cur.Next
				cur.Next = nxt.Next
				nxt.Next = cur
				prev.Next = nxt

				swapped = true
				prev = nxt // 交换后 prev 指向新左节点
				// cur 保持不变（现在是右节点），继续和它新的后继比较
			} else {
				// 不交换，整体前进一格
				prev = cur
				cur = cur.Next
			}
		}

		if !swapped { // 本轮无交换，已排序
			break
		}
	}
	return dummy.Next
}
