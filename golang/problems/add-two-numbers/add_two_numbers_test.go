package addtwonumbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{
		2,
		&ListNode{4, &ListNode{3, nil}},
	}

	l2 := ListNode{
		5,
		&ListNode{6, &ListNode{4, &ListNode{1, nil}}},
	}

	want := &ListNode{
		7,
		&ListNode{0, &ListNode{8, &ListNode{1, nil}}},
	}

	out := addTwoNumbers(&l1, &l2)

	for {
		outV := out.Val
		wantV := want.Val

		if outV != wantV {
			t.Fatalf(`Out = %v, Want = %v`, outV, wantV)
		}

		out = out.Next
		want = want.Next

		if out == nil || want == nil {
			break
		}
	}

	if !(out == nil && want == nil) {
		t.Fatalf(`Mismatched lengths. %v, %v`, out, want)
	}

}
