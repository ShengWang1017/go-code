package linked_list

// ReverseKGroup LC25K个一组翻转链表
// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 你不能只是单纯地改变节点内部的值，而是需要实际进行节点交换。
func ReverseKGroup(head *ListNode, k int) *ListNode {
	var pre *ListNode
	cur := head
	tails := make([]*ListNode, 0)
	heads := make([]*ListNode, 0)

	for cur != nil {
		tails = append(tails, cur)
		var temp *ListNode
		for i := 0; i < k; i++ {
			temp = cur.Next
			cur.Next = pre
			pre = cur
			cur = temp
			if cur == nil {
				break
			}
		}
		heads = append(heads, pre)
	}

	for i := range heads {
		if i == len(heads)-1 {
			tails[i].Next = nil
			break
		}
		tails[i].Next = heads[i+1]
	}
	return heads[0]
}
