package main

import (
	"fmt"
)

func (this *MyLinkedList) Log() {
	temp := this
	for temp != nil {
		fmt.Println("value =", temp.Val)
		temp = temp.Next
	}
}

func main() {

	node := MyLinkedList{}
	node.AddAtHead(1)
	node.AddAtTail(3)
	node.AddAtIndex(1, 2)
	fmt.Println(node.Get(1))
	node.DeleteAtIndex(1)
	fmt.Println(node.Get(1))
	node.Log()
}

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {

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
func (this *MyLinkedList) AddAtHead(val int) {
	old := this.Next
	this.Next = &MyLinkedList{
		Val:  val,
		Next: old,
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	temp := this
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &MyLinkedList{
		Val: val,
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {

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
		p.Next = &MyLinkedList{
			Val: val,
		}
	} else {
		old := p.Next
		p.Next = &MyLinkedList{
			Val:  val,
			Next: old,
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
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
