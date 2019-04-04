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

	node1.Val = 10
	node1.Next = node2

	node2.Val = 20
	node2.Next = node3

	node3.Val = 30
	node3.Next = node4

	node4.Val = 40
	node4.Next = node2

	//fmt.Println(LinkList.HasCycle(node1))

	fmt.Println(LinkList.DetectCycle(node1))
}
