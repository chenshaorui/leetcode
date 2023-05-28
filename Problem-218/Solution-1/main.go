package Solution_1

import (
	"container/heap"
	"sort"
)

type Item struct {
	right  int
	height int
}

type PriorityQueue []*Item

func NewPriorityQueue() PriorityQueue {
	return make([]*Item, 0)
}

func (priorityQueue *PriorityQueue) Len() int {
	return len(*priorityQueue)
}

func (priorityQueue *PriorityQueue) Less(i, j int) bool {
	return (*priorityQueue)[i].height > (*priorityQueue)[j].height
}

func (priorityQueue *PriorityQueue) Swap(i, j int) {
	(*priorityQueue)[i], (*priorityQueue)[j] = (*priorityQueue)[j], (*priorityQueue)[i]
}

func (priorityQueue *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*priorityQueue = append(*priorityQueue, item)
}

func (priorityQueue *PriorityQueue) Pop() interface{} {
	length := len(*priorityQueue)

	item := (*priorityQueue)[length-1]
	*priorityQueue = (*priorityQueue)[:length-1]

	return item
}

func getSkyline(buildings [][]int) [][]int {
	buildingTotal := len(buildings)

	boundaries := make([]int, 0, buildingTotal*2)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}

	sort.Ints(boundaries)

	skyline := make([][]int, 0)

	priorityQueue := NewPriorityQueue()
	heap.Init(&priorityQueue)

	i := 0
	for _, boundary := range boundaries {
		for i < buildingTotal && buildings[i][0] <= boundary {
			heap.Push(&priorityQueue, &Item{
				right:  buildings[i][1],
				height: buildings[i][2],
			})
			i++
		}

		for len(priorityQueue) > 0 && priorityQueue[0].right <= boundary {
			heap.Pop(&priorityQueue)
		}

		maxHeight := 0
		if len(priorityQueue) > 0 {
			maxHeight = priorityQueue[0].height
		}

		if len(skyline) == 0 || maxHeight != skyline[len(skyline)-1][1] {
			skyline = append(skyline, []int{boundary, maxHeight})
		}
	}

	return skyline
}
