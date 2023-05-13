package Solution_1

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
	isVisitedMatrix[i][j] = true

	if grid[i][j] != '1' {
		return
	}

	above := i - 1
	if above >= 0 && !isVisitedMatrix[above][j] {
		search(grid, isVisitedMatrix, rowNum, colNum, above, j)
	}

	below := i + 1
	if below < rowNum && !isVisitedMatrix[below][j] {
		search(grid, isVisitedMatrix, rowNum, colNum, below, j)
	}

	left := j - 1
	if left >= 0 && !isVisitedMatrix[i][left] {
		search(grid, isVisitedMatrix, rowNum, colNum, i, left)
	}

	right := j + 1
	if right < colNum && !isVisitedMatrix[i][right] {
		search(grid, isVisitedMatrix, rowNum, colNum, i, right)
	}
}
