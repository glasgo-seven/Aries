package Aries


type ListNode struct {
	Value	interface{}
	Next	*ListNode
	Prev	*ListNode
}

func ListNode_Init(value interface{}) *ListNode {
	return &ListNode{
		Prev	:	nil,
		Value	:	value,
		Next	:	nil,
	}
}

func ListNode_Create(value interface{}, prev *ListNode, next *ListNode) *ListNode {
	return &ListNode{
		Value	:	value,
		Next	:	next,
		Prev	:	prev,
	}
}
