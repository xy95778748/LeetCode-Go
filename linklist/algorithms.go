package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

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

/*
	19. 删除链表的倒数第N个节点
	https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
*/
func removeNthFromEnd(head *SingleList, n int) *SingleList {
	if head.Next == nil {
		return nil
	}

	fast, slow := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

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

/*
	206.反转链表
	https://leetcode-cn.com/problems/reverse-linked-list/submissions/
*/
// for循环 时间, 空间复杂度O(n)
func reverseList(head *SingleList) *SingleList {
	if head == nil || head.Next == nil {
		return head
	}

	new := &SingleList{-1, nil}
	temp := head
	for temp != nil {
		new.Next = &SingleList{temp.Val, new.Next}
		temp = temp.Next
	}
	return new.Next
}

// 时间O(n), 空间O(1)
func ReverseList1(head *SingleList) *SingleList {
	cur := head
	var pre *SingleList = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}
	return pre
}

// 递归
func reverseList2(head *SingleList) *SingleList {

	if head == nil || head.Next == nil {
		return head
	}
	return reverse(nil, head)
}

func reverse(prev *SingleList, curr *SingleList) *SingleList {
	if curr == nil {
		return prev
	}
	next := curr.Next
	curr.Next = prev
	return reverse(curr, next)
}

/*
	328. 奇偶链表
	https://leetcode-cn.com/problems/odd-even-linked-list/
*/
func oddEvenList(head *SingleList) *SingleList {

	if head == nil || head.Next == nil {
		return head
	}

	tempHead := head
	odd, even := tempHead, tempHead.Next

	for even != nil && even.Next != nil {
		temp := odd.Next
		odd.Next = even.Next
		even.Next = even.Next.Next
		odd.Next.Next = temp
		odd = odd.Next
		even = even.Next
	}
	return head
}

/*
	876. 链表的中间结点
	https://leetcode-cn.com/problems/middle-of-the-linked-list/comments/
*/
func middleNode(head *SingleList) *SingleList {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

/*
	234. 回文链表
	https://leetcode-cn.com/problems/palindrome-linked-list/
*/
func isPalindrome(head *SingleList) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid := middleNode(head)  // 找中点
	end := ReverseList1(mid) // 反转后半段

	for head != nil && end != nil {
		if head.Val == end.Val {
			head, end = head.Next, end.Next
		} else {
			return false
		}
	}
	return true
}

/*
	237. 删除链表中的节点
	https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
*/

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head1, head2 := l1, l2
	newHead := &ListNode{}
	result := newHead
	index := 1
	for head1.Next != nil && head2.Next != nil {
		if index%2 == 1 {
			newHead.Next = head1
			head1, newHead = head1, newHead.Next
		} else {
			newHead.Next = head2
			head2, newHead = head2.Next, newHead.Next
		}
		index++
	}
	return result
}
