package reversenodesinkgroup

// 25. K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/submissions/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归做法
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

// 翻转链表、递归做法
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	root := reverse(head.Next)
	head.Next.Next, head.Next = head, nil

	return root
}

//////////////////////////////

// 直接翻转
func reverseKGroup2(head *ListNode, k int) *ListNode {
	dumy := &ListNode{Next: head}
	prev := dumy
	for prev != nil {
		head, tail := reverseK2(prev.Next, k)
		prev.Next, prev = head, tail
	}
	return dumy.Next
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

// 翻转链表的前k个节点，不足k个不翻转，返回翻转部分的头尾节点
func reverseK2(head *ListNode, k int) (*ListNode, *ListNode) {
	p, cnt := head, 0
	for p != nil && cnt != k {
		cnt, p = cnt+1, p.Next
	}
	if cnt < k {
		return head, nil
	}
	var prev *ListNode
	root := head
	for k > 0 {
		head.Next, head, prev, k = prev, head.Next, head, k-1
	}
	root.Next = head
	return prev, root
}
