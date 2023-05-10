package Solution_1

func longestPalindrome(s string) string {
	strLen := len(s)

	palindrome := s[:1]

	if strLen > 1 {
		isPalindromeMatrix := make([][]bool, strLen)

		isPalindromeMatrix[0] = make([]bool, strLen)
		for i := 0; i < strLen; i++ {
			isPalindromeMatrix[0][i] = true
		}

		isPalindromeMatrix[1] = make([]bool, strLen-1)
		for i := 0; i < strLen-1; i++ {
			if s[i] == s[i+1] {
				isPalindromeMatrix[1][i] = true
				palindrome = s[i : i+2]
			}
		}

		for i := 2; i < strLen; i++ {
			isPalindromeMatrix[i] = make([]bool, strLen-i)
			for j := 0; j < strLen-i; j++ {
				if s[j] == s[j+i] && isPalindromeMatrix[i-2][j+1] {
					isPalindromeMatrix[i][j] = true
					palindrome = s[j : j+i+1]
				}
			}
		}
	}

	return palindrome
}
