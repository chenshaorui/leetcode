package Solution_2

type Queue struct {
	elements [][2]int
}

func NewQueue() *Queue {
	return &Queue{
		elements: [][2]int{},
	}
}

func (queue *Queue) Push(element [2]int) {
	queue.elements = append(queue.elements, element)
}

func (queue *Queue) Pop() [2]int {
	var element [2]int

	if len(queue.elements) > 0 {
		element = queue.elements[0]
		queue.elements = queue.elements[1:]
	}

	return element
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.elements) == 0
}

func numIslands(grid [][]byte) int {
	islandNum := 0

	rowNum := len(grid)
	if rowNum > 0 {
		colNum := len(grid[0])

		isVisitedMatrix := make([][]bool, rowNum)
		for i := 0; i < rowNum; i++ {
			isVisitedMatrix[i] = make([]bool, colNum)
		}

		for i := 0; i < rowNum; i++ {
			for j := 0; j < colNum; j++ {
				if grid[i][j] == '1' && !isVisitedMatrix[i][j] {
					islandNum++
					search(grid, isVisitedMatrix, rowNum, colNum, i, j)
				}
			}
		}
	}

	return islandNum
}

func search(grid [][]byte, isVisitedMatrix [][]bool, rowNum, colNum int, i, j int) {
	queue := NewQueue()

	queue.Push([2]int{i, j})

	for !queue.IsEmpty() {
		location := queue.Pop()

		isVisitedMatrix[location[0]][location[1]] = true

		above := location[0] - 1
		if above >= 0 && grid[above][location[1]] == '1' && !isVisitedMatrix[above][location[1]] {
			queue.Push([2]int{above, location[1]})
		}

		below := location[0] + 1
		if below < rowNum && grid[below][location[1]] == '1' && !isVisitedMatrix[below][location[1]] {
			queue.Push([2]int{below, location[1]})
		}

		left := location[1] - 1
		if left >= 0 && grid[location[0]][left] == '1' && !isVisitedMatrix[location[0]][left] {
			queue.Push([2]int{location[0], left})
		}

		right := location[1] + 1
		if right < colNum && grid[location[0]][right] == '1' && !isVisitedMatrix[location[0]][right] {
			queue.Push([2]int{location[0], right})
		}
	}
}
