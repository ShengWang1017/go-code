package main

import (
	"fmt"
	"go-code/linked_list"
)

func main() {
	listNode1 := linked_list.ArrayConvToListNode([]int{4, 1, 8, 4, 5})
	listNode2 := linked_list.ArrayConvToListNode([]int{5, 6, 1, 8, 4, 5})

	result := linked_list.GetIntersectionNode(listNode1, listNode2)
	fmt.Printf("result : %v", result)

}
