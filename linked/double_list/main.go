package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data     int
	previous *Node
	next     *Node
}

type LinkedList struct {
	head *Node
}

func NewList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(data int) {
	new := &Node{data: data}
	if l.head == nil {
		l.head = new
		return
	}
	previous := l.head
	current := l.head.next
	for current != nil {
		previous = current
		current = current.next
	}
	current = new
	current.previous = previous
	previous.next = current
}

func (l *LinkedList) PrintForward() {
	if l.head == nil {
		fmt.Println("nil")
		return
	}
	s := strings.Builder{}
	current := l.head
	for current != nil {
		s.WriteString(fmt.Sprint(current.data))
		if current.next != nil {
			s.WriteString(" -> ")
		}
		current = current.next
	}
	fmt.Println(s.String())
}

func (l *LinkedList) PrintBackForward() {
	if l.head == nil {
		fmt.Println("nil")
	}
	current := l.head
	for current != nil {
		if current.next == nil {
			break
		}
		current = current.next
	}
	s := strings.Builder{}
	for current != nil {
		s.WriteString(fmt.Sprint(current.data))
		if current.previous != nil {
			s.WriteString(" <- ")
		}
		current = current.previous
	}
	fmt.Println(s.String())
}

func (l *LinkedList) Delete(data int) {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		if current.data == data {
			if current.previous == nil {
				current.next.previous = nil
				l.head = current.next
				return
			} else if current.next == nil {
				current.previous.next = nil
				return
			}
			current.previous.next = current.next
			current.next.previous = current.previous
			return
		}
		current = current.next
	}
	fmt.Println("not found")
}

func main() {
	ll := NewList()
	ll.PrintForward()
	ll.Insert(1)
	ll.PrintForward()
	ll.Insert(2)
	ll.Insert(3)
	ll.PrintForward()
	ll.PrintBackForward()
	ll.Delete(3)
	ll.Delete(1)
	ll.PrintForward()
	ll.PrintBackForward()
}
