package binarytree

import "github.com/just-mitch/leetcode-go/interfaces"

type DFSPreOrderNodeIterator[T any] struct {
	stack []*Node[T]
}

func (it *DFSPreOrderNodeIterator[T]) HasNext() bool {
	return len(it.stack) > 0
}

func (it *DFSPreOrderNodeIterator[T]) Next() *Node[T] {

	node := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]

	if node.Right != nil {
		it.stack = append(it.stack, node.Right)
	}
	if node.Left != nil {
		it.stack = append(it.stack, node.Left)
	}

	return node

}

func NewDFSPreOrderNodeIterator[T any](t *BinaryTree[T]) interfaces.Iterator[Node[T]] {
	it := &DFSPreOrderNodeIterator[T]{
		stack: []*Node[T]{},
	}
	if t != nil && t.Root != nil {
		it.stack = append(it.stack, t.Root)
	}
	return it
}

func (t *BinaryTree[T]) ToPreOrderSlice() []T {
	a := make([]T, 0, 10)
	it := NewDFSPreOrderNodeIterator[T](t)
	for it.HasNext() {
		a = append(a, it.Next().Val)
	}
	return a
}
