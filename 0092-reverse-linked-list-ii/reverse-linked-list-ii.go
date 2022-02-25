package reverselinkedlistii

// 92. 反转链表 II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	virtual := &ListNode{Next: head}
	prev := virtual
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	prev.Next = reverseK(prev.Next, right-left+1)
	return virtual.Next
}

// 翻转链表
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}

// 返还链表的k个节点 k <= 节点数
func reverseK(head *ListNode, k int) *ListNode {
	var prev *ListNode
	root := head
	for k > 0 {
		head.Next, prev, head = prev, head, head.Next
		k--
	}
	root.Next = head
	return prev
}
