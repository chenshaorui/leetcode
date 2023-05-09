package Solution_1

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	var stack []byte

	for _, char := range s {
		if _, ok := pairs[byte(char)]; ok {
			stack = append(stack, byte(char))
		} else if len(stack) > 0 && byte(char) == pairs[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
