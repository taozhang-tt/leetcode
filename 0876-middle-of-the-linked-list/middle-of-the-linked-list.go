package middleofthelinkedlist

// 876. 链表的中间结点

type ListNode struct {
	Val  int
	Next *ListNode
}

// 返回链表中间节点，有两个中间节点时，返回后一个
func middleNode(head *ListNode) *ListNode {
	prev := &ListNode{Next: head}
	slow, fast := prev, prev
	for fast != nil {
		if fast.Next == nil {
			return slow.Next
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

// 返回链表中间节点，有两个中间节点时，返回前一个
func middleNode2(head *ListNode) *ListNode {
	prev := &ListNode{Next: head}
	slow, fast := prev, prev
	for fast != nil {
		if fast.Next == nil {
			return slow
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

// 返回链表中间节点的前一个节点，有两个中间节点时，以后一个为准
func middleNode3(head *ListNode) *ListNode {
	prev := &ListNode{Next: &ListNode{Next: head}}
	slow, fast := prev, prev
	for fast != nil {
		if fast.Next == nil {
			return slow
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
