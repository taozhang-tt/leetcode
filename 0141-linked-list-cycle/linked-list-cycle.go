package linkedlistcycle

// 141. 环形链表
// https://leetcode.cn/problems/linked-list-cycle/

type ListNode struct {
	Val  int
	Next *ListNode
}

// set 判重
func hasCycle(head *ListNode) bool {
	set := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := set[head]; ok {
			return true
		} else {
			set[head] = struct{}{}
		}
		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycle2(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
