# LIST

##### Content:

* list_Node.go
  ```go
  type ListNode struct { ... }
  func ListNode_Init(value interface{}) *ListNode { ... }
  func ListNode_Create(value interface{}, prev *ListNode, next *ListNode) *ListNode { ... }
  ```

* list_LinkedList.go
  ```go
  type LinkedList struct { ... }
  func (l *LinkedList) View() { ... }
  func (l *LinkedList) ViewFull() { ... }
  func (l *LinkedList) Head() *ListNode { ... }
  func (l *LinkedList) Tail() *ListNode { ... }
  func (l *LinkedList) Append(node *ListNode) { ... }
  func (l *LinkedList) PopHead() *ListNode { ... }
  ```

* list_LinkedList_Single.go
  ```go
  type SinglyLinkedList struct { ... }
  func SinglyLinkedList_Init() *SinglyLinkedList { ... }
  func (l *SinglyLinkedList) InsertAfter(pointerNode *ListNode, newNode *ListNode) { ... }
  ```

* list_LinkedList_Double.go
  ```go
  type DoublyLinkedList struct { ... }
  func DoublyLinkedList_Init() *DoublyLinkedList { ... }
  func (l *DoublyLinkedList) InsertBefore(pointerNode *ListNode, newNode *ListNode) { ... }
  func (l *DoublyLinkedList) InsertAfter(pointerNode *ListNode, newNode *ListNode) { ... }
  ```
