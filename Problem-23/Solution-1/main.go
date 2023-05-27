package Solution_1

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

func NewPriorityQueue() PriorityQueue {
	return make([]*ListNode, 0)
}

func (priorityQueue *PriorityQueue) Len() int {
	return len(*priorityQueue)
}

func (priorityQueue *PriorityQueue) Less(i, j int) bool {
	return (*priorityQueue)[i].Val < (*priorityQueue)[j].Val
}

func (priorityQueue *PriorityQueue) Swap(i, j int) {
	(*priorityQueue)[i], (*priorityQueue)[j] = (*priorityQueue)[j], (*priorityQueue)[i]
}

func (priorityQueue *PriorityQueue) Push(x interface{}) {
	listNode := x.(*ListNode)
	*priorityQueue = append(*priorityQueue, listNode)
}

func (priorityQueue *PriorityQueue) Pop() interface{} {
	length := len(*priorityQueue)

	listNode := (*priorityQueue)[length-1]
	*priorityQueue = (*priorityQueue)[:length-1]

	return listNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	priorityQueue := NewPriorityQueue()
	heap.Init(&priorityQueue)

	for _, head := range lists {
		if head != nil {
			heap.Push(&priorityQueue, head)
		}
	}

	var head, curr *ListNode

	for priorityQueue.Len() > 0 {
		next := heap.Pop(&priorityQueue).(*ListNode)

		if curr == nil {
			head = next
		} else {
			curr.Next = next
		}
		curr = next

		if next.Next != nil {
			heap.Push(&priorityQueue, next.Next)
		}
	}

	if curr != nil {
		curr.Next = nil
	}

	return head
}
