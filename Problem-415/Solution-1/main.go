package Solution_1

import "strconv"

func addStrings(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)

	sum := ""
	carry := uint8(0)

	i, j := len1-1, len2-1
	for i >= 0 || j >= 0 {
		var digit1, digit2 uint8

		if i < 0 {
			digit1 = 0
		} else {
			digit1 = num1[i] - '0'
			i--
		}

		if j < 0 {
			digit2 = 0
		} else {
			digit2 = num2[j] - '0'
			j--
		}

		digitSum := digit1 + digit2 + carry
		carry = digitSum / 10

		sum = strconv.Itoa(int(digitSum%10)) + sum
	}

	if carry == 1 {
		sum = "1" + sum
	}

	return sum
}
