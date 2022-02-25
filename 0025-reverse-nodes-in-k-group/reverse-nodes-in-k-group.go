package reversenodesinkgroup

// 25. K 个一组翻转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	prev := &ListNode{Next: head}
	for i := 0; i < k; i++ {
		prev = prev.Next
		if prev == nil {
			return head
		}
	}
	next := reverseKGroup(prev.Next, k)
	prev.Next = nil
	reverse(head)
	head.Next = next
	return prev
}

// 翻转链表
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev.Next
}

// k 个一组切割链表
func cutK(head *ListNode, k int) []*ListNode {
	ret := make([]*ListNode, 0)
	if head == nil {
		return ret
	}
	prev := &ListNode{Next: head}
	for i := 0; i < k; i++ {
		prev = prev.Next
		if prev == nil {
			ret = append(ret, head)
			return ret
		}
	}
	ret = append(ret, head)
	ret = append(ret, cutK(prev.Next, k)...)
	prev.Next = nil
	return ret
}

// 翻转链表的前k个节点，k <= 节点数
func reverseK(head *ListNode, k int) *ListNode {
	root := head
	var prev *ListNode
	for k > 0 {
		head.Next, head, prev = prev, head.Next, head
		k--
	}
	root.Next = head
	return prev
}
