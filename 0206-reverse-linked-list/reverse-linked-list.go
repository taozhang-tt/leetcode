package reverselinkedlist

// 206. 翻转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 直接翻转
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}

// 递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	final := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return final
}
