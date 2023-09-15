package Aries

import (

)


type DoublyLinkedList struct {
	LinkedList
}

func DoublyLinkedList_Init() *DoublyLinkedList {
	return &DoublyLinkedList{}
}


func (l *DoublyLinkedList) InsertBefore(pointerNode *ListNode, newNode *ListNode) {
	newNode.Prev = pointerNode.Prev
	newNode.Prev.Next = newNode
	pointerNode.Prev = newNode
	newNode.Next = pointerNode
}

func (l *DoublyLinkedList) InsertAfter(pointerNode *ListNode, newNode *ListNode) {
	newNode.Next = pointerNode.Next
	newNode.Next.Prev = newNode
	pointerNode.Next = newNode
	newNode.Prev = pointerNode
}
