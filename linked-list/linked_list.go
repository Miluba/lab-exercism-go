package linkedlist

import (
	"errors"
)

// Define List and Node types here.
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

var ErrEmptyList = errors.New("Empty list")

type List struct {
	root *Node
	size int
}

func NewList(args ...interface{}) *List {
	ll := &List{}
	ll.size = len(args)
	if ll.size == 0 {
		return ll
	}
	for _, v := range args {
		ll.PushBack(v)
	}
	return ll
}

func (l *List) PushFront(v interface{}) {
	l.size++
	if l.root == nil {
		l.root = &Node{
			Val: v,
		}
		return
	}
	n := &Node{
		Val:  v,
		next: l.root,
	}
	l.root.prev = n
	l.root = n
}

func (l *List) PushBack(v interface{}) {
	t := l.Last()
	l.size++
	if t == nil {
		l.root = &Node{
			Val: v,
		}
		return
	}
	n := &Node{
		Val:  v,
		prev: t,
	}
	t.next = n
}

func (l *List) PopFront() (interface{}, error) {
	if l.root == nil {
		return nil, ErrEmptyList
	}
	l.size--
	v := l.root.Val
	l.root = l.root.next
	if l.root != nil {
		l.root.prev = nil
	}
	return v, nil
}

func (l *List) PopBack() (interface{}, error) {
	last := l.Last()
	head := l.root
	if last == nil && head == nil {
		return nil, ErrEmptyList
	}
	l.size--
	if last != nil && head != nil && last.prev == nil {
		l.root.next = nil
		l.root = nil
		return last.Val, nil
	}
	if last != nil && head != nil && last.prev != nil {
		v := last.Val
		last.prev.next = nil
		last = last.prev
		return v, nil
	}
	panic("bad impl")
}

func (l *List) Reverse() {
	if l.size < 2 {
		return
	}
	var tmp *Node
	c := l.root
	for c != nil {
		tmp = c.prev
		c.prev = c.next
		c.next = tmp
		c = c.prev
	}
	if tmp != nil {
		l.root = tmp.prev
	}
}

func (l *List) First() *Node {
	return l.root
}

func (l *List) Last() *Node {
	n := l.root
	if n == nil {
		return nil
	}
	for {
		if n.next == nil {
			return n
		}
		n = n.next
	}

}
