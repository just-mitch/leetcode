// a simple, singly linked list implementation
package linkedlist

import (
	"fmt"

	"github.com/just-mitch/leetcode-go/interfaces"
)

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.Val)
}

type LinkedList[T comparable] struct {
	Head *Node[T]
}

type NodeIterator[T comparable] struct {
	Current *Node[T]
}

func (i *NodeIterator[T]) HasNext() bool {
	return i.Current != nil
}

func (i *NodeIterator[T]) Next() *Node[T] {
	current := i.Current
	i.Current = i.Current.Next
	return current
}

func (n *Node[T]) Iterator() interfaces.Iterator[Node[T]] {
	return &NodeIterator[T]{n}
}

func (l *LinkedList[T]) Iterator() interfaces.Iterator[Node[T]] {
	return l.Head.Iterator()
}

func (l *LinkedList[T]) String() string {
	s := ""
	iterator := l.Iterator()
	for iterator.HasNext() {
		s += fmt.Sprintf("%v ", iterator.Next())
	}
	return s
}

func (l *LinkedList[T]) ToString() string {
	return l.String()
}
