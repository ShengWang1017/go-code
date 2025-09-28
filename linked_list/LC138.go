package linked_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := map[*Node]int{}
	cur := head
	var nodes []*Node

	for i := 0; cur != nil; i++ {
		nodeMap[cur] = i
		newNode := &Node{Val: cur.Val}
		nodes = append(nodes, newNode)
		if i != 0 {
			nodes[i-1].Next = nodes[i]
		}
		cur = cur.Next
	}
	cur = head
	for i := 0; i < len(nodeMap); i++ {
		if cur.Random == nil {
			nodes[i].Random = nil
		} else {
			index := nodeMap[cur.Random]
			nodes[i].Random = nodes[index]
		}
		cur = cur.Next
	}

	return nodes[0]

}
