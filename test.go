package goLeoLib

import (
	"fmt"
)



func Test_LinkedList() {
	var l LinkedList = *InitLinkedList()

	l.Append(&LinkedListNode{Value: 1, Next: nil})
	l.Append(&LinkedListNode{Value: 2, Next: nil})
	l.Append(&LinkedListNode{Value: 3, Next: nil})
	l.Append(&LinkedListNode{Value: 4, Next: nil})

	l.View()
	fmt.Println(l.Head())
	fmt.Println(l.Tail())

	l.InsertAfter(&LinkedListNode{Value: 0, Next: nil}, l.Head())

	fmt.Println()
	l.View()
}

func Test_DoublyLinkedList() {
	var l DoublyLinkedList = *InitDoublyLinkedList()

	l.Append(&DoublyLinkedListNode{Value: 1, Prev: nil, Next: nil})
	l.Append(&DoublyLinkedListNode{Value: 2, Prev: nil, Next: nil})
	l.Append(&DoublyLinkedListNode{Value: 3, Prev: nil, Next: nil})
	l.Append(&DoublyLinkedListNode{Value: 4, Prev: nil, Next: nil})

	l.View()
	fmt.Println(l.Head())
	fmt.Println(l.Tail())

	l.InsertBefore(&DoublyLinkedListNode{Value: 0, Prev: nil, Next: nil}, l.Tail())
	l.InsertAfter(&DoublyLinkedListNode{Value: 0, Prev: nil, Next: nil}, l.Head())

	fmt.Println()
	l.View()
}