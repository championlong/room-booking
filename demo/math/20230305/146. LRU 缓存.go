package main

type LRUCache struct {
	capacity   int
	nodeMap    map[int]*Node
	head, tail *Node
}
type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		capacity: capacity,
		nodeMap:  make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) DeleteNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
func (this *LRUCache) AddHead(node *Node) {
	this.head.Next.Prev = node
	node.Next = this.head.Next
	this.head.Next = node
	node.Prev = this.head
}
func (this *LRUCache) Get(key int) int {
	if v, ok := this.nodeMap[key]; ok {
		this.DeleteNode(v)
		this.AddHead(v)
		return v.Value
	}
	return -1
}
func (this *LRUCache) Removetail() {
	node := this.tail.Prev
	this.DeleteNode(node)
	delete(this.nodeMap, node.Key)
}
func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.nodeMap[key]; ok {
		v.Value = value
		this.DeleteNode(v)
		this.AddHead(v)
		return
	} else {
		if this.capacity == len(this.nodeMap) {
			this.Removetail()
		}
		node := &Node{Key: key, Value: value}
		this.AddHead(node)
		this.nodeMap[key] = node
	}
}
