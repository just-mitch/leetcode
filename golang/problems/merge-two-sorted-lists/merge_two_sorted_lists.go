package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var head = &ListNode{}

	curr, i, j := head, list1, list2

	for i != nil && j != nil {
		var tmp *ListNode
		if i.Val <= j.Val {
			tmp = i
			i = i.Next
			curr.Next = tmp
		} else {
			tmp = j
			j = j.Next
			curr.Next = tmp
		}
		curr = curr.Next
	}
	if i != nil {
		curr.Next = i
	} else if j != nil {
		curr.Next = j
	}

	return head.Next

}
