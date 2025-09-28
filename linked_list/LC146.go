package linked_list

type BidNode struct {
	Value      int
	prev, next *BidNode
}
type LRUCache struct {
	Cap     int
	Length  int
	nodeMap map[int]*BidNode
	Head    *BidNode
	Tail    *BidNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{Cap: capacity,
		nodeMap: make(map[int]*BidNode),
		Length:  0,
		Head:    &BidNode{},
		Tail:    &BidNode{},
	}
	cache.Head.next = cache.Tail
	cache.Tail.prev = cache.Head
	return cache
}

func (this *LRUCache) Get(key int) int {
	return -1
}

func (this *LRUCache) Put(key int, value int) {

}
