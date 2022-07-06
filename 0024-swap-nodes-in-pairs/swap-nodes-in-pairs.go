package swapnodesinpairs

// 24. 两两交换链表中的节点

type ListNode struct {
	Val  int
	Next *ListNode
}

// 直接反转
// 反转两个节点 a -> b，需要记录 a 的前缀节点 prev
func swapPairs(head *ListNode) *ListNode {
	virtual := &ListNode{Next: head}
	prev := virtual
	for prev.Next != nil && prev.Next.Next != nil {
		a, b := prev.Next, prev.Next.Next
		prev.Next, b.Next, a.Next, prev = b, a, b.Next, a
	}
	return virtual.Next
}

// 递归
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	a, b, c := head, head.Next, swapPairs2(head.Next.Next)
	a.Next, b.Next = c, a
	return b
}

// 直接反转
// 反转两个节点 a -> b，需要记录 a 的前缀节点 prev
func swapPairs3(head *ListNode) *ListNode {
	dumy := &ListNode{Next: head}
	prev := dumy
	for head != nil && head.Next != nil {
		prev.Next, head.Next, head.Next.Next, head, prev = head.Next, head.Next.Next, head, head.Next.Next, head
	}
	return dumy.Next
}

// 递归
func swapPairs4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next, root := swapPairs4(head.Next.Next), head.Next
	head.Next, head.Next.Next = next, head
	return root
}
