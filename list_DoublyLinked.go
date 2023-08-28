package goLeoLib

import (
	"fmt"
)



type DoublyLinkedList struct {
	head *DoublyLinkedListNode
}

func InitDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil}
}

func (l *DoublyLinkedList) View() {
	pointer := l.head
	for pointer != nil {
		fmt.Printf("%v [%p] -> \n", pointer, pointer)
		pointer = pointer.Next
	}
	fmt.Print("nil;\n\n")
}

func (l *DoublyLinkedList) Head() *DoublyLinkedListNode {
	return l.head
}

func (l *DoublyLinkedList) Tail() *DoublyLinkedListNode {
	pointer := l.head
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	return pointer
}

func (l *DoublyLinkedList) Append(node *DoublyLinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.Prev = l.Tail()
		l.Tail().Next = node
	}
}

func (l *DoublyLinkedList) InsertBefore(newNode *DoublyLinkedListNode, pointerNode *DoublyLinkedListNode) {
	newNode.Prev = pointerNode.Prev
	newNode.Prev.Next = newNode
	pointerNode.Prev = newNode
	newNode.Next = pointerNode
}

func (l *DoublyLinkedList) InsertAfter(newNode *DoublyLinkedListNode, pointerNode *DoublyLinkedListNode) {
	newNode.Next = pointerNode.Next
	newNode.Next.Prev = newNode
	pointerNode.Next = newNode
	newNode.Prev = pointerNode
}




