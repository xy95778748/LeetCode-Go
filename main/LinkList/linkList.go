package LinkList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针相对于哈希表更快
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
	one_step := head
	two_step := head
	for two_step != nil && one_step != nil {
		one_step = one_step.Next
		two_step = two_step.Next
		if two_step == nil {
			break
		}
		two_step = two_step.Next
		if one_step == two_step {
			break
		}
	}
	if two_step == nil {
		return nil
	}
	one_step = head
	for two_step != one_step {
		one_step = one_step.Next
		two_step = two_step.Next
	}
	return one_step
}
