package Solution_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nodes1, nodes2 []*ListNode

	for l1 != nil {
		nodes1 = append(nodes1, l1)
		l1 = l1.Next
	}

	for l2 != nil {
		nodes2 = append(nodes2, l2)
		l2 = l2.Next
	}

	var head *ListNode
	carry := 0

	i, j := len(nodes1)-1, len(nodes2)-1
	for i >= 0 || j >= 0 {
		var digit1, digit2 int

		if i < 0 {
			digit1 = 0
		} else {
			digit1 = nodes1[i].Val
			i--
		}

		if j < 0 {
			digit2 = 0
		} else {
			digit2 = nodes2[j].Val
			j--
		}

		digitSum := digit1 + digit2 + carry
		carry = digitSum / 10

		head = &ListNode{
			Val:  digitSum % 10,
			Next: head,
		}
	}

	if carry == 1 {
		head = &ListNode{
			Val:  1,
			Next: head,
		}
	}

	return head
}
