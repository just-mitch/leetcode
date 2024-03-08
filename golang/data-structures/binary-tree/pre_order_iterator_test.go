package binarytree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPreOrderIterator(t *testing.T) {

	type args struct {
		l *BinaryTree[int]
	}
	tests := []struct {
		args args
		want []int
	}{
		{args{
			&BinaryTree[int]{},
		}, []int{}},
		{args{
			&BinaryTree[int]{
				&Node[int]{2, &Node[int]{0, nil,
					&Node[int]{1, nil, nil},
				}, &Node[int]{4,
					&Node[int]{3, nil, nil},
					&Node[int]{5, nil, nil}}}},
		},
			[]int{2, 0, 1, 4, 3, 5}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("to-preorder-slice-%v", i), func(t *testing.T) {
			if got := tt.args.l.ToPreOrderSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%T, ToPreOrderSlice() = %v, want %v", tt.args.l, got, tt.want)
			}
		})
	}
}
