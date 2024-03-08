package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func getContents(l *ListNode) (int, *ListNode) {
	if l == nil {
		return 0, nil
	} else {
		return l.Val, l.Next
	}
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {

	var head *ListNode

	sum := &ListNode{0, nil}

	head = sum

	for {
		l1Val, l1Next := getContents(l1)
		l2Val, l2Next := getContents(l2)

		carry := 0
		if iSum := sum.Val + l1Val + l2Val; iSum >= 10 {
			sum.Val = iSum - 10
			carry = 1
		} else {
			sum.Val = iSum
		}

		l1 = l1Next
		l2 = l2Next

		if carry == 1 {
			sum.Next = &ListNode{
				1,
				nil,
			}
		} else if l1 != nil || l2 != nil {
			sum.Next = &ListNode{
				0,
				nil,
			}
		} else {
			break
		}

		sum = sum.Next

	}

	return head

}
