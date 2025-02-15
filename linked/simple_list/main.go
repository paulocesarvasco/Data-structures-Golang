package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) Delete(data int) {
	if l.head.data == data {
		l.head = l.head.next
		return
	}
	current := l.head.next
	previous := l.head
	for current != nil {
		if current.data == data {
			previous.next = current.next
			return
		}
		previous = current
		current = current.next
	}
}

func (l *LinkedList) Print() {
	list := ""
	current := l.head
	for current != nil {
		list += fmt.Sprint(current.data)
		if current.next != nil {
			list += " -> "
		}
		current = current.next
	}
	fmt.Println(list)
}

func main() {
	ll := NewLinkedList()
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Delete(4)
	ll.Insert(5)
	ll.Delete(1)
	ll.Print()
}
