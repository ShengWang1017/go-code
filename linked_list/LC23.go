package linked_list

func MergeKLists(lists []*ListNode) *ListNode {

	head := &ListNode{}
	cur := head

	for {

		var minNode *ListNode

		// 纵向比较k个链，找到最小的节点
		k := 0
		for i, node := range lists {
			if node == nil {
				continue
			}
			if minNode == nil || minNode.Val > node.Val {
				minNode = node
				k = i
			}
		}
		if minNode == nil {
			break
		}
		cur.Next = minNode
		cur = cur.Next
		lists[k] = lists[k].Next
	}

	return head.Next
}
