package util

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

type CircularDoubleLinkedList[T any] struct {
	head *Node[T]
	len  int
}

func NewCircularLinkedList[T any]() *CircularDoubleLinkedList[T] {
	return &CircularDoubleLinkedList[T]{}
}

func (c *CircularDoubleLinkedList[T]) Add(val T) {
	el := &Node[T]{val: val}

	if c.head == nil {
		c.head, el.next, el.prev = el, el, el
		c.len = 1
		return
	}

	c.head.prev.next, el.prev, el.next, c.head.prev = el, c.head.prev, c.head, el
	c.len++
}

func (c *CircularDoubleLinkedList[T]) AddAfterNode(val T, node *Node[T]) {
	el := &Node[T]{val: val}
	el.prev, el.next = node, node.next
	node.next.prev, node.next = el, el
	c.len++
}

func (c *CircularDoubleLinkedList[T]) Reverse(from, to int) {
	n := to - from + 1
	if to < from {
		n = (c.len - from + 1) + to
	}

	for i := 0; i < n/2; i++ {
		ff := (from + i) % c.len
		tt := (to - i)
		if tt < 0 {
			tt = c.len + tt
		}
		lhs, rhs := c.GetNode(ff), c.GetNode(tt)
		lhs.val, rhs.val = rhs.val, lhs.val

	}
}

func (c *CircularDoubleLinkedList[T]) Get(n int) T {
	return c.GetNode(n).val
}

func (c *CircularDoubleLinkedList[T]) GetNode(n int) *Node[T] {
	var ret = c.head
	for i := 0; i < n; i++ {
		ret = ret.next
	}
	return ret
}

func (c *CircularDoubleLinkedList[T]) GetNodeFunc(cmp func(v T) bool) (*Node[T], bool) {
	if cmp(c.head.val) {
		return c.head, true
	}
	n := c.head.next

	for n != c.head {
		if cmp(n.val) {
			return n, true
		}
	}

	return nil, false

}

func (c *CircularDoubleLinkedList[T]) NodeValue(n *Node[T]) T {
	return n.val
}

func (c *CircularDoubleLinkedList[T]) StepFromNode(n *Node[T], steps int) *Node[T] {
	for i := 0; i < steps; i++ {
		n = n.next
	}
	return n
}

func (c *CircularDoubleLinkedList[T]) String() string {
	var sb strings.Builder

	sep := ""
	c.Each(func(t T) {
		sb.WriteString(fmt.Sprintf("%s%v", sep, t))
		sep = " -> "
	})
	return sb.String()
}

func (c *CircularDoubleLinkedList[T]) Each(f func(T)) {
	last, curr := c.head.prev, c.head
	visited := map[*Node[T]]struct{}{}
	for {
		if _, ok := visited[curr]; ok {
			panic(fmt.Sprintf("already visited %v", curr.val))
		}
		f(curr.val)
		if last == curr {
			break
		}
		visited[curr] = struct{}{}
		curr = curr.next
	}
}

func (c *CircularDoubleLinkedList[T]) Len() int {
	n := 0
	c.Each(func(t T) { n++ })
	return n
}
