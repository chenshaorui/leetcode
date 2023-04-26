package Solution_1

type LinkedListNode struct {
	key   int
	value int
	prev  *LinkedListNode
	next  *LinkedListNode
}

func NewEmptyLinkedListNode() *LinkedListNode {
	return &LinkedListNode{
		key:   0,
		value: 0,
		prev:  nil,
		next:  nil,
	}
}

func NewLinkedListNode(key int, value int) *LinkedListNode {
	return &LinkedListNode{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

type LRUCache struct {
	nodeMap  map[int]*LinkedListNode
	head     *LinkedListNode
	tail     *LinkedListNode
	size     int
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := NewEmptyLinkedListNode()
	tail := NewEmptyLinkedListNode()

	head.next = tail
	tail.prev = head

	return LRUCache{
		nodeMap:  map[int]*LinkedListNode{},
		head:     head,
		tail:     tail,
		size:     0,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodeMap[key]; ok {
		this.moveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.nodeMap[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node = NewLinkedListNode(key, value)
		this.nodeMap[key] = node
		this.addToHead(node)
		this.size++

		if this.size > this.capacity {
			tail := this.removeTail()
			delete(this.nodeMap, tail.key)
			this.size--
		}
	}
}

func (this *LRUCache) addToHead(node *LinkedListNode) {
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
}

func (this *LRUCache) removeNode(node *LinkedListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *LinkedListNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *LinkedListNode {
	tail := this.tail.prev
	this.removeNode(tail)
	return tail
}
