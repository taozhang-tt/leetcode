package linkedlistcycle

// 141. 环形链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// set 判重
func hasCycle(head *ListNode) bool {
	storage := make(map[*ListNode]bool)
	for head != nil {
		if storage[head] {
			return true
		} else {
			storage[head] = true
		}
		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycle2(head *ListNode) bool {
	virtual := &ListNode{Next: head}
	slow, fast := virtual, virtual
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
