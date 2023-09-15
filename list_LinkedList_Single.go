package Aries

import (
	
)


type SinglyLinkedList struct {
	LinkedList
}

func SinglyLinkedList_Init() *SinglyLinkedList {
	return &SinglyLinkedList{}
}


func (l *SinglyLinkedList) InsertAfter(pointerNode *ListNode, newNode *ListNode) {
	newNode.Next = pointerNode.Next
	pointerNode.Next = newNode
}
