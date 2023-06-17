package Solution_1

func minDistance(word1 string, word2 string) int {
	word1Len, word2Len := len(word1), len(word2)

	distanceMatrix := make([][]int, word1Len+1)
	for i := 0; i <= word1Len; i++ {
		distanceMatrix[i] = make([]int, word2Len+1)
		distanceMatrix[i][0] = i
	}

	for i := 1; i <= word2Len; i++ {
		distanceMatrix[0][i] = i
	}

	for i := 1; i <= word1Len; i++ {
		for j := 1; j <= word2Len; j++ {
			k, l := i-1, j-1

			distance := distanceMatrix[k][l]
			if word1[k] != word2[l] {
				distance += 1
			}

			distanceMatrix[i][j] = min(distanceMatrix[k][j]+1, distanceMatrix[i][l]+1, distance)
		}
	}

	return distanceMatrix[word1Len][word2Len]
}

func min(num1, num2, num3 int) int {
	minNum := num1

	if num2 < minNum {
		minNum = num2
	}

	if num3 < minNum {
		minNum = num3
	}

	return minNum
}
