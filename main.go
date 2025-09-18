package main

import (
	"fmt"
	"go-code/linked_list"
)

func main() {
	// 创建测试链表
	list1 := linked_list.ArrayConvToListNode([]int{1, 4, 5})
	list2 := linked_list.ArrayConvToListNode([]int{1, 3, 4})
	list3 := linked_list.ArrayConvToListNode([]int{2, 6})

	// 构造链表数组
	lists := []*linked_list.ListNode{list1, list2, list3}

	// 调用MergeKLists方法
	result := linked_list.MergeKLists(lists)

	// 打印结果
	fmt.Println("Merged list:")
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}
