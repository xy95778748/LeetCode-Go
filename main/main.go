package main

import (
	"LeetCode-Go/linklist"
	"fmt"
)

func main() {

	node13 := &linklist.ListNode{4, nil}
	node12 := &linklist.ListNode{2, node13}
	node11 := &linklist.ListNode{1, node12}

	node23 := &linklist.ListNode{4, nil}
	node22 := &linklist.ListNode{3, node23}
	node21 := &linklist.ListNode{1, node22}

	result := linklist.MergeTwoLists(node11, node21)

	fmt.Println(result)
}
