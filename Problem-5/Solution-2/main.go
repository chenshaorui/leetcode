package Solution_2

func longestPalindrome(s string) string {
	strLen := len(s)

	maxLen := 0
	palindrome := ""

	for i := 0; i < strLen; i++ {
		j := 1
		for j <= i && i+j < strLen && s[i-j] == s[i+j] {
			j++
		}
		j--

		length := j*2 + 1
		if length > maxLen {
			palindrome = s[i-j : i+j+1]
			maxLen = length
		}

		if maxLen >= (strLen-i)*2-2 || maxLen >= i*2+2 {
			break
		}

		next := i + 1
		if next < strLen && s[i] == s[next] {
			j = 1
			for j <= i && next < strLen-1 && s[i-j] == s[next+j] {
				j++
			}
			j--

			length = j*2 + 2
			if length > maxLen {
				palindrome = s[i-j : next+j+1]
				maxLen = length
			}
		}

		if maxLen >= (strLen-i)*2-3 || maxLen >= i*2+3 {
			break
		}
	}

	return palindrome
}
