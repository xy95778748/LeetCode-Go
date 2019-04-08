package LinkList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针相对于哈希表更快
// 链表是否有环
func HasCycle(head *ListNode) bool {

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 判断链表是否有环, 有环返回node
func DetectCycle(head *ListNode) *ListNode {

	dict := make(map[*ListNode]int)

	temp := head
	for temp != nil && temp.Next != nil {
		if dict[temp] == 1 {
			return temp
		} else {
			dict[temp] = 1
			temp = temp.Next
		}
	}
	return nil
}

// 最佳答案
func detectCycle(head *ListNode) *ListNode {
	oneStep := head
	twoStep := head
	for twoStep != nil && oneStep != nil {
		oneStep = oneStep.Next
		twoStep = twoStep.Next
		if twoStep == nil {
			break
		}
		twoStep = twoStep.Next
		if oneStep == twoStep {
			break
		}
	}
	if twoStep == nil {
		return nil
	}
	oneStep = head
	for twoStep != oneStep {
		oneStep = oneStep.Next
		twoStep = twoStep.Next
	}
	return oneStep
}

// 相交链表 返回相交的node
/*
 * 如果两个链表相交, 那么一定存在如在关系
 * 1. headA = headA独有 + 相交部分
 * 2. headB = headB独有 + 相交部分
 * 3. headA + headB = headA独有 + 相交部分(1) + headB独有 + 相交部分(2)
 * 4. headB + headA = headB独有 + 相交部分(1) + headA独有 + 相交部分(2)
 * 结论: 第一个指针遍历headA后遍历headB, 第二个指针遍历headB后遍历headA
         交点一定是两个链表相交的起点
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cursorA := headA
	cursorB := headB
	if cursorA == nil || cursorB == nil {
		return nil
	}

	for cursorA != cursorB {
		if cursorA == nil {
			cursorA = headB
		} else {
			cursorA = cursorA.Next
		}

		if cursorB == nil {
			cursorB = headA
		} else {
			cursorB = cursorB.Next
		}
	}
	return cursorA
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	if head.Next == nil {
		return nil
	}

	fast, slow := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 删除第一个元素
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}

func ReverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	newHead := new(ListNode)
	newHead.Next = nil
	newHead.Val = head.Val

	tempHead := head.Next

	for tempHead != nil {
		node := new(ListNode)
		node.Val = tempHead.Val
		node.Next = newHead
		newHead = node
		tempHead = tempHead.Next
	}
	return newHead
}

func RemoveElements(head *ListNode, val int) *ListNode {

	return head
}
