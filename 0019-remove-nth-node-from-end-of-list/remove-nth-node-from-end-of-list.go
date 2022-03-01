package removenthnodefromendoflist

// 19. 删除链表的倒数第 N 个结点

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	virtual := &ListNode{Next: head}
	prev := virtual
	for n > 0 {
		head = head.Next
		n--
	}
	for head != nil {
		head, prev = head.Next, prev.Next
	}
	prev.Next = prev.Next.Next
	return virtual.Next
}

func list2Slice(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}
