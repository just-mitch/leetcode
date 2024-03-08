package binarytree

import (
	"fmt"

	"github.com/just-mitch/leetcode-go/interfaces"
)

type Node[T any] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.Val)
}

type BinaryTree[T any] struct {
	Root *Node[T]
}

func (t *BinaryTree[T]) String() string {
	return ""
}

type DFSInOrderNodeIterator[T any] struct {
	current *Node[T]
	stack   []*Node[T]
}

func (it *DFSInOrderNodeIterator[T]) HasNext() bool {
	return it.current != nil || len(it.stack) > 0
}

func (it *DFSInOrderNodeIterator[T]) Next() *Node[T] {

	for it.current != nil {
		it.stack = append(it.stack, it.current)
		it.current = it.current.Left
	}

	it.current = it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]

	node := it.current
	it.current = it.current.Right

	return node
}

func (n *Node[T]) Iterator() interfaces.Iterator[Node[T]] {
	return &DFSInOrderNodeIterator[T]{n, make([]*Node[T], 0, 10)}
}

func (t *BinaryTree[T]) Iterator() interfaces.Iterator[Node[T]] {
	return t.Root.Iterator()
}

func (t *BinaryTree[T]) ToSlice() []T {
	a := make([]T, 0, 10)
	it := t.Iterator()
	for it.HasNext() {
		a = append(a, it.Next().Val)
	}
	return a
}
