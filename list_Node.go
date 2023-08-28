package goLeoLib



type LinkedListNode struct {
	Value interface{}
	Next *LinkedListNode
}

func Init_LinkedListNode(value interface{}) LinkedListNode {
	return LinkedListNode{Value: value, Next: nil}
}

func Create_LinkedListNode(value interface{}, next *LinkedListNode) LinkedListNode {
	return LinkedListNode{Value: value, Next: next}
}



type DoublyLinkedListNode struct {
	Value interface{}
	Prev *DoublyLinkedListNode
	Next *DoublyLinkedListNode
}

func Init_DoublyLinkedListNode(value interface{}) DoublyLinkedListNode {
	return DoublyLinkedListNode{Value: value, Prev: nil, Next: nil}
}

func Create_DoublyLinkedListNode(value interface{}, prev *DoublyLinkedListNode, next *DoublyLinkedListNode) DoublyLinkedListNode {
	return DoublyLinkedListNode{Value: value, Prev: prev, Next: next}
}
