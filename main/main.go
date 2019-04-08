package main

import (
	"LeetCode-Go/main/LinkList"
	"fmt"
)

func main() {

	node1 := new(LinkList.ListNode)
	node2 := new(LinkList.ListNode)
	node3 := new(LinkList.ListNode)
	node4 := new(LinkList.ListNode)
	node5 := new(LinkList.ListNode)

	node1.Val = 10
	node1.Next = nil

	node2.Val = 10
	node2.Next = node3

	node3.Val = 30
	node3.Next = node4

	node4.Val = 40
	node4.Next = node5

	node5.Val = 50
	node5.Next = nil

	//fmt.Println(LinkList.HasCycle(node1))

	//fmt.Println(LinkList.DetectCycle(node1))

	//node :=  LinkList.RemoveNthFromEnd(node1, 2)

	//node := LinkList.ReverseList(node1)

	node := LinkList.RemoveElements(node1, 10)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

