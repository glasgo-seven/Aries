package Aries

import (
	"fmt"
)


type LinkedList struct {
	head	*ListNode
}

func (l *LinkedList) View() {
	pointer := l.head
	for pointer != nil {
		fmt.Printf("[%v] -- ", pointer.Value)
		pointer = pointer.Next
	}
	fmt.Print("nil;\n\n")
}

func (l *LinkedList) ViewFull() {
	pointer := l.head
	for pointer != nil {
		fmt.Printf("%v [%p] -- \n", pointer, pointer)
		pointer = pointer.Next
	}
	fmt.Print("nil;\n\n")
}

func (l *LinkedList) Head() *ListNode {
	return l.head
}

func (l *LinkedList) Tail() *ListNode {
	pointer := l.head
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	return pointer
}

func (l *LinkedList) Append(node *ListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.Prev = l.Tail()
		l.Tail().Next = node
	}
}

func (l *LinkedList) PopHead() *ListNode {
	ret := l.head

	if l.head.Next != nil {
		l.head = l.head.Next
		l.head.Prev = nil
	} else {
		l.head = nil
	}
	
	return ret
}
