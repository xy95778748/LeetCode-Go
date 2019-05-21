package linklist

/*
	141. 环形链表
	https://leetcode-cn.com/problems/linked-list-cycle/
*/

func HasCycle(head *SingleList) bool {

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

/*
	142. 环形链表 II
	https://leetcode-cn.com/problems/linked-list-cycle-ii/
	// 链表有环：
	// 假设链表起点到入环的第1个节点A的长度为a，到快慢指针相遇的节点B的长度为a+b，
	// 节点B到节点A的距离为c，由于快指针的速度是慢指针的2倍，所以快指针走过的距离
	// 始终是慢指针的两倍。当它们第1次相遇时，快指针走过的距离：2*(a+b) = a+b+c+b, 即 a == c
*/
func DetectCycle(head *SingleList) *SingleList {

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

/*
	160. 相交链表
	https://leetcode-cn.com/problems/intersection-of-two-linked-lists/solution/
*/
// 方法一
func GetIntersectionNode(headA, headB *SingleList) *SingleList {

	cursorA := headA
	cursorB := headB

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

// 方法二
func GetIntersectionNode1(headA, headB *SingleList) *SingleList {

	if headA == nil || headB == nil {
		return nil
	}

	tempA, tempB := headA, headB
	lenA, lenB := 0, 0

	for tempA != nil {
		tempA = tempA.Next
		lenA++
	}

	for tempB != nil {
		tempB = tempB.Next
		lenB++
	}

	tempA, tempB = headA, headB

	if lenA > lenB {
		for lenA-lenB > 0 {
			tempA = tempA.Next
			lenA--
		}
	} else {
		for lenB-lenA > 0 {
			tempB = tempB.Next
			lenB--
		}
	}

	for tempA != nil && tempB != nil {
		if tempA == tempB {
			return tempA
		}
		tempA = tempA.Next
		tempB = tempB.Next
	}
	return nil
}
