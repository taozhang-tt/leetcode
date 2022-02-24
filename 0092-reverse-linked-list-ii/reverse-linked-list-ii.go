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

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}

// 返还链表的K个节点
// K <= 节点数
func reverseK(head *ListNode, K int) *ListNode {
	var prev *ListNode
	root := head
	for K > 0 {
		head.Next, prev, head = prev, head, head.Next
		K--
	}
	root.Next = head
	return prev
}
