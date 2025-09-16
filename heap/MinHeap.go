package heap

import (
	"container/heap"
	"go-code/linked_list"
)

type MinHeap []*linked_list.ListNode

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*linked_list.ListNode))
}

func (h *MinHeap) Pop() any {
	n := len(*h)
	last := (*h)[n-1]
	*h = (*h)[:n-1]
	return last
}

// mergeKLists	LC23 合并K个升序链表
// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
func mergeKLists(lists []*linked_list.ListNode) *linked_list.ListNode {
	head := &linked_list.ListNode{}
	cur := head

	h := &MinHeap{}
	heap.Init(h)
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	for h.Len() > 0 {
		node := heap.Pop(h).(*linked_list.ListNode)
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return head.Next

}
