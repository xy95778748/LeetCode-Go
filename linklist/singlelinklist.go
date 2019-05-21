package linklist

/*
 链表头 Val = 0
 链表头不算链表的元素, 所以 Get(0) AddAtHead() 实际都在this.next
*/

type SingleList struct {
	Val  int
	Next *SingleList
}

/** Initialize your data structure here. */
func Constructor() SingleList {
	return SingleList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *SingleList) Get(index int) int {

	if index < 0 {
		return -1
	}

	p := this.Next
	if p == nil {
		return -1
	}
	for i := 1; i <= index; i++ {
		if p.Next == nil {
			return -1
		}
		p = p.Next
	}
	if p == nil {
		return -1
	}
	return p.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *SingleList) AddAtHead(val int) {
	old := this.Next
	this.Next = &SingleList{
		Val:  val,
		Next: old,
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *SingleList) AddAtTail(val int) {
	temp := this
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &SingleList{
		Val: val,
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *SingleList) AddAtIndex(index int, val int) {

	if index < 0 {
		this.AddAtHead(val)
	}

	p := this
	for i := 1; i <= index; i++ {
		if p.Next == nil {
			return
		}
		p = p.Next
	}
	if p.Next == nil {
		p.Next = &SingleList{
			Val: val,
		}
	} else {
		old := p.Next
		p.Next = &SingleList{
			Val:  val,
			Next: old,
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *SingleList) DeleteAtIndex(index int) {
	p := this
	for i := 1; i <= index; i++ {
		if p.Next == nil {
			return
		}
		p = p.Next
	}
	if p.Next == nil {
		return
	}
	p.Next = p.Next.Next
}
