package goLeoLib

import (
	"fmt"
)



type LinkedList struct {
	head *LinkedListNode
}

func InitLinkedList() *LinkedList {
	return &LinkedList{nil}
}

func (l *LinkedList) View() {
	pointer := l.head
	for pointer != nil {
		fmt.Printf("%v [%p] -> \n", pointer, pointer)
		pointer = pointer.Next
	}
	fmt.Print("nil;\n\n")
}

func (l *LinkedList) Head() *LinkedListNode {
	return l.head
}

func (l *LinkedList) Tail() *LinkedListNode {
	pointer := l.head
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	return pointer
}

func (l *LinkedList) Append(node *LinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		l.Tail().Next = node
	}
}

func (l *LinkedList) InsertAfter(newNode *LinkedListNode, pointerNode *LinkedListNode) {
	newNode.Next = pointerNode.Next
	pointerNode.Next = newNode
}




